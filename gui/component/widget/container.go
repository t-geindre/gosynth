package widget

import (
	"gosynth/gui/component"
)

type Container struct {
	*component.Component
}

func NewContainer() *Container {
	c := &Container{
		Component: component.NewComponent(),
	}

	c.Component.GetGraphic().Disable()

	return c
}
