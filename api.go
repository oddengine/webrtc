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
	mtx_       sync.Mutex
	observers_ = make(map[uintptr]interface{})
)

// RTCConstraints dictionary is used to describe a set of rtc library.
type RTCConstraints struct {
	KeyframeInterval int64 // ms. 0: auto.
	Logger           struct {
		Directory string
		MaxSize   int64
		History   int64
	}
}

func InitializeLibrary(path string, constraints *RTCConstraints) error {
	file := C.CString(path)

	var config C.raw_rtc_constraints_t
	config.keyframe_interval = C.int64_t(constraints.KeyframeInterval)
	config.logger.directory = C.CString(constraints.Logger.Directory)
	config.logger.max_size = C.size_t(constraints.Logger.MaxSize)
	config.logger.history = C.size_t(constraints.Logger.History)

	defer func() {
		C.free(unsafe.Pointer(file))
		C.free(unsafe.Pointer(config.logger.directory))
	}()

	eno := int(C.InitializeLibrary(file, config))
	if eno != 0 {
		fmt.Printf("Failed to initialize library: %d\n", eno)
		return fmt.Errorf("error %d", eno)
	}
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
	Release()
}

type MediaStreamTrackInterface interface {
	ID() string
	Kind() string
	Muted() bool
	State() string
	GetSource() MediaSourceInterface
	Stop()
	Release()

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
	Release()

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
	Release()
}

type RtpReceiverInterface interface {
	Track() MediaStreamTrackInterface
	Streams() []MediaStreamInterface
	GetParameters() interface{}
	GetStats() map[string]interface{}
	Release()
}

type RtpTransceiverInterface interface {
	Direction() string
	Mid() string
	Receiver() RtpReceiverInterface
	Sender() RtpSenderInterface
	SetCodecPreferences(codecs []RtpCodecCapability)
	SetDirection(new_direction string) error
	Stop()
	Release()
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
	ConnectionState() string
	IceConnectionState() string
	IceGatheringState() string
	SignalingState() string
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
	Release()

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

func LogInfo(message string) {
	msg := C.CString(message)
	defer func() {
		C.free(unsafe.Pointer(msg))
	}()

	C.LogInfo(msg)
}

func LogWarn(message string) {
	msg := C.CString(message)
	defer func() {
		C.free(unsafe.Pointer(msg))
	}()

	C.LogWarn(msg)
}

func LogError(message string) {
	msg := C.CString(message)
	defer func() {
		C.free(unsafe.Pointer(msg))
	}()

	C.LogError(msg)
}

func LogInfof(format string, args ...interface{}) {
	LogInfo(fmt.Sprintf(format, args...))
}

func LogWarnf(format string, args ...interface{}) {
	LogWarn(fmt.Sprintf(format, args...))
}

func LogErrorf(format string, args ...interface{}) {
	LogError(fmt.Sprintf(format, args...))
}

//export __onsignalingchange__
func __onsignalingchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnSignalingChange(C.GoString(new_state))
	}
}

//export __ondatachannel__
func __ondatachannel__(target unsafe.Pointer, data_channel unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnDataChannel(data_channel)
	}
}

//export __onrenegotiationneeded__
func __onrenegotiationneeded__(target unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnRenegotiationNeeded()
	}
}

//export __onconnectionchange__
func __onconnectionchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnConnectionChange(C.GoString(new_state))
	}
}

//export __oniceconnectionchange__
func __oniceconnectionchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnIceConnectionChange(C.GoString(new_state))
	}
}

//export __onicegatheringchange__
func __onicegatheringchange__(target unsafe.Pointer, new_state *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnIceGatheringChange(C.GoString(new_state))
	}
}

//export __onicecandidate__
func __onicecandidate__(target unsafe.Pointer, candidate *C.char, sdp_mid *C.char, sdp_mline_index C.int) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnIceCandidate(&IceCandidate{
			Candidate:     C.GoString(candidate),
			SDPMid:        C.GoString(sdp_mid),
			SDPMLineIndex: int(sdp_mline_index),
		})
	}
}

//export __onicecandidateerror__
func __onicecandidateerror__(target unsafe.Pointer, address *C.char, port C.int, url *C.char, error_code C.int, error_text *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	pc := (*PeerConnection)(target)
	if pc != nil {
		pc.OnIceCandidateError(C.GoString(address), int(port), C.GoString(url), int(error_code), C.GoString(error_text))
	}
}

//export __ontrack__
func __ontrack__(target unsafe.Pointer, transceiver unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
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
		pc.OnTrack(track, streams...)
	}
}

//export __oncreatesessiondescriptionsuccess__
func __oncreatesessiondescriptionsuccess__(target unsafe.Pointer, typ *C.char, sdp *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	ob := (*CreateSessionDescriptionObserver)(target)
	if ob != nil && ob.OnSuccess != nil {
		ob.OnSuccess(SessionDescription{
			Type: C.GoString(typ),
			SDP:  C.GoString(sdp),
		})
	}
	ob.release()
}

//export __oncreatesessiondescriptionfailure__
func __oncreatesessiondescriptionfailure__(target unsafe.Pointer, name *C.char, message *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	ob := (*CreateSessionDescriptionObserver)(target)
	if ob != nil && ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
	ob.release()
}

//export __onsetsessiondescriptionsuccess__
func __onsetsessiondescriptionsuccess__(target unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	ob := (*SetSessionDescriptionObserver)(target)
	if ob != nil && ob.OnSuccess != nil {
		ob.OnSuccess()
	}
	ob.release()
}

//export __onsetsessiondescriptionfailure__
func __onsetsessiondescriptionfailure__(target unsafe.Pointer, name *C.char, message *C.char) {
	defer func() {
		if err := recover(); err != nil {
			LogErrorf("Unexpected error occurred: %v", err)
			debug.PrintStack()
		}
	}()

	ob := (*SetSessionDescriptionObserver)(target)
	if ob != nil && ob.OnFailure != nil {
		ob.OnFailure(new(RTCError).Init(C.GoString(name), C.GoString(message)))
	}
	ob.release()
}
