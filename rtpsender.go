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
	"unsafe"
)

type RtpSender struct {
	fd unsafe.Pointer
}

func (me *RtpSender) Init() *RtpSender {
	return me
}

func (me *RtpSender) SetTrack(track *MediaStreamTrack) bool {
	eno := (int)(C.RtpSenderSetTrack(me.fd, track.fd))
	if eno != 0 {
		fmt.Printf("Failed to RtpSenderSetTrack\n")
		return false
	}
	return true
}

func (me *RtpSender) Track() *MediaStreamTrack {
	track := new(MediaStreamTrack).Init()
	track.fd = C.RtpSenderGetTrack(me.fd)
	if track.fd == nil {
		fmt.Printf("Failed to RtpSenderGetTrack\n")
		return nil
	}
	return track
}

func (me *RtpSender) SetStreams(stream_ids ...string) {
	ids := make([]*C.char, 0)
	for _, id := range stream_ids {
		ids = append(ids, C.CString(id))
	}
	C.RtpSenderSetStreams(me.fd, (C.size_t)(len(ids)), (**C.char)(&ids[0]))
}

func (me *RtpSender) Streams() []string {
	var (
		dst        C.raw_array_t
		buf        = make([]unsafe.Pointer, 8)
		stream_ids = make([]string, 0)
	)

	dst.size = 8
	dst.elements = (*unsafe.Pointer)(&buf[0])

	C.RtpSenderGetStreams(me.fd, &dst)
	for i := 0; i < (int)(dst.size); i++ {
		ptr := (unsafe.Pointer)(&buf[i])
		cs := *(**C.char)(ptr)
		defer func(cs *C.char) {
			C.free(unsafe.Pointer(cs))
		}(cs)
		id := C.GoString(cs)
		stream_ids = append(stream_ids, id)
	}
	return stream_ids
}

func (me *RtpSender) SetParameters(parameters interface{}) {

}

func (me *RtpSender) GetParameters() interface{} {
	return nil
}

func (me *RtpSender) GetStats() map[string]interface{} {
	return nil
}
