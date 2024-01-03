package demo

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

	return c
}
