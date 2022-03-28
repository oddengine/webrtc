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
	me.fd = C.CreatePeerConnectionFactory(unsafe.Pointer(me))
	return me
}

func (me *PeerConnectionFactory) CreatePeerConnection(config RTCConfiguration) (*PeerConnection, error) {
	pc := new(PeerConnection).Init(config)
	pc.fd = C.CreatePeerConnection(me.fd, unsafe.Pointer(pc))
	return pc, nil
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
		fmt.Printf("Failed to CreateAudioTrack: id=%s\n", id)
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
		fmt.Printf("Failed to CreateVideoTrack: id=%s\n", id)
		return nil, fmt.Errorf("unknown error occurred")
	}
	return track, nil
}

func NewPeerConnectionFactory() *PeerConnectionFactory {
	return new(PeerConnectionFactory).Init()
}
