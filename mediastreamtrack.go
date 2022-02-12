package rawrtc

/*
#cgo CFLAGS : -g -I../../include
#cgo LDFLAGS: -L../../lib -ldl

#include "api.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type MediaStreamTrack struct {
	fd unsafe.Pointer

	OnEnded  func()
	OnMute   func()
	OnUnmute func()
}

func (me *MediaStreamTrack) Init() *MediaStreamTrack {
	return me
}

func (me *MediaStreamTrack) ID() string {
	id := C.MediaStreamTrackGetID(me.fd)
	defer func() {
		C.free(unsafe.Pointer(id))
	}()
	return C.GoString(id)
}

func (me *MediaStreamTrack) Kind() string {
	kind := C.MediaStreamTrackGetKind(me.fd)
	defer func() {
		C.free(unsafe.Pointer(kind))
	}()
	return C.GoString(kind)
}

func (me *MediaStreamTrack) Muted() bool {
	return (int)(C.MediaStreamTrackGetMuted(me.fd)) != 0
}

func (me *MediaStreamTrack) State() string {
	return C.GoString(C.MediaStreamTrackGetState(me.fd))
}

func (me *MediaStreamTrack) GetSource() *MediaSource {
	source := new(MediaSource).Init()
	source.fd = C.MediaStreamTrackGetSource(me.fd)
	if source.fd == nil {
		fmt.Printf("Failed to MediaStreamTrackGetSource\n")
		return nil
	}
	return source
}

func (me *MediaStreamTrack) Stop() {
	C.MediaStreamTrackStop(me.fd)
}
