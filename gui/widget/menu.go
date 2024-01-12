package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"gosynth/gui/theme"
)

type Menu struct {
	*component.Component
	options   map[string]func()
	container component.IComponent
	opened    bool
}

func NewMenu() *Menu {
	m := &Menu{
		Component: component.NewComponent(),
	}

	m.GetLayout().SetAbsolutePositioning(true)
	m.GetLayout().SetPadding(5, 5, 5, 5)

	m.GetLayout().AddListener(&m, layout.UpdateStartsEvent, func(e event.IEvent) {
		for _, c := range m.GetChildren() {
			c.GetLayout().SetFill(100/float64(len(m.GetChildren())) - 1)
		}
		m.GetLayout().SetSize(200, float64(len(m.GetChildren()))*30)
	})

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := m.GetGraphic().GetImage()
		img.Fill(theme.Colors.Background)
		w, h := m.GetLayout().GetSize()
		vector.StrokeRect(img, 0, 0, float32(w), float32(h), 1, theme.Colors.BackgroundInverted, false)
	})

	return m
}

func (m *Menu) SetContainer(c component.IComponent) {
	if m.GetParent() != nil {
		m.GetParent().Remove(m)
		m.Open()
	}
	m.container = c
}

func (m *Menu) AddOption(label string, callback func()) {
	op := component.NewContainer()
	op.GetLayout().SetContentOrientation(layout.Horizontal)
	op.GetLayout().SetPadding(3, 3, 5, 5)
	lbl := NewLabel(label, LabelPositionTop)
	op.Append(lbl)
	op.AddListener(&m, control.LeftMouseUpEvent, func(e event.IEvent) {
		callback()
	})
	op.Append(component.NewFiller(100))
	op.AddListener(&m, control.MouseEnterEvent, func(e event.IEvent) {
		control.Cursor.Push(ebiten.CursorShapePointer)
		lbl.SetBackgroundColor(theme.Colors.BackgroundInverted)
		lbl.SetColor(theme.Colors.TextInverted)
		op.SetBackgroundColor(theme.Colors.BackgroundInverted)

	})
	op.AddListener(&m, control.MouseLeaveEvent, func(e event.IEvent) {
		control.Cursor.Pop()
		lbl.SetBackgroundColor(theme.Colors.Background)
		lbl.SetColor(theme.Colors.Text)
		op.SetBackgroundColor(theme.Colors.Background)
	})
	m.Append(op)
	m.GetLayout().ScheduleUpdate()
}

func (m *Menu) Open() {
	if m.container == nil {
		return
	}
	if m.opened {
		m.Close()
	}
	m.container.Append(m)
	m.opened = true
	x, y := ebiten.CursorPosition()
	sx, sy := m.GetParent().GetLayout().GetAbsolutePosition()
	m.GetLayout().SetPosition(float64(x)-sx, float64(y)-sy)
}

func (m *Menu) Close() {
	if m.container == nil || m.GetParent() == nil {
		return
	}
	m.container.Remove(m)
	m.opened = false
}
