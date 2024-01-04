package widget

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/theme"
)

type Plug struct {
	*component.Image
	inverted bool
}

func NewPlug() *Plug {
	p := &Plug{
		Image: component.NewImage(theme.Images.Plug),
	}

	return p
}
