package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/widget"
	"gosynth/gui/theme"
)

const ModuleUWidth = 65
const ModuleHeight = 500

type Module struct {
	*component.Component
	mouseDelta *control.MouseDelta
	outerType  component.IComponent
	title      string
}

func NewModule(title string, widthUnit int, outerType component.IComponent) *Module {
	m := &Module{
		Component:  component.NewComponent(),
		outerType:  outerType,
		mouseDelta: control.NewMouseDelta(),
		title:      title,
	}

	l := m.GetLayout()
	l.GetPadding().SetAll(5)
	l.SetAbsolutePositioning(true)
	l.GetSize().Set(widthUnit*ModuleUWidth, ModuleHeight)

	text := widget.NewText(title, theme.Fonts.Title, theme.Colors.Text)
	text.GetLayout().GetMargin().SetBottom(5)
	m.Append(text)

	m.GetGraphic().GetDispatcher().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := m.GetGraphic().GetImage()
		image.Fill(theme.Colors.Background)
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 2, theme.Colors.Off, false)
	})

	m.GetDispatcher().AddListener(&m, control.LeftMouseDownEvent, func(e event.IEvent) {
		m.mouseDelta.Start()
		m.GetGraphic().GetOptions().ColorScale.ScaleAlpha(0.95)
		e.StopPropagation()
	})

	m.GetDispatcher().AddListener(&m, control.LeftMouseUpEvent, func(e event.IEvent) {
		m.mouseDelta.Stop()
		m.GetGraphic().GetOptions().ColorScale.Reset()
		e.StopPropagation()
	})

	m.GetDispatcher().AddListener(&m, control.FocusEvent, func(e event.IEvent) {
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
		m.GetLayout().GetPosition().MoveBy(m.mouseDelta.GetDelta())
	}
}
