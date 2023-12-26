package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/gui/theme"
)

type Module struct {
	Node
	MouseLDown             bool
	LastMouseX, LastMouseY int
}

func NewModule(width, height int, inode INode) *Module {
	m := &Module{}
	m.Node = *NewNode(width, height, inode)

	// Layout
	m.Image.Fill(theme.Colors.Background)
	vector.StrokeRect(m.Image, 0, 0, float32(width), float32(height), 2, theme.Colors.Off, false)

	tl := NewLabel(width, 35, "VCA", theme.Fonts.Title)
	tl.SetPosition(0, 0)
	m.Append(tl)

	return m
}

func (m *Module) Update() error {
	if m.MouseLDown {
		x, y := ebiten.CursorPosition()
		m.MoveBy(x-m.LastMouseX, y-m.LastMouseY)
		m.LastMouseX, m.LastMouseY = x, y
	}

	return m.Node.Update()
}
func (m *Module) MouseLeftDown(target INode) {
	if m.GetParent() != nil {
		m.GetParent().MoveFront(m)
	}
	if m.GetINode() == target {
		m.MouseLDown = true
		m.LastMouseX, m.LastMouseY = ebiten.CursorPosition()
		m.Options.ColorScale.ScaleAlpha(0.8)
	}
	m.Node.MouseLeftDown(target)
}

func (m *Module) MouseLeftUp(target INode) {
	if m.GetINode() == target {
		m.Options.ColorScale.Reset()
		m.MouseLDown = false
	}
	m.Node.MouseLeftUp(target)
}
