package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/behavior"
	"gosynth/gui/component"
	"gosynth/gui/graphic"
	"gosynth/gui/theme"
	"gosynth/gui/widget"
)

const ModuleUWidth float64 = 65
const ModuleHeight float64 = 500

type Module struct {
	*component.Component
	outerType component.IComponent
	title     string
}

func NewModule(title string, widthUnit int, outerType component.IComponent) *Module {
	m := &Module{
		Component: component.NewComponent(),
		outerType: outerType,
		title:     title,
	}

	behavior.NewDraggable(m)
	behavior.NewFocusable(m)

	l := m.GetLayout()
	l.SetAbsolutePositioning(true)
	l.GetPadding().SetAll(10)
	l.GetSize().Set(float64(widthUnit)*ModuleUWidth, ModuleHeight)

	text := widget.NewText(title, widget.TextSizeTitle)
	text.GetLayout().GetMargin().SetBottom(10)
	m.Append(text)

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := m.GetGraphic().GetImage()
		image.Fill(theme.Colors.Background)
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 2, theme.Colors.Off, false)
	})

	return m
}
