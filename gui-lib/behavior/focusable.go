package behavior

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
)

type Focusable struct {
	node component.IComponent
}

func NewFocusable(node component.IComponent) *Focusable {
	f := &Focusable{
		node: node,
	}

	node.AddListener(&f, control.FocusEvent, func(e event.IEvent) {
		if p := node.GetParent(); p != nil {
			p.MoveFront(node)
		}
	})

	return f
}

func (f *Focusable) Remove() {
	f.node.RemoveListener(&f, control.FocusEvent)
}
