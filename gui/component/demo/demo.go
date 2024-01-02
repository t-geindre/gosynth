package demo

import (
	"fmt"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/layout"
	"gosynth/gui/component/widget"
	"gosynth/gui/theme"
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

	for i := float64(0); i < 2; i++ {
		draggable := NewDraggable(nil)

		f := NewFiller(50)
		f.Append(widget.NewText("50%", theme.Fonts.Title))
		draggable.Append(f)

		draggable.GetLayout().GetPosition().SetX(i * 310)
		draggable.GetLayout().GetSize().Set(300, 300)
		draggable.GetLayout().GetPadding().SetAll(10)
		draggable.Append(NewResize(30, 30))

		if i > 0 {
			draggable.GetLayout().SetContentOrientation(layout.Horizontal)
		}

		for j := float64(0); j < 5; j++ {
			btn := NewButton()

			if i > 0 {
				if j > 0 {
					btn.GetLayout().GetMargin().SetLeft(10)
				}
				btn.GetLayout().GetMargin().SetTop(10)
				btn.GetLayout().GetMargin().SetBottom(10)
			} else {
				if j > 0 {
					btn.GetLayout().GetMargin().SetTop(10)
				}
				btn.GetLayout().GetMargin().SetLeft(10)
				btn.GetLayout().GetMargin().SetRight(10)
			}
			btn.GetLayout().GetWantedSize().Set((j+1)*50, (j+1)*50)
			btn.Append(widget.NewText(fmt.Sprintf("%d", int((j+1)*50)), theme.Fonts.Title))
			draggable.Append(btn)
		}

		fe := NewFiller(50)
		fe.Append(widget.NewText("50%", theme.Fonts.Title))
		draggable.Append(fe)

		d.Append(draggable)
	}

	return d
}
