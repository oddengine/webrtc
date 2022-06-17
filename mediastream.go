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
		C.Free(unsafe.Pointer(id))
	}()
	return C.GoString(id)
}

func (me *MediaStream) AddTrack(track *MediaStreamTrack) bool {
	eno := (int)(C.MediaStreamAddTrack(me.fd, track.fd))
	if eno != 0 {
		LogErrorf("Failed to MediaStreamAddTrack.")
		return false
	}
	return true
}

func (me *MediaStream) RemoveTrack(track *MediaStreamTrack) bool {
	eno := (int)(C.MediaStreamRemoveTrack(me.fd, track.fd))
	if eno != 0 {
		LogErrorf("Failed to MediaStreamRemoveTrack.")
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
		LogErrorf("Failed to MediaStreamFindAudioTrack: id=%s", id)
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
		LogErrorf("Failed to MediaStreamFindVideoTrack: id=%s", id)
		return nil
	}
	return track
}

func (me *MediaStream) Release() {
	C.MediaStreamRelease(me.fd)
}
