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

var (
	factory ILoggerFactory
	logger  ILogger
)

func InitializeLibrary(path string, constraints *DefaultWriterConstraints) error {
	file := C.CString(path)
	defer func() {
		C.free(unsafe.Pointer(file))
	}()

	eno := int(C.InitializeLibrary(file))
	if eno != 0 {
		Errorf("Failed to initialize library: %d", eno)
		return fmt.Errorf("error %d", eno)
	}

	factory = NewDefaultLoggerFactory(constraints)
	logger = factory.NewLogger("RTC")
	return nil
}

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
	RemoveTrack(sender RtpSenderInterface) error
	AddTransceiver(kind string, init RtpTransceiverInit) (RtpTransceiverInterface, error)
	CreateOffer(observer *CreateSessionDescriptionObserver)
	CreateAnswer(observer *CreateSessionDescriptionObserver)
	SetLocalDescription(observer *SetSessionDescriptionObserver, desc *SessionDescription)
	SetRemoteDescription(observer *SetSessionDescriptionObserver, desc *SessionDescription)
	AddIceCandidate(candidate *IceCandidate) bool
	GetReceivers() []RtpReceiverInterface
	GetSenders() []RtpSenderInterface
	GetTransceivers() []RtpTransceiverInterface
	GetStats() map[string]interface{}
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

//export __onloggerwriterresize__
func __onloggerwriterresize__(target unsafe.Pointer) {
	ob := (*DefaultWriter)(target)
	ob.OnResize()
}

//export __onsignalingchange__
func __onsignalingchange__(target unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(target)
	ob.OnSignalingChange(C.GoString(new_state))
}

//export __ondatachannel__
func __ondatachannel__(target unsafe.Pointer, data_channel unsafe.Pointer) {
	ob := (*PeerConnection)(target)
	ob.OnDataChannel(data_channel)
}

//export __onrenegotiationneeded__
func __onrenegotiationneeded__(target unsafe.Pointer) {
	ob := (*PeerConnection)(target)
	ob.OnRenegotiationNeeded()
}

//export __onconnectionchange__
func __onconnectionchange__(target unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(target)
	ob.OnConnectionChange(C.GoString(new_state))
}

//export __oniceconnectionchange__
func __oniceconnectionchange__(target unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(target)
	ob.OnIceConnectionChange(C.GoString(new_state))
}

//export __onicegatheringchange__
func __onicegatheringchange__(target unsafe.Pointer, new_state *C.char) {
	ob := (*PeerConnection)(target)
	ob.OnIceGatheringChange(C.GoString(new_state))
}

//export __onicecandidate__
func __onicecandidate__(target unsafe.Pointer, candidate *C.char, sdp_mid *C.char, sdp_mline_index C.int) {
	ob := (*PeerConnection)(target)
	ob.OnIceCandidate(&IceCandidate{
		Candidate:     C.GoString(candidate),
		SDPMid:        C.GoString(sdp_mid),
		SDPMLineIndex: int(sdp_mline_index),
	})
}

//export __onicecandidateerror__
func __onicecandidateerror__(target unsafe.Pointer, address *C.char, port C.int, url *C.char, error_code C.int, error_text *C.char) {
	ob := (*PeerConnection)(target)
	ob.OnIceCandidateError(C.GoString(address), int(port), C.GoString(url), int(error_code), C.GoString(error_text))
}

//export __ontrack__
func __ontrack__(target unsafe.Pointer, transceiver unsafe.Pointer) {
	trans := new(RtpTransceiver).Init()
	trans.fd = transceiver

	receiver := trans.Receiver()
	track := receiver.Track()
	streams := receiver.Streams()

	ob := (*PeerConnection)(target)
	ob.OnTrack(track, streams...)
}

//export __oncreatesessiondescriptionsuccess__
func __oncreatesessiondescriptionsuccess__(target unsafe.Pointer, typ *C.char, sdp *C.char) {
	ob := (*CreateSessionDescriptionObserver)(target)
	if ob.OnSuccess != nil {
		ob.OnSuccess(SessionDescription{
			Type: C.GoString(typ),
			SDP:  C.GoString(sdp),
		})
	}
}

//export __oncreatesessiondescriptionfailure__
func __oncreatesessiondescriptionfailure__(target unsafe.Pointer, name *C.char, message *C.char) {
	ob := (*CreateSessionDescriptionObserver)(target)
	if ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
}

//export __onsetsessiondescriptionsuccess__
func __onsetsessiondescriptionsuccess__(target unsafe.Pointer) {
	ob := (*SetSessionDescriptionObserver)(target)
	if ob.OnSuccess != nil {
		ob.OnSuccess()
	}
}

//export __onsetsessiondescriptionfailure__
func __onsetsessiondescriptionfailure__(target unsafe.Pointer, name *C.char, message *C.char) {
	ob := (*SetSessionDescriptionObserver)(target)
	if ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
}
