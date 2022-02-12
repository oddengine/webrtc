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

type MediaStream struct {
	fd unsafe.Pointer

	OnAddTrack    func(track *MediaStreamTrack)
	OnRemoveTrack func(track *MediaStreamTrack)
}

func (me *MediaStream) Init() *MediaStream {
	return me
}

func (me *MediaStream) ID() string {
	id := C.MediaStreamGetID(me.fd)
	defer func() {
		C.free(unsafe.Pointer(id))
	}()
	return C.GoString(id)
}

func (me *MediaStream) AddTrack(track *MediaStreamTrack) bool {
	eno := (int)(C.MediaStreamAddTrack(me.fd, track.fd))
	if eno != 0 {
		fmt.Printf("Failed to MediaStreamAddTrack\n")
		return false
	}
	return true
}

func (me *MediaStream) RemoveTrack(track *MediaStreamTrack) bool {
	eno := (int)(C.MediaStreamRemoveTrack(me.fd, track.fd))
	if eno != 0 {
		fmt.Printf("Failed to MediaStreamRemoveTrack\n")
		return false
	}
	return true
}

func (me *MediaStream) GetAudioTracks() []*MediaStreamTrack {
	var (
		size   C.size_t = 8
		array           = make([]unsafe.Pointer, 8)
		tracks          = make([]*MediaStreamTrack, 0)
	)

	C.MediaStreamGetAudioTracks(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < (int)(size); i++ {
		track := new(MediaStreamTrack).Init()
		track.fd = array[i]
		tracks = append(tracks, track)
	}
	return tracks
}

func (me *MediaStream) GetVideoTracks() []*MediaStreamTrack {
	var (
		size   C.size_t = 8
		array           = make([]unsafe.Pointer, 8)
		tracks          = make([]*MediaStreamTrack, 0)
	)

	C.MediaStreamGetVideoTracks(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < (int)(size); i++ {
		track := new(MediaStreamTrack).Init()
		track.fd = array[i]
		tracks = append(tracks, track)
	}
	return tracks
}

func (me *MediaStream) FindAudioTrack(id string) *MediaStreamTrack {
	track_id := C.CString(id)
	defer func() {
		C.free(unsafe.Pointer(track_id))
	}()

	track := new(MediaStreamTrack).Init()
	track.fd = C.MediaStreamFindAudioTrack(me.fd, track_id)
	if track.fd == nil {
		fmt.Printf("Failed to MediaStreamFindAudioTrack: id=%s\n", id)
		return nil
	}
	return track
}

func (me *MediaStream) FindVideoTrack(id string) *MediaStreamTrack {
	track_id := C.CString(id)
	defer func() {
		C.free(unsafe.Pointer(track_id))
	}()

	track := new(MediaStreamTrack).Init()
	track.fd = C.MediaStreamFindVideoTrack(me.fd, track_id)
	if track.fd == nil {
		fmt.Printf("Failed to MediaStreamFindVideoTrack: id=%s\n", id)
		return nil
	}
	return track
}
