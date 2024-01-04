package demo

import (
	"fmt"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"gosynth/gui/theme"
)

type Demo struct {
	*component.Root
}

func NewDemo() *Demo {
	d := &Demo{
		Root: component.NewRoot(),
	}

	color := randomColor()
	d.GetGraphic().AddListener(&d, graphic.DrawEvent, func(e event.IEvent) {
		image := d.GetGraphic().GetImage()
		image.Fill(color)
	})

	textColor := randomColor()
	bgTextColor := colorInverse(textColor)

	for i := float64(0); i < 2; i++ {
		draggable := NewWindow(nil)

		f := component.NewFiller(40)
		f.Append(component.NewText("50%", theme.Fonts.Title, textColor, bgTextColor))
		draggable.Append(f)

		draggable.GetLayout().SetPosition(i*310, 0)
		draggable.GetLayout().SetSize(300, 300)
		draggable.GetLayout().SetPadding(10, 10, 10, 10)
		draggable.Append(NewResize(30, 30))

		if i > 0 {
			draggable.GetLayout().SetContentOrientation(layout.Horizontal)
		}

		for j := float64(0); j < 5; j++ {
			btn := NewButton()
			btn.GetLayout().SetMargin(20, 20, 20, 20)
			btn.GetLayout().SetWantedSize((j+1)*50, (j+1)*50)
			btn.Append(component.NewText(fmt.Sprintf("%d", int((j+1)*50)), theme.Fonts.Title, textColor, bgTextColor))
			draggable.Append(btn)
		}

		fe := component.NewFiller(50)
		fe.Append(component.NewText("50%", theme.Fonts.Title, textColor, bgTextColor))
		draggable.Append(fe)

		d.Append(draggable)
	}

	return d
}
