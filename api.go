package rawrtc

/*
#cgo CFLAGS : -g -I../../include
#cgo LDFLAGS: -L../../lib -ldl

#include "api.h"
*/
import "C"
import (
	"unsafe"
)

func LoadLibrary(path string) {
	file := C.CString(path)
	defer func() {
		C.free(unsafe.Pointer(file))
	}()

	C.LoadLibrary(file)
}

const (
	TRACK_KIND_AUDIO = "audio"
	TRACK_KIND_VIDEO = "video"
)

const (
	SOURCE_STATE_INITIALIZING uint32 = iota
	SOURCE_STATE_LIVE
	SOURCE_STATE_ENDED
	SOURCE_STATE_MUTED
)

const (
	TRACK_STATE_LIVE uint32 = iota
	TRACK_STATE_ENDED
)

const (
	RTP_TRANSCEIVER_DIRECTION_SENDRECV = "sendrecv"
	RTP_TRANSCEIVER_DIRECTION_SENDONLY = "sendonly"
	RTP_TRANSCEIVER_DIRECTION_RECVONLY = "recvonly"
	RTP_TRANSCEIVER_DIRECTION_INACTIVE = "inactive"
)

type PeerConnectionFactoryInterface interface {
	CreatePeerConnection(config RTCConfiguration) (PeerConnectionInterface, error)
	CreateAudioTrack(id string, source MediaSourceInterface) (MediaStreamTrackInterface, error)
	CreateVideoTrack(id string, source MediaSourceInterface) (MediaStreamTrackInterface, error)
}

type MediaSourceInterface interface {
	Remote() bool
	State() string
}

type MediaStreamTrackInterface interface {
	ID() string
	Kind() string
	Muted() bool
	State() string
	GetSource() MediaSourceInterface
	Stop()

	// OnEnded  func()
	// OnMute   func()
	// OnUnmute func()
}

type MediaStreamInterface interface {
	ID() string
	AddTrack(track MediaStreamTrackInterface) bool
	RemoveTrack(track MediaStreamTrackInterface) bool
	GetAudioTracks() []MediaStreamTrackInterface
	GetVideoTracks() []MediaStreamTrackInterface
	FindAudioTrack(id string) MediaStreamTrackInterface
	FindVideoTrack(id string) MediaStreamTrackInterface

	// OnAddTrack    func(track MediaStreamTrackInterface)
	// OnRemoveTrack func(track MediaStreamTrackInterface)
}

type RtpSenderInterface interface {
	SetTrack(track MediaStreamTrackInterface) bool
	Track() MediaStreamTrackInterface
	SetStreams(stream_ids ...string)
	Streams() []string
	SetParameters(parameters interface{})
	GetParameters() interface{}
	GetStats() map[string]interface{}
}

type RtpReceiverInterface interface {
	Track() MediaStreamTrackInterface
	Streams() []MediaStreamInterface
	GetParameters() interface{}
	GetStats() map[string]interface{}
}

type RtpTransceiverInterface interface {
	Direction() string
	Mid() string
	Receiver() RtpReceiverInterface
	Sender() RtpSenderInterface
	SetDirection(new_direction string) error
	Stop()
}

type RtpTransceiverInit struct {
	Direction string
	StreamIDs []string
}

type PeerConnectionInterface interface {
	AddTrack(track MediaStreamTrackInterface, streams ...MediaStreamInterface) (RtpSenderInterface, error)
	RemoveTrack(track MediaStreamTrackInterface) error
	AddTransceiver(kind string, init RtpTransceiverInit) (RtpTransceiverInterface, error)
	CreateOffer(observer *CreateSessionDescriptionObserver)
	CreateAnswer(observer *CreateSessionDescriptionObserver)
	SetLocalDescription(observer *SetSessionDescriptionObserver, desc *SessionDescription)
	SetRemoteDescription(observer *SetSessionDescriptionObserver, desc *SessionDescription)
	AddIceCandidate(candidate *IceCandidate) bool
	GetReceivers() []RtpReceiverInterface
	GetSenders() []RtpSenderInterface
	GetTransceivers() []RtpTransceiverInterface
	Close()

	// OnSignalingChange     func(new_state string)
	// OnDataChannel         func(data_channel interface{})
	// OnRenegotiationNeeded func()
	// OnConnectionChange    func(new_state string)
	// OnIceConnectionChange func(new_state string)
	// OnIceGatheringChange  func(new_state string)
	// OnIceCandidate        func(candidate *IceCandidate)
	// OnIceCandidateError   func(address string, port int, url string, error_code int, error_text string)
	// OnTrack               func(track MediaStreamTrackInterface, streams ...MediaStreamInterface)
}

//export __onsignalingchange__
func __onsignalingchange__(observer unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(observer)
	ob.OnSignalingChange(C.GoString(new_state))
}

//export __ondatachannel__
func __ondatachannel__(observer unsafe.Pointer, data_channel unsafe.Pointer) {
	ob := (*PeerConnection)(observer)
	ob.OnDataChannel(data_channel)
}

//export __onrenegotiationneeded__
func __onrenegotiationneeded__(observer unsafe.Pointer) {
	ob := (*PeerConnection)(observer)
	ob.OnRenegotiationNeeded()
}

//export __onconnectionchange__
func __onconnectionchange__(observer unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(observer)
	ob.OnConnectionChange(C.GoString(new_state))
}

//export __oniceconnectionchange__
func __oniceconnectionchange__(observer unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(observer)
	ob.OnIceConnectionChange(C.GoString(new_state))
}

//export __onicegatheringchange__
func __onicegatheringchange__(observer unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(observer)
	ob.OnIceGatheringChange(C.GoString(new_state))
}

//export __onicecandidate__
func __onicecandidate__(observer unsafe.Pointer, candidate *C.char, sdp_mid *C.char, sdp_mline_index C.int) {
	ob := (*PeerConnection)(observer)
	ob.OnIceCandidate(&IceCandidate{
		Candidate:     C.GoString(candidate),
		SDPMid:        C.GoString(sdp_mid),
		SDPMLineIndex: int(sdp_mline_index),
	})
}

//export __onicecandidateerror__
func __onicecandidateerror__(observer unsafe.Pointer, address *C.char, port C.int, url *C.char, error_code C.int, error_text *C.char) {
	ob := (*PeerConnection)(observer)
	ob.OnIceCandidateError(C.GoString(address), int(port), C.GoString(url), int(error_code), C.GoString(error_text))
}

//export __ontrack__
func __ontrack__(observer unsafe.Pointer, transceiver unsafe.Pointer) {
	trans := new(RtpTransceiver).Init()
	trans.fd = transceiver

	receiver := trans.Receiver()
	track := receiver.Track()
	streams := receiver.Streams()

	ob := (*PeerConnection)(observer)
	ob.OnTrack(track, streams...)
}

//export __oncreatesessiondescriptionsuccess__
func __oncreatesessiondescriptionsuccess__(observer unsafe.Pointer, typ *C.char, sdp *C.char) {
	ob := (*CreateSessionDescriptionObserver)(observer)
	if ob.OnSuccess != nil {
		ob.OnSuccess(SessionDescription{
			Type: C.GoString(typ),
			SDP:  C.GoString(sdp),
		})
	}
}

//export __oncreatesessiondescriptionfailure__
func __oncreatesessiondescriptionfailure__(observer unsafe.Pointer, name *C.char, message *C.char) {
	ob := (*CreateSessionDescriptionObserver)(observer)
	if ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
}

//export __onsetsessiondescriptionsuccess__
func __onsetsessiondescriptionsuccess__(observer unsafe.Pointer) {
	ob := (*SetSessionDescriptionObserver)(observer)
	if ob.OnSuccess != nil {
		ob.OnSuccess()
	}
}

//export __onsetsessiondescriptionfailure__
func __onsetsessiondescriptionfailure__(observer unsafe.Pointer, name *C.char, message *C.char) {
	ob := (*SetSessionDescriptionObserver)(observer)
	if ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
}
