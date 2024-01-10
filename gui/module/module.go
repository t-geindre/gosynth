package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
	"gosynth/gui/theme"
	"gosynth/gui/widget"
)

const ModuleUWidth float64 = 65
const ModuleHeight float64 = 500

type Module struct {
	*component.Component
	title string
}

func NewModule(title string, widthUnit int) *Module {
	m := &Module{
		Component: component.NewComponent(),
		title:     title,
	}

	l := m.GetLayout()
	l.SetAbsolutePositioning(true)
	l.SetPadding(10, 10, 10, 10)
	l.SetSize(float64(widthUnit)*ModuleUWidth, ModuleHeight)

	text := widget.NewTitle(title, widget.TitlePositionTop)
	m.Append(text)

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := m.GetGraphic().GetImage()
		image.Fill(theme.Colors.Background)
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 2, theme.Colors.Off, false)
	})

	return m
}
