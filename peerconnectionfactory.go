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

type PeerConnectionFactory struct {
	fd unsafe.Pointer
}

func (me *PeerConnectionFactory) Init() *PeerConnectionFactory {
	return me
}

func (me *PeerConnectionFactory) CreatePeerConnection(config RTCConfiguration) (*PeerConnection, error) {
	pc := new(PeerConnection).Init(config)
	pc.fd = C.CreatePeerConnection(me.fd, unsafe.Pointer(pc))
	return pc, nil
}

func (me *PeerConnectionFactory) GetRtpSenderCapabilities(kind string) RtpCapabilities {
	var (
		capabilities RtpCapabilities
	)

	return capabilities
}

func (me *PeerConnectionFactory) GetRtpReceiverCapabilities(kind string) RtpCapabilities {
	var (
		capabilities RtpCapabilities
	)

	k := C.CString(kind)
	defer func() {
		C.free((unsafe.Pointer(k)))
	}()

	rtp_capabilities := C.GetRtpReceiverCapabilities(me.fd, k)
	size := int(rtp_capabilities.size)

	for i := 0; i < size; i++ {
		arr := (*[1024]C.raw_rtp_codec_capability_t)(unsafe.Pointer(rtp_capabilities.codecs))[:size:size]
		src := arr[i]
		dst := new(RtpCodecCapability)
		dst.fd = src.fd
		dst.MimeType = C.GoString(src.mime_type)
		dst.ClockRate = int(src.clock_rate)
		dst.Channels = int(src.channels)
		capabilities.Codecs = append(capabilities.Codecs, *dst)
	}
	return capabilities
}

func (me *PeerConnectionFactory) CreateAudioTrack(id string, source *MediaSource) (*MediaStreamTrack, error) {
	var (
		track_id = C.CString(id)
	)

	defer func() {
		C.free(unsafe.Pointer(track_id))
	}()

	track := new(MediaStreamTrack).Init()
	track.fd = C.CreateAudioTrack(me.fd, unsafe.Pointer(track), track_id, source.fd)
	if track.fd == nil {
		logger_.Errorf("Failed to CreateAudioTrack: id=%s", id)
		return nil, fmt.Errorf("unknown error occurred")
	}
	return track, nil
}

func (me *PeerConnectionFactory) CreateVideoTrack(id string, source *MediaSource) (*MediaStreamTrack, error) {
	var (
		track_id = C.CString(id)
	)

	defer func() {
		C.free(unsafe.Pointer(track_id))
	}()

	track := new(MediaStreamTrack).Init()
	track.fd = C.CreateVideoTrack(me.fd, unsafe.Pointer(track), track_id, source.fd)
	if track.fd == nil {
		logger_.Errorf("Failed to CreateVideoTrack: id=%s", id)
		return nil, fmt.Errorf("unknown error occurred")
	}
	return track, nil
}

func NewPeerConnectionFactory() *PeerConnectionFactory {
	f := new(PeerConnectionFactory).Init()
	f.fd = C.CreatePeerConnectionFactory(unsafe.Pointer(f))
	return f
}
