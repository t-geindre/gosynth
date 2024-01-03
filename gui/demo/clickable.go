package demo

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/control"
	"gosynth/gui/graphic"
	"image/color"
)

type Clickable struct {
	*component.Component
	color    color.RGBA
	onCol    color.RGBA
	offColor color.RGBA
}

func NewButton() *Clickable {
	c := randomColor()
	b := &Clickable{
		Component: component.NewComponent(),
		color:     c,
		offColor:  c,
		onCol:     colorInverse(c),
	}

	b.GetGraphic().AddListener(&b, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := b.GetGraphic().GetImage()
		img.Fill(b.color)
	})

	b.AddListener(&b, control.LeftMouseDownEvent, func(e event.IEvent) {
		b.color = b.onCol
		b.GetGraphic().ScheduleUpdate()
		e.StopPropagation()
	})

	b.AddListener(&b, control.LeftMouseUpEvent, func(e event.IEvent) {
		b.color = b.offColor
		b.GetGraphic().ScheduleUpdate()
		e.StopPropagation()
	})

	return b
}
