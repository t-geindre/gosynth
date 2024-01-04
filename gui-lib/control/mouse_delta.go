package control

import "github.com/hajimehoshi/ebiten/v2"

type MouseDelta struct {
	LastX, LastY int
	Active       bool
}

func NewMouseDelta() *MouseDelta {
	return &MouseDelta{
		LastX: 0,
		LastY: 0,
	}
}

func (m *MouseDelta) GetDelta() (int, int) {
	if !m.Active {
		return 0, 0
	}

	x, y := ebiten.CursorPosition()
	dx, dy := x-m.LastX, y-m.LastY
	m.LastX, m.LastY = x, y

	return dx, dy
}

func (m *MouseDelta) Start() {
	m.Active = true
	m.LastX, m.LastY = ebiten.CursorPosition()
	Cursor.Push(ebiten.CursorShapeMove)
}

func (m *MouseDelta) Stop() {
	m.Active = false
	Cursor.Pop()
}

func (m *MouseDelta) IsActive() bool {
	return m.Active
}
