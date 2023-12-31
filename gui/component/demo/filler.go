package demo

import (
	"gosynth/gui/component"
)

type Filler struct {
	*component.Component
}

func NewFiller(fill int) *Filler {
	f := &Filler{
		Component: component.NewComponent(),
	}

	f.GetLayout().SetFill(fill)

	return f
}
