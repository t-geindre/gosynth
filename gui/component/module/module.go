package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/theme"
)

const ModuleUWidth = 65
const ModuleHeight = 500

type Module struct {
	*component.Component
	MouseDelta *control.MouseDelta
	OuterType  component.IComponent
}

func NewModule(outerType component.IComponent) *Module {
	m := &Module{
		Component:  component.NewComponent(),
		OuterType:  outerType,
		MouseDelta: control.NewMouseDelta(),
	}

	l := m.GetLayout()
	l.GetPadding().SetAll(10)
	l.SetAbsolutePositioning(true)

	m.GetGraphic().SetUpdateFunc(func() {
		image := m.GetGraphic().GetImage()
		image.Fill(theme.Colors.Background)
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 2, theme.Colors.Off, false)
	})

	m.GetDispatcher().AddListener(&m, control.LeftMouseDownEvent, func(e event.IEvent) {
		m.MouseDelta.Start()
		m.GetGraphic().GetOptions().ColorScale.ScaleAlpha(0.95)
		e.StopPropagation()
	})

	m.GetDispatcher().AddListener(&m, control.LeftMouseUpEvent, func(e event.IEvent) {
		m.MouseDelta.Stop()
		m.GetGraphic().GetOptions().ColorScale.Reset()
		e.StopPropagation()
	})

	m.GetDispatcher().AddListener(&m, control.FocusEvent, func(e event.IEvent) {
		if p := m.GetParent(); p != nil {
			p.MoveFront(m.OuterType)
		}
	})

	return m
}

func (m *Module) Update() {
	if m.MouseDelta.IsActive() {
		m.GetLayout().GetPosition().MoveBy(m.MouseDelta.GetDelta())
	}
}
