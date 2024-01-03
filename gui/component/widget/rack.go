package widget

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/graphic"
	"image/color"
)

type Rack struct {
	*component.Component
	mouseDelta *control.MouseDelta
	scale      float64
}

func NewRack() *Rack {
	r := &Rack{
		Component:  component.NewComponent(),
		mouseDelta: control.NewMouseDelta(),
	}

	r.scale = .1
	r.GetGraphic().GetOptions().GeoM.Scale(r.scale, r.scale)

	r.GetGraphic().AddListener(&r, graphic.DrawEvent, func(e event.IEvent) {
		r.GetGraphic().GetImage().Fill(color.RGBA{R: 26, G: 26, B: 26, A: 255})
	})

	r.	AddListener(&r, control.LeftMouseDownEvent, func(e event.IEvent) {
		r.mouseDelta.Start()
		e.StopPropagation()
	})

	r.	AddListener(&r, control.LeftMouseUpEvent, func(e event.IEvent) {
		r.mouseDelta.Stop()
		e.StopPropagation()
	})

	r.GetLayout().SetFill(100)

	return r
}

func (r *Rack) Update() {
	if r.mouseDelta.IsActive() {
		dx, dy := r.mouseDelta.GetDelta()
		for _, c := range r.GetChildren() {
			c.GetLayout().GetPosition().MoveBy(float64(dx), float64(dy))
		}
	}

	r.Component.Update()
}
