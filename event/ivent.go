package event

type IEvent interface {
	GetId() Id
	GetSource() any
	IsPropagationStopped() bool
	StopPropagation()
}
