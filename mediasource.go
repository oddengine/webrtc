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

type MediaSource struct {
	fd unsafe.Pointer
}

func (me *MediaSource) Init() *MediaSource {
	return me
}

func (me *MediaSource) Remote() bool {
	return false
}

func (me *MediaSource) State(track *MediaStreamTrack) string {
	return ""
}
