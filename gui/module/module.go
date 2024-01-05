package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui-lib/behavior"
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

	behavior.NewDraggable(m)
	behavior.NewFocusable(m)

	m.AddListener(&m, behavior.DragEvent, func(e event.IEvent) {
		dEv := e.(*behavior.DragEventDetails)
		px, py := m.GetLayout().GetPosition()
		m.GetLayout().SetPosition(float64(dEv.DeltaX)+px, float64(dEv.DeltaY)+py)
		e.StopPropagation()
	})

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
