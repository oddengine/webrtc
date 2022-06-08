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
	"runtime/debug"
	"sync"
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
	factory_   ILoggerFactory
	logger_    ILogger
	mtx_       sync.Mutex
	observers_ = make(map[uintptr]interface{})
)

// RTCConstraints dictionary is used to describe a set of rtc library.
type RTCConstraints struct {
	KeyframeInterval int64 // ms. 0: auto.
}

func InitializeLibrary(path string, constraints *RTCConstraints, writer *DefaultWriterConstraints) error {
	file := C.CString(path)
	defer func() {
		C.free(unsafe.Pointer(file))
	}()

	var config C.raw_rtc_constraints_t
	config.keyframe_interval = C.int64_t(constraints.KeyframeInterval)

	eno := int(C.InitializeLibrary(file, config))
	if eno != 0 {
		Errorf("Failed to initialize library: %d", eno)
		return fmt.Errorf("error %d", eno)
	}

	factory_ = NewDefaultLoggerFactory(writer)
	logger_ = factory_.NewLogger("RTC")
	return nil
}

type PeerConnectionFactoryInterface interface {
	CreatePeerConnection(config RTCConfiguration) (PeerConnectionInterface, error)
	GetRtpSenderCapabilities(kind string) RtpCapabilities
	GetRtpReceiverCapabilities(kind string) RtpCapabilities
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
	SetCodecPreferences(codecs []RtpCodecCapability)
	SetDirection(new_direction string) error
	Stop()
}

type RtpCapabilities struct {
	Codecs []RtpCodecCapability
}

type RtpCodecCapability struct {
	fd unsafe.Pointer

	MimeType  string
	ClockRate int
	Channels  int
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
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	dw := (*DefaultWriter)(target)
	if dw != nil {
		dw.OnResize()
	}
}

//export __onsignalingchange__
func __onsignalingchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnSignalingChange(in)")
		pc.OnSignalingChange(C.GoString(new_state))
		println("OnSignalingChange(out)")
	}
}

//export __ondatachannel__
func __ondatachannel__(target unsafe.Pointer, data_channel unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnDataChannel(in)")
		pc.OnDataChannel(data_channel)
		println("OnDataChannel(out)")
	}
}

//export __onrenegotiationneeded__
func __onrenegotiationneeded__(target unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnRenegotiationNeeded(in)")
		pc.OnRenegotiationNeeded()
		println("OnRenegotiationNeeded(out)")
	}
}

//export __onconnectionchange__
func __onconnectionchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnConnectionChange(in)")
		pc.OnConnectionChange(C.GoString(new_state))
		println("OnConnectionChange(out)")
	}
}

//export __oniceconnectionchange__
func __oniceconnectionchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnIceConnectionChange(in)")
		pc.OnIceConnectionChange(C.GoString(new_state))
		println("OnIceConnectionChange(out)")
	}
}

//export __onicegatheringchange__
func __onicegatheringchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnIceGatheringChange(in)")
		pc.OnIceGatheringChange(C.GoString(new_state))
		println("OnIceGatheringChange(out)")
	}
}

//export __onicecandidate__
func __onicecandidate__(target unsafe.Pointer, candidate *C.char, sdp_mid *C.char, sdp_mline_index C.int) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnIceCandidate(in)")
		pc.OnIceCandidate(&IceCandidate{
			Candidate:     C.GoString(candidate),
			SDPMid:        C.GoString(sdp_mid),
			SDPMLineIndex: int(sdp_mline_index),
		})
		println("OnIceCandidate(out)")
	}
}

//export __onicecandidateerror__
func __onicecandidateerror__(target unsafe.Pointer, address *C.char, port C.int, url *C.char, error_code C.int, error_text *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnIceCandidateError(in)")
		pc.OnIceCandidateError(C.GoString(address), int(port), C.GoString(url), int(error_code), C.GoString(error_text))
		println("OnIceCandidateError(out)")
	}
}

//export __ontrack__
func __ontrack__(target unsafe.Pointer, transceiver unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	trans := new(RtpTransceiver).Init()
	trans.fd = transceiver

	receiver := trans.Receiver()
	track := receiver.Track()
	streams := receiver.Streams()

	pc := (*PeerConnection)(target)
	if pc != nil {
		println("OnTrack(in)")
		pc.OnTrack(track, streams...)
		println("OnTrack(out)")
	}
}

//export __oncreatesessiondescriptionsuccess__
func __oncreatesessiondescriptionsuccess__(target unsafe.Pointer, typ *C.char, sdp *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	println("OnCreateSessionDescriptionSuccess(in)")
	ob := (*CreateSessionDescriptionObserver)(target)
	if ob != nil && ob.OnSuccess != nil {
		ob.OnSuccess(SessionDescription{
			Type: C.GoString(typ),
			SDP:  C.GoString(sdp),
		})
	}
	ob.release()
	println("OnCreateSessionDescriptionSuccess(out)")
}

//export __oncreatesessiondescriptionfailure__
func __oncreatesessiondescriptionfailure__(target unsafe.Pointer, name *C.char, message *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	println("OnCreateSessionDescriptionFailure(in)")
	ob := (*CreateSessionDescriptionObserver)(target)
	if ob != nil && ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
	ob.release()
	println("OnCreateSessionDescriptionFailure(out)")
}

//export __onsetsessiondescriptionsuccess__
func __onsetsessiondescriptionsuccess__(target unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	println("OnSetSessionDescriptionSuccess(in)")
	ob := (*SetSessionDescriptionObserver)(target)
	if ob != nil && ob.OnSuccess != nil {
		ob.OnSuccess()
	}
	ob.release()
	println("OnSetSessionDescriptionSuccess(out)")
}

//export __onsetsessiondescriptionfailure__
func __onsetsessiondescriptionfailure__(target unsafe.Pointer, name *C.char, message *C.char) {
	defer func() {
		if err := recover(); err != nil {
			logger_.Errorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	println("OnSetSessionDescriptionFailure(in)")
	ob := (*SetSessionDescriptionObserver)(target)
	if ob != nil && ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
	ob.release()
	println("OnSetSessionDescriptionFailure(out)")
}
