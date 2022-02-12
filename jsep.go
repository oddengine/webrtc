package rawrtc

type CreateSessionDescriptionObserver struct {
	OnSuccess func(desc SessionDescription)
	OnFailure func(err *RTCError)
}

func (me *CreateSessionDescriptionObserver) Init(onSuccess func(desc SessionDescription), onFailure func(err *RTCError)) *CreateSessionDescriptionObserver {
	me.OnSuccess = onSuccess
	me.OnFailure = onFailure
	return me
}

func NewCreateSessionDescriptionObserver(onSuccess func(desc SessionDescription), OnFailure func(err *RTCError)) *CreateSessionDescriptionObserver {
	return new(CreateSessionDescriptionObserver).Init(onSuccess, OnFailure)
}

type SetSessionDescriptionObserver struct {
	OnSuccess func()
	OnFailure func(err *RTCError)
}

func (me *SetSessionDescriptionObserver) Init(onSuccess func(), onFailure func(err *RTCError)) *SetSessionDescriptionObserver {
	me.OnSuccess = onSuccess
	me.OnFailure = onFailure
	return me
}

func NewSetSessionDescriptionObserver(onSuccess func(), onFailure func(err *RTCError)) *SetSessionDescriptionObserver {
	return new(SetSessionDescriptionObserver).Init(onSuccess, onFailure)
}
