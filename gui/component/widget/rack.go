package widget

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"image/color"
)

type Rack struct {
	*component.Component
	MouseDelta *control.MouseDelta
}

func NewRack() *Rack {
	r := &Rack{
		Component:  component.NewComponent(),
		MouseDelta: control.NewMouseDelta(),
	}

	r.GetGraphic().SetUpdateFunc(func() {
		image := r.GetGraphic().GetImage()
		image.Fill(color.RGBA{R: 26, G: 26, B: 26, A: 255})
	})

	r.GetDispatcher().AddListener(&r, control.LeftMouseDownEvent, func(e event.IEvent) {
		r.MouseDelta.Start()
		e.StopPropagation()
	})

	r.GetDispatcher().AddListener(&r, control.LeftMouseUpEvent, func(e event.IEvent) {
		r.MouseDelta.Stop()
		e.StopPropagation()
	})

	return r
}

func (r *Rack) Update() {
	r.GetGraphic().ScheduleUpdate()

	if r.MouseDelta.IsActive() {
		dx, dy := r.MouseDelta.GetDelta()
		for _, c := range r.GetChildren() {
			c.GetLayout().GetPosition().MoveBy(dx, dy)
		}
	}

	r.Component.Update()
}
