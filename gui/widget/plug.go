package widget

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui/theme"
)

type Plug struct {
	*component.Image
	inverted bool
}

func NewPlug() *Plug {
	p := &Plug{
		Image: component.NewImage(theme.Images.Plug),
	}

	p.AddListener(&p, control.LeftMouseDownEvent, func(e event.IEvent) {
		e.StopPropagation()
	})

	return p
}
