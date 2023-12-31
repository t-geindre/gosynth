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

	for k := 0; k < 5; k++ {
		for l := 0; l < 4; l++ {
			draggable := NewDraggable(nil)
			dl := draggable.GetLayout()
			dl.GetPosition().Set(k*300, l*300)
			dl.GetSize().Set(300, 300)
			dl.GetPadding().SetAll(10)
			dl.SetContentOrientation(layout.Horizontal)
			d.Append(draggable)

			resize := NewResize(300, 300)
			draggable.Append(resize)

			draggable.Append(NewFiller(50))

			containers := make([]*Container, 0)
			for i := 0; i < 5; i++ {
				container := NewContainer()
				if i > 0 {
					container.GetLayout().GetMargin().SetLeft(10)
				}
				container.GetLayout().GetWantedSize().SetWidth(100)
				draggable.Append(container)
				containers = append(containers, container)
			}

			for _, contains := range containers {
				contains.Append(NewFiller(50))
				btnCount := 5
				for j := 0; j < btnCount; j++ {
					clickable := NewButton()
					clickable.GetLayout().GetWantedSize().SetHeight(100)
					if j > 0 {
						clickable.GetLayout().GetMargin().SetTop(10)
					}
					contains.Append(clickable)
				}

				contains.Append(NewFiller(50))
			}

			draggable.Append(NewFiller(50))
		}
	}

	return d
}
