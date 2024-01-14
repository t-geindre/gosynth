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
		vector.StrokeRect(img, 1, 1, float32(w)-1, float32(h)-1, 1, theme.Colors.BackgroundInverted, false)
	})

	return m
}

func (m *Menu) SetParent(c component.IComponent) {
	if m.container == nil {
		m.container = c.GetRoot()
		m.container.AddListener(&m, control.FocusEvent, func(e event.IEvent) {
			m.Close()
		})
	}
	if !m.opened && c != nil {
		c.Remove(m)
		return
	}
	m.Component.SetParent(c)
}

func (m *Menu) AddOption(label string, callback func()) {
	op := component.NewContainer()
	op.GetLayout().SetContentOrientation(layout.Horizontal)
	op.GetLayout().SetPadding(3, 3, 5, 5)
	lbl := NewLabel(label, LabelPositionTop)
	op.Append(lbl)
	op.AddListener(&m, control.LeftMouseUpEvent, func(e event.IEvent) {
		callback()
		m.Close()
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
	m.opened = true
	m.container.Append(m)
	x, y := ebiten.CursorPosition()
	m.GetLayout().SetPosition(float64(x), float64(y))
}

func (m *Menu) Close() {
	if m.container == nil || m.GetParent() == nil {
		return
	}
	m.container.Remove(m)
	m.opened = false
}
