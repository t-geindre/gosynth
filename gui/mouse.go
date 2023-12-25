package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/event"
)

type MouseEvent struct {
	x, y int
}

type Mouse struct {
	LeftIsDown   bool
	LastX, LastY int
	IsDragging   bool
	event.Dispatcher
	Events struct {
		Click     event.Event
		DragStart event.Event
		DragDelta event.Event
		DragEnd   event.Event
	}
}

func NewMouse() *Mouse {
	m := &Mouse{}
	m.Dispatcher.Init()
	m.Events.Click = m.Dispatcher.RegisterEvent()
	m.Events.DragStart = m.Dispatcher.RegisterEvent()
	m.Events.DragDelta = m.Dispatcher.RegisterEvent()
	m.Events.DragEnd = m.Dispatcher.RegisterEvent()

	return m
}

func (m *Mouse) Update() {
	if m.LeftIsDown {
		x, y := ebiten.CursorPosition()
		dx, dy := x-m.LastX, y-m.LastY
		if dx != 0 || dy != 0 {
			if !m.IsDragging {
				m.IsDragging = true
				m.Dispatch(m.Events.DragStart, nil)
			}
			m.Dispatch(m.Events.DragStart, MouseEvent{dx, dy})
			m.LastX, m.LastY = x, y
		}
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			m.LeftIsDown = false
			if m.IsDragging {
				m.IsDragging = false
				m.Dispatch(m.Events.DragEnd, nil)
				return
			}
			x, y := ebiten.CursorPosition()
			m.Dispatch(m.Events.Click, MouseEvent{x, y})
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		m.LeftIsDown = true
		m.LastX, m.LastY = ebiten.CursorPosition()
	}
}
