package event

type Id uint16

type Event struct {
	Id                 Id
	PropagationStopped bool
	Source             any
}

func NewEvent(id Id, src any) *Event {
	return &Event{
		Id:     id,
		Source: src,
	}
}

func (e *Event) StopPropagation() {
	e.PropagationStopped = true
}

func (e *Event) IsPropagationStopped() bool {
	return e.PropagationStopped
}

func (e *Event) GetId() Id {
	return e.Id
}

func (e *Event) GetSource() any {
	return e.Source
}
