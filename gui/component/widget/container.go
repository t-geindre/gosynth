package widget

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/theme"
)

type Container struct {
	*component.Component
	inverted bool
}

func NewContainer() *Container {
	c := &Container{
		Component: component.NewComponent(),
	}

	c.GetGraphic().GetDispatcher().AddListener(&c, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := c.GetGraphic().GetImage()
		img.Clear()

		if c.inverted {
			c.GetGraphic().GetImage().Fill(theme.Colors.BackgroundInverted)
		}
	})

	return c
}

func (c *Container) SetInverted(inverted bool) {
	c.inverted = inverted
	c.GetGraphic().ScheduleUpdate()
}

func (c *Container) GetInverted() bool {
	return c.inverted
}
