package event

type Event uint16
type ListenerArgs any
type Listener struct {
	Call func(event ListenerArgs)
	Id   any
}

type Dispatcher struct {
	EventCounter Event
	Listeners    map[Event][]Listener
}

func (d *Dispatcher) Init() {
	d.Listeners = make(map[Event][]Listener)
}

func (d *Dispatcher) RegisterEvent() Event {
	d.EventCounter++
	return d.EventCounter
}

func (d *Dispatcher) AddListener(Id any, event Event, call func(event ListenerArgs)) {
	listener := Listener{
		Call: call,
		Id:   Id,
	}

	d.Listeners[event] = append(d.Listeners[event], listener)
}

func (d *Dispatcher) RemoveListener(Id any, event Event) {
	for i, listener := range d.Listeners[event] {
		if listener.Id == Id {
			d.Listeners[event] = append(d.Listeners[event][:i], d.Listeners[event][i+1:]...)
			return
		}
	}
}

func (d *Dispatcher) Dispatch(event Event, args ListenerArgs) {
	for _, listener := range d.Listeners[event] {
		listener.Call(args)
	}
}
