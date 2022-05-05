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

type CreateSessionDescriptionObserver struct {
	fd unsafe.Pointer

	OnSuccess func(desc SessionDescription)
	OnFailure func(err *RTCError)
}

func (me *CreateSessionDescriptionObserver) Init(onSuccess func(desc SessionDescription), onFailure func(err *RTCError)) *CreateSessionDescriptionObserver {
	me.OnSuccess = onSuccess
	me.OnFailure = onFailure
	return me
}

func (me *CreateSessionDescriptionObserver) release() {
	C.CreateSessionDescriptionObserverRelease(me.fd)
	delete(observers_, uintptr(unsafe.Pointer(me)))
}

func NewCreateSessionDescriptionObserver(onSuccess func(desc SessionDescription), OnFailure func(err *RTCError)) *CreateSessionDescriptionObserver {
	observer := new(CreateSessionDescriptionObserver).Init(onSuccess, OnFailure)
	observer.fd = C.CreateCreateSessionDescriptionObserver(unsafe.Pointer(observer))
	return observer
}

type SetSessionDescriptionObserver struct {
	fd unsafe.Pointer

	OnSuccess func()
	OnFailure func(err *RTCError)
}

func (me *SetSessionDescriptionObserver) Init(onSuccess func(), onFailure func(err *RTCError)) *SetSessionDescriptionObserver {
	me.OnSuccess = onSuccess
	me.OnFailure = onFailure
	return me
}

func (me *SetSessionDescriptionObserver) release() {
	C.SetSessionDescriptionObserverRelease(me.fd)
	delete(observers_, uintptr(unsafe.Pointer(me)))
}

func NewSetSessionDescriptionObserver(onSuccess func(), onFailure func(err *RTCError)) *SetSessionDescriptionObserver {
	observer := new(SetSessionDescriptionObserver).Init(onSuccess, onFailure)
	observer.fd = C.CreateSetSessionDescriptionObserver(unsafe.Pointer(observer))
	return observer
}
