package behavior

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/control"
)

type Focusable struct {
	node component.IComponent
}

func NewFocusable(node component.IComponent) *Focusable {
	f := &Focusable{
		node: node,
	}

	node.AddListener(&f, control.LeftMouseDownEvent, func(e event.IEvent) {
		fEvent := event.NewEvent(FocusEvent, node)
		node.Dispatch(fEvent)

		if p := node.GetParent(); p != nil && !fEvent.IsPropagationStopped() {
			p.MoveFront(node)
		}
	})

	return f
}

func (f *Focusable) Remove() {
	f.node.RemoveListener(&f, control.LeftMouseDownEvent)
}
