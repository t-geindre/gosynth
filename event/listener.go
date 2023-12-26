package event

type ListenerArgs any
type Listener struct {
	Call func(event IEvent)
	Ref  any
}
