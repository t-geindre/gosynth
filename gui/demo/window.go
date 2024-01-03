package demo

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/behavior"
	"gosynth/gui/component"
	"gosynth/gui/control"
	"gosynth/gui/graphic"
)

type Window struct {
	*component.Component
	mouseOver bool
}

func NewWindow(outerType component.IComponent) *Window {
	w := &Window{
		Component: component.NewComponent(),
	}

	w.GetLayout().SetAbsolutePositioning(true)

	behavior.NewDraggable(w)
	behavior.NewFocusable(w)

	c := randomColor()
	bCol := colorInverse(c)

	w.GetGraphic().AddListener(&w, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := w.GetGraphic().GetImage()
		image.Fill(c)

		strokeWidth := float32(2)
		if w.mouseOver {
			strokeWidth = 20
		}
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), strokeWidth, bCol, true)
	})

	w.AddListener(&w, control.MouseEnterEvent, func(e event.IEvent) {
		w.mouseOver = true
		w.GetGraphic().ScheduleUpdate()
	})

	w.AddListener(&w, control.MouseLeaveEvent, func(e event.IEvent) {
		w.mouseOver = false
		w.GetGraphic().ScheduleUpdate()
	})

	return w
}
