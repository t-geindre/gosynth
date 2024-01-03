package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	control2 "gosynth/gui/control"
	"gosynth/gui/graphic"
	"gosynth/gui/theme"
	"gosynth/gui/widget"
)

const ModuleUWidth float64 = 65
const ModuleHeight float64 = 500

type Module struct {
	*component.Component
	mouseDelta *control2.MouseDelta
	outerType  component.IComponent
	title      string
}

func NewModule(title string, widthUnit int, outerType component.IComponent) *Module {
	m := &Module{
		Component:  component.NewComponent(),
		outerType:  outerType,
		mouseDelta: control2.NewMouseDelta(),
		title:      title,
	}

	l := m.GetLayout()
	l.GetPadding().SetAll(10)
	l.SetAbsolutePositioning(true)
	l.GetSize().Set(float64(widthUnit)*ModuleUWidth, ModuleHeight)

	text := widget.NewText(title, theme.Fonts.Title)
	text.GetLayout().GetMargin().SetBottom(10)
	m.Append(text)

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := m.GetGraphic().GetImage()
		image.Fill(theme.Colors.Background)
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 2, theme.Colors.Off, false)
	})

	m.AddListener(&m, control2.LeftMouseDownEvent, func(e event.IEvent) {
		m.mouseDelta.Start()
		m.GetGraphic().GetOptions().ColorScale.ScaleAlpha(0.95)
		e.StopPropagation()
	})

	m.AddListener(&m, control2.LeftMouseUpEvent, func(e event.IEvent) {
		m.mouseDelta.Stop()
		m.GetGraphic().GetOptions().ColorScale.Reset()
		e.StopPropagation()
	})

	m.AddListener(&m, control2.FocusEvent, func(e event.IEvent) {
		if p := m.GetParent(); p != nil {
			if m.outerType != nil {
				p.MoveFront(m.outerType)
				return
			}
			p.MoveFront(m)
		}
	})

	return m
}

func (m *Module) Update() {
	if m.mouseDelta.IsActive() {
		dx, dy := m.mouseDelta.GetDelta()
		m.GetLayout().GetPosition().MoveBy(float64(dx), float64(dy))
	}

	m.Component.Update()
}
