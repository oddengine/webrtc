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

type RtpReceiver struct {
	fd unsafe.Pointer
}

func (me *RtpReceiver) Init() *RtpReceiver {
	return me
}

func (me *RtpReceiver) Track() *MediaStreamTrack {
	track := new(MediaStreamTrack).Init()
	track.fd = C.RtpReceiverGetTrack(me.fd)
	if track.fd == nil {
		fmt.Printf("Failed to RtpReceiverGetTrack\n")
		return nil
	}
	return track
}

func (me *RtpReceiver) Streams() []*MediaStream {
	var (
		size    C.size_t = 8
		array            = make([]unsafe.Pointer, 8)
		streams          = make([]*MediaStream, 0)
	)

	C.RtpReceiverGetStreams(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < (int)(size); i++ {
		stream := new(MediaStream).Init()
		stream.fd = array[i]
		streams = append(streams, stream)
	}
	return streams
}

func (me *RtpReceiver) GetParameters() interface{} {
	return nil
}

func (me *RtpReceiver) GetStats() map[string]interface{} {
	return nil
}
