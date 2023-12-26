package event

type Dispatcher struct {
	Listeners map[Id][]Listener
}

func (d *Dispatcher) Init() {
	d.Listeners = make(map[Id][]Listener)
}

func (d *Dispatcher) AddListener(ref any, id Id, call func(e IEvent)) {
	listener := Listener{
		Call: call,
		Ref:  ref,
	}

	d.Listeners[id] = append(d.Listeners[id], listener)
}

func (d *Dispatcher) RemoveListener(ref any, id Id) {
	for i, listener := range d.Listeners[id] {
		if listener.Ref == ref {
			d.Listeners[id] = append(d.Listeners[id][:i], d.Listeners[id][i+1:]...)
			return
		}
	}
}

func (d *Dispatcher) Dispatch(e IEvent) {
	for _, listener := range d.Listeners[e.GetId()] {
		listener.Call(e)
	}
}
