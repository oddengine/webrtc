package rawrtc

/*
#cgo CFLAGS : -g -I../../include
#cgo linux LDFLAGS: -L../../lib -ldl
#cgo windows LDFLAGS: -L../../lib

#include "api.h"
*/
import "C"
import (
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
		logger_.Errorf("Failed to RtpSenderSetTrack.")
		return false
	}
	return true
}

func (me *RtpSender) Track() *MediaStreamTrack {
	track := new(MediaStreamTrack).Init()
	track.fd = C.RtpSenderGetTrack(me.fd)
	if track.fd == nil {
		logger_.Errorf("Failed to RtpSenderGetTrack.")
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
		size       C.size_t = 8
		array               = make([]unsafe.Pointer, 8)
		stream_ids          = make([]string, 0)
	)

	C.RtpSenderGetStreams(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < int(size); i++ {
		item := (*C.char)(array[i])
		defer func(item *C.char) {
			C.free(unsafe.Pointer(item))
		}(item)

		id := C.GoString(item)
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

func (me *RtpSender) Release() {
	C.RtpSenderRelease(me.fd)
}
