package widget

import (
	"gosynth/gui/component"
)

type Filler struct {
	*component.Component
}

func NewFiller(fill float64) *Filler {
	f := &Filler{
		Component: component.NewComponent(),
	}

	f.GetLayout().SetFill(fill)

	return f
}
