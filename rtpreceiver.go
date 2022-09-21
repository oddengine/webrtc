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
		LogErrorf("Failed to RtpReceiverGetTrack.")
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
	for i := 0; i < int(size); i++ {
		stream := new(MediaStream).Init()
		stream.fd = array[i]
		streams = append(streams, stream)
	}
	return streams
}

func (me *RtpReceiver) GetParameters() RtpParameters {
	var (
		parameters RtpParameters
		dst        C.raw_rtp_parameters_t
	)

	C.RtpReceiverGetParameters(me.fd, &dst)
	size := int(dst.size)
	if size > 0 {
		defer func() {
			C.Free(unsafe.Pointer(dst.codecs))
		}()
	}

	for i := 0; i < size; i++ {
		arr := (*[1024]C.raw_rtp_codec_parameters_t)(unsafe.Pointer(dst.codecs))[:size:size]
		src := arr[i]
		defer (func(src C.raw_rtp_codec_parameters_t) {
			C.free(unsafe.Pointer(src.fd))
			C.Free(unsafe.Pointer(src.mime_type))
		})(src)

		dst := new(RtpCodecParameters)
		dst.fd = src.fd
		dst.PayloadType = int(src.payload_type)
		dst.MimeType = C.GoString(src.mime_type)
		dst.ClockRate = int(src.clock_rate)
		dst.Channels = int(src.channels)
		parameters.Codecs = append(parameters.Codecs, *dst)
	}
	return parameters
}

func (me *RtpReceiver) GetStats() map[string]interface{} {
	return nil
}

func (me *RtpReceiver) Release() {
	C.RtpReceiverRelease(me.fd)
}
