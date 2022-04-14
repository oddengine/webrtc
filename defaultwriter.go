package rawrtc

/*
#cgo CFLAGS : -g -I../../include
#cgo linux LDFLAGS: -L../../lib -ldl
#cgo windows LDFLAGS: -L../../lib

#include "api.h"
*/
import "C"
import (
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
	"unsafe"

	"gitlab.xthktech.cn/xthk-media/rawrtc/utils"
)

// Schedule Types.
const (
	SCHEDULE_DAILY    = "daily"
	SCHEDULE_DURATION = "duration"
)

// DefaultWriterConstraints dictionary is used to describe a set of rotatable writer.
type DefaultWriterConstraints struct {
	Directory string
	FileName  string
	Level     string
	Rotation  struct {
		MaxSize  int64
		Schedule struct {
			Type     string
			Duration string
		}
		History int
	}
}

// DefaultWriter represents a rotatable writer.
type DefaultWriter struct {
	fd          unsafe.Pointer
	constraints *DefaultWriterConstraints
	mtx         sync.Mutex
	files       []string
	ticker      *time.Ticker
}

// Init this class.
func (me *DefaultWriter) Init(constraints *DefaultWriterConstraints) *DefaultWriter {
	var config C.raw_default_writer_constraints_t
	config.rotation.maxsize = C.int(constraints.Rotation.MaxSize)

	me.constraints = constraints
	me.fd = C.CreateDefaultWriter(unsafe.Pointer(me), &config)

	err := utils.MkdirAll(constraints.Directory)
	if err != nil {
		panic(err)
	}
	err = me.readdir()
	if err != nil {
		panic(err)
	}
	err = me.rotate()
	if err != nil {
		panic(err)
	}
	return me
}

// Write writes len(p) bytes from p to the underlying data stream.
func (me *DefaultWriter) Write(p []byte) (int, error) {
	me.mtx.Lock()
	defer me.mtx.Unlock()

	if me.constraints.Rotation.MaxSize > 0 && me.Size()+int64(len(p)) >= me.constraints.Rotation.MaxSize {
		err := me.rotate()
		if err != nil {
			Errorf("Failed to rotate log: %s", err)
			return 0, err
		}
	}

	n := (int)(C.WriterWrite(me.fd, (*C.char)(unsafe.Pointer(&p[0])), (C.size_t)(len(p))))
	if n <= 0 {
		Errorf("Failed to WriterWrite.")
		return 0, fmt.Errorf("error %d", n)
	}
	return n, nil
}

func (me *DefaultWriter) Size() int64 {
	return (int64)(C.WriterSize(me.fd))
}

func (me *DefaultWriter) readdir() error {
	arr, err := os.ReadDir(me.constraints.Directory)
	if err != nil {
		Errorf("Failed to ReadDir: %v", err)
		return err
	}

	for _, entry := range arr {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			Warnf("Failed to get file info: %v", err)
			continue
		}
		me.files = append(me.files, info.Name())
	}
	sort.Strings(me.files)
	return nil
}

func (me *DefaultWriter) rotate() error {
	var (
		delay time.Duration
	)

	// Close current file.
	if me.fd != nil {
		C.WriterClose(me.fd)
	}

	// Remove history files.
	if me.constraints.Rotation.History > 0 {
		for len(me.files) >= me.constraints.Rotation.History {
			name := me.files[0]
			me.files = me.files[1:]

			err := os.Remove(me.constraints.Directory + name)
			if err != nil {
				Errorf("Failed to remove log: %s", err)
			}
		}
	}

	// Create a new file.
	now := time.Now()
	name := now.Format(me.constraints.FileName)
	path := C.CString(me.constraints.Directory + name)
	defer func() {
		C.free(unsafe.Pointer(path))
	}()

	eno := (int)(C.WriterOpen(me.fd, path))
	if eno != 0 {
		Errorf("Failed to WriterOpen: %d", eno)
		return fmt.Errorf("error %d", eno)
	}
	me.files = append(me.files, name)
	Debugf(0, "New log: file=%s", me.constraints.Directory+name)

	// Start ticker.
	if me.ticker != nil {
		me.ticker.Stop()
		me.ticker = nil
	}

	switch me.constraints.Rotation.Schedule.Type {
	case SCHEDULE_DAILY:
		t, err := time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02 ")+me.constraints.Rotation.Schedule.Duration, time.Local)
		if err != nil {
			Errorf("Failed to parse time: %s", err)
			return err
		}
		if now.After(t) {
			t = t.Add(24 * time.Hour)
		}
		delay = time.Until(t)
	case SCHEDULE_DURATION:
		d, err := time.ParseDuration(me.constraints.Rotation.Schedule.Duration)
		if err != nil {
			Errorf("Failed to parse duration: %s", err)
			return err
		}
		delay = d
	}

	if delay > 0 {
		Debugf(0, "About to rotate logger: delay=%dns", delay)
		me.ticker = time.NewTicker(delay)
		go me.wait()
	}
	return nil
}

func (me *DefaultWriter) OnResize() {
	// TODO(spencer@lau): If a goroutine calls Write(), which triggered OnResize,
	// it will cause a deadlock here. For now, we only call Write() in c++.
	me.mtx.Lock()
	defer me.mtx.Unlock()

	err := me.rotate()
	if err != nil {
		Errorf("Failed to rotate log: %s", err)
		return
	}
}

func (me *DefaultWriter) wait() {
	<-me.ticker.C

	me.mtx.Lock()
	defer me.mtx.Unlock()

	err := me.rotate()
	if err != nil {
		Errorf("Failed to rotate log: %s", err)
		return
	}
}
