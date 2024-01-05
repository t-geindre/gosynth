package connection

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui/theme"
)

type PlugDirection uint8

const (
	PlugDirectionIn PlugDirection = iota
	PlugDirectionOut
)

type Plug struct {
	*component.Image
	inverted  bool
	direciton PlugDirection
}

func NewPlug(direction PlugDirection) *Plug {
	p := &Plug{
		Image:     component.NewImage(theme.Images.Plug),
		direciton: direction,
	}

	p.AddListener(&p, control.LeftMouseDownEvent, func(e event.IEvent) {
		p.Dispatch(event.NewEvent(ConnectionStartEvent, p))
		e.StopPropagation()
	})

	p.AddListener(&p, control.LeftMouseUpEvent, func(e event.IEvent) {
		p.Dispatch(event.NewEvent(ConnectionStopEvent, p))
		e.StopPropagation()
	})

	p.AddListener(&p, control.MouseEnterEvent, func(e event.IEvent) {
		p.Dispatch(event.NewEvent(ConnectionEnterEvent, p))
		e.StopPropagation()
	})

	p.AddListener(&p, control.MouseLeaveEvent, func(e event.IEvent) {
		p.Dispatch(event.NewEvent(ConnectionLeaveEvent, p))
		e.StopPropagation()
	})

	return p
}

func (p *Plug) GetDirection() PlugDirection {
	return p.direciton
}
