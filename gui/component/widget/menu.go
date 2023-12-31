package widget

import (
	"gosynth/gui/component"
	"image/color"
)

type Menu struct {
	*component.Component
}

func NewMenu() *Menu {
	m := &Menu{
		Component: component.NewComponent(),
	}

	m.GetLayout().GetSize().Set(100, 100)
	m.GetGraphic().SetUpdateFunc(func() {
		m.GetGraphic().GetImage().Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	})

	return m
}
