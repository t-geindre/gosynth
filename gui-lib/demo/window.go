package demo

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	behavior2 "gosynth/gui-lib/behavior"
	component2 "gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui-lib/graphic"
)

type Window struct {
	*component2.Component
	mouseOver bool
}

func NewWindow(outerType component2.IComponent) *Window {
	w := &Window{
		Component: component2.NewComponent(),
	}

	w.GetLayout().SetAbsolutePositioning(true)

	behavior2.NewDraggable(w)
	behavior2.NewFocusable(w)

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
