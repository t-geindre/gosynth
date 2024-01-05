package connection

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui/theme"
	audio "gosynth/module"
)

type PlugDirection uint8

const (
	PlugDirectionIn PlugDirection = iota
	PlugDirectionOut
)

type Plug struct {
	*component.Image
	inverted    bool
	direction   PlugDirection
	audioModule audio.IModule
	port        audio.Port
	isOn        bool
}

func NewPlug(direction PlugDirection, module audio.IModule, port audio.Port) *Plug {
	p := &Plug{
		Image:     component.NewImage(theme.Images.Plug),
		direction: direction,
	}

	p.Bind(module, port)

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

	p.AddListener(&p, component.UpdateEvent, func(e event.IEvent) {
		on := false
		if p.audioModule != nil {
			if p.direction == PlugDirectionIn {
				on = p.audioModule.ReceiveInput(p.port) != nil
			} else {
				on = p.audioModule.ReceiveOutput(p.port) != nil
			}
		}

		if p.isOn != on {
			p.isOn = on
			p.GetGraphic().ScheduleUpdate()
		}
	})

	return p
}

func (p *Plug) GetDirection() PlugDirection {
	return p.direction
}

func (p *Plug) Bind(module audio.IModule, port audio.Port) {
	p.audioModule = module
	p.port = port
}

func (p *Plug) GetBinding() (audio.IModule, audio.Port) {
	return p.audioModule, p.port
}
