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
	"time"
	"unsafe"
)

const (
	PeerConnectionStateNew          = "new"
	PeerConnectionStateConnecting   = "connecting"
	PeerConnectionStateConnected    = "connected"
	PeerConnectionStateDisconnected = "disconnected"
	PeerConnectionStateFailed       = "failed"
	PeerConnectionStateClosed       = "closed"
)

type RTCConfiguration struct {
}

type PeerConnection struct {
	fd unsafe.Pointer

	OnSignalingChange     func(new_state string)
	OnDataChannel         func(data_channel interface{})
	OnRenegotiationNeeded func()
	OnConnectionChange    func(new_state string)
	OnIceConnectionChange func(new_state string)
	OnIceGatheringChange  func(new_state string)
	OnIceCandidate        func(candidate *IceCandidate)
	OnIceCandidateError   func(address string, port int, url string, error_code int, error_text string)
	OnTrack               func(track *MediaStreamTrack, streams ...*MediaStream)
}

func (me *PeerConnection) Init(config RTCConfiguration) *PeerConnection {
	return me
}

func (me *PeerConnection) AddTrack(track *MediaStreamTrack, streams ...*MediaStream) (*RtpSender, error) {
	var (
		err C.raw_rtc_error_t
	)

	arr := make([]unsafe.Pointer, 0)
	for _, stream := range streams {
		arr = append(arr, stream.fd)
	}
	sender := new(RtpSender).Init()
	sender.fd = C.PeerConnectionAddTrack(me.fd, track.fd, (C.size_t)(len(arr)), (*unsafe.Pointer)(&arr[0]), &err)
	if sender.fd == nil {
		logger_.Errorf("Failed to PeerConnectionAddTrack: name=%s, message=%s", C.GoString(err.name), C.GoString(err.message))
		return nil, fmt.Errorf("%s: %s", C.GoString(err.name), C.GoString(err.message))
	}
	return sender, nil
}

func (me *PeerConnection) RemoveTrack(sender *RtpSender) error {
	var (
		err C.raw_rtc_error_t
	)

	eno := (int)(C.PeerConnectionRemoveTrack(me.fd, sender.fd, &err))
	if eno != 0 {
		logger_.Errorf("Failed to PeerConnectionRemoveTrack: name=%s, message=%s", C.GoString(err.name), C.GoString(err.message))
		return fmt.Errorf("%s: %s", C.GoString(err.name), C.GoString(err.message))
	}
	return nil
}

func (me *PeerConnection) AddTransceiver(media_type string, init RtpTransceiverInit) (*RtpTransceiver, error) {
	var (
		err C.raw_rtc_error_t
	)

	kind := C.CString(media_type)
	ids := make([]*C.char, 0)
	for _, id := range init.StreamIDs {
		ids = append(ids, C.CString(id))
	}
	var cini C.raw_rtp_transceiver_init_t
	cini.direction = C.CString(init.Direction)
	cini.size = (C.size_t)(len(ids))
	cini.stream_ids = (**C.char)(&ids[0])
	defer func() {
		C.free(unsafe.Pointer(kind))
		C.free(unsafe.Pointer(cini.direction))
		for i := range ids {
			C.free(unsafe.Pointer(ids[i]))
		}
	}()

	transceiver := new(RtpTransceiver).Init()
	transceiver.fd = C.PeerConnectionAddTransceiver(me.fd, kind, &cini, &err)
	if transceiver.fd == nil {
		logger_.Errorf("Failed to PeerConnectionAddTransceiver: name=%s, message=%s", C.GoString(err.name), C.GoString(err.message))
		return nil, fmt.Errorf("%s: %s", C.GoString(err.name), C.GoString(err.message))
	}
	return transceiver, nil
}

func (me *PeerConnection) CreateOffer(observer *CreateSessionDescriptionObserver) {
	mtx_.Lock()
	ptr := uintptr(unsafe.Pointer(observer))
	observers_[ptr] = observer
	mtx_.Unlock()
	C.PeerConnectionCreateOffer(me.fd, observer.fd)
}

func (me *PeerConnection) CreateAnswer(observer *CreateSessionDescriptionObserver) {
	mtx_.Lock()
	ptr := uintptr(unsafe.Pointer(observer))
	observers_[ptr] = observer
	mtx_.Unlock()
	C.PeerConnectionCreateAnswer(me.fd, observer.fd)
}

func (me *PeerConnection) SetLocalDescription(observer *SetSessionDescriptionObserver, desc *SessionDescription) {
	var description C.raw_session_description_t
	description.typ = C.CString(desc.Type)
	description.sdp = C.CString(desc.SDP)
	defer func() {
		C.free(unsafe.Pointer(description.typ))
		C.free(unsafe.Pointer(description.sdp))
	}()

	mtx_.Lock()
	ptr := uintptr(unsafe.Pointer(observer))
	observers_[ptr] = observer
	mtx_.Unlock()
	C.PeerConnectionSetLocalDescription(me.fd, observer.fd, &description)
}

func (me *PeerConnection) SetRemoteDescription(observer *SetSessionDescriptionObserver, desc *SessionDescription) {
	var description C.raw_session_description_t
	description.typ = C.CString(desc.Type)
	description.sdp = C.CString(desc.SDP)
	defer func() {
		C.free(unsafe.Pointer(description.typ))
		C.free(unsafe.Pointer(description.sdp))
	}()

	mtx_.Lock()
	ptr := uintptr(unsafe.Pointer(observer))
	observers_[ptr] = observer
	mtx_.Unlock()
	C.PeerConnectionSetRemoteDescription(me.fd, observer.fd, &description)
}

func (me *PeerConnection) AddIceCandidate(candidate *IceCandidate) error {
	var (
		err C.raw_rtc_error_t
	)

	var cand C.raw_ice_candidate_t
	cand.candidate = C.CString(candidate.Candidate)
	cand.sdp_mid = C.CString(candidate.SDPMid)
	cand.sdp_mline_index = C.int(candidate.SDPMLineIndex)
	defer func() {
		C.free(unsafe.Pointer(cand.candidate))
		C.free(unsafe.Pointer(cand.sdp_mid))
	}()

	eno := (int)(C.PeerConnectionAddIceCandidate(me.fd, &cand, &err))
	if eno != 0 {
		logger_.Errorf("Failed to PeerConnectionAddIceCandidate: name=%s, message=%s", C.GoString(err.name), C.GoString(err.message))
		return fmt.Errorf("%s: %s", C.GoString(err.name), C.GoString(err.message))
	}
	return nil
}

func (me *PeerConnection) GetReceivers() []*RtpReceiver {
	var (
		size      C.size_t = 8
		array              = make([]unsafe.Pointer, 8)
		receivers          = make([]*RtpReceiver, 0)
	)

	C.PeerConnectionGetReceivers(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < (int)(size); i++ {
		receiver := new(RtpReceiver).Init()
		receiver.fd = array[i]
		receivers = append(receivers, receiver)
	}
	return receivers
}

func (me *PeerConnection) GetSenders() []*RtpSender {
	var (
		size    C.size_t = 8
		array            = make([]unsafe.Pointer, 8)
		senders          = make([]*RtpSender, 0)
	)

	C.PeerConnectionGetSenders(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < (int)(size); i++ {
		sender := new(RtpSender).Init()
		sender.fd = array[i]
		senders = append(senders, sender)
	}
	return senders
}

func (me *PeerConnection) GetTransceivers() []*RtpTransceiver {
	var (
		size         C.size_t = 8
		array                 = make([]unsafe.Pointer, 8)
		transceivers          = make([]*RtpTransceiver, 0)
	)

	C.PeerConnectionGetTransceivers(me.fd, &size, (*unsafe.Pointer)(&array[0]))
	for i := 0; i < (int)(size); i++ {
		transceiver := new(RtpTransceiver).Init()
		transceiver.fd = array[i]
		transceivers = append(transceivers, transceiver)
	}
	return transceivers
}

func (me *PeerConnection) GetStats() map[string]interface{} {
	return nil
}

func (me *PeerConnection) Close() {
	C.PeerConnectionClose(me.fd)
}

func (me *PeerConnection) Release() {
	// TODO(spencer@lau): Since we must call this function in or after event handler of
	// PeerConnectionState::kClosed, we do it asynchronously and wait for 5 seconds.
	go func(me *PeerConnection) {
		time.Sleep(5 * time.Second)
		C.PeerConnectionRelease(me.fd)
	}(me)
}
