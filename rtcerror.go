package rawrtc

import "fmt"

type RTCError struct {
	name    string
	message string
}

func (me *RTCError) Init(name string, message string) *RTCError {
	me.name = name
	me.message = message
	return me
}

func (me *RTCError) Name() string {
	return me.name
}

func (me *RTCError) Message() string {
	return me.message
}

func (me *RTCError) ToString() string {
	return fmt.Sprintf("[name=%s, message=%s]", me.name, me.message)
}
