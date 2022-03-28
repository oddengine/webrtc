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

type RtpTransceiver struct {
	fd unsafe.Pointer
}

func (me *RtpTransceiver) Init() *RtpTransceiver {
	return me
}

func (me *RtpTransceiver) Direction() string {
	return C.GoString(C.RtpTransceiverGetDirection(me.fd))
}

func (me *RtpTransceiver) Mid() string {
	mid := C.RtpTransceiverGetMid(me.fd)
	defer func() {
		C.free(unsafe.Pointer(mid))
	}()
	return C.GoString(mid)
}

func (me *RtpTransceiver) Receiver() *RtpReceiver {
	receiver := new(RtpReceiver).Init()
	receiver.fd = C.RtpTransceiverGetReceiver(me.fd)
	if receiver.fd == nil {
		fmt.Printf("Failed to RtpTransceiverGetReceiver\n")
		return nil
	}
	return receiver
}

func (me *RtpTransceiver) Sender() *RtpSender {
	sender := new(RtpSender).Init()
	sender.fd = C.RtpTransceiverGetSender(me.fd)
	if sender.fd == nil {
		fmt.Printf("Failed to RtpTransceiverGetSender\n")
		return nil
	}
	return sender
}

func (me *RtpTransceiver) SetDirection(new_direction string) error {
	var (
		err C.raw_rtc_error_t
	)

	direction := C.CString(new_direction)
	defer func() {
		C.free(unsafe.Pointer(direction))
	}()

	eno := (int)(C.RtpTransceiverSetDirection(me.fd, direction, &err))
	if eno != 0 {
		fmt.Printf("Failed to RtpTransceiverSetDirection: name=%s, message=%s\n", C.GoString(err.name), C.GoString(err.message))
		return fmt.Errorf("%s: %s", C.GoString(err.name), C.GoString(err.message))
	}
	return nil
}

func (me *RtpTransceiver) Stop() {
	C.RtpTransceiverStop(me.fd)
}
