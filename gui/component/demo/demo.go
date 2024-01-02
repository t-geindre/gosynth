package demo

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/layout"
)

type Demo struct {
	*component.Component
}

func NewDemo() *Demo {
	d := &Demo{
		Component: component.NewComponent(),
	}

	color := randomColor()
	d.GetGraphic().GetDispatcher().AddListener(&d, graphic.DrawEvent, func(e event.IEvent) {
		image := d.GetGraphic().GetImage()
		image.Fill(color)
	})

	draggable := NewDraggable(nil)
	draggable.Append(NewFiller(50))
	draggable.GetLayout().GetSize().Set(300, 300)
	draggable.GetLayout().GetPadding().SetAll(10)
	draggable.GetLayout().SetContentOrientation(layout.Horizontal)
	draggable.Append(NewResize(100, 100))
	for i := float64(0); i < 5; i++ {
		btn := NewButton()
		if i > 0 {
			btn.GetLayout().GetMargin().SetLeft(10)
		}
		btn.GetLayout().GetMargin().SetTop(10)
		btn.GetLayout().GetMargin().SetBottom(10)
		btn.GetLayout().GetWantedSize().Set((i+1)*50, 100)
		draggable.Append(btn)
	}

	d.Append(draggable)

	return d
}
