package component

import (
	"gosynth/event"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"image/color"
)

type Container struct {
	*Component
	inverted bool
	bgColor  color.Color
}

func NewHContainter() *Container {
	c := NewContainer()
	c.GetLayout().SetContentOrientation(layout.Horizontal)
	return c
}

func NewVContainter() *Container {
	c := NewContainer()
	c.GetLayout().SetContentOrientation(layout.Vertical)
	return c
}

func NewContainer() *Container {
	c := &Container{
		Component: NewComponent(),
		bgColor:   color.RGBA{R: 0, G: 0, B: 0, A: 0},
	}

	c.GetGraphic().AddListener(&c, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		c.GetGraphic().GetImage().Fill(c.bgColor)
	})

	return c
}

func (c *Container) SetBackgroundColor(color color.Color) {
	c.bgColor = color
	c.GetGraphic().ScheduleUpdate()
}
