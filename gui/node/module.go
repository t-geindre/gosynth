package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/gui/primitive"
	"gosynth/gui/theme"
)

type Module struct {
	Node
	MouseLDown             bool
	LastMouseX, LastMouseY int
}

func NewModule(width, height int) *Module {
	m := &Module{}
	m.Node = *NewNode(width, height, m)

	// Layout
	m.Image.Fill(theme.Colors.Background)
	title := primitive.Title("VCA", width, 10)
	m.Image.DrawImage(title, nil)
	vector.StrokeRect(m.Image, 0, 0, float32(width), float32(height), 2, theme.Colors.Off, false)
	yOffset := title.Bounds().Dy() + 10

	// Components
	sl := NewSlider(width-20, height-300)
	sl.SetPosition(10, yOffset)
	sl.SetRange(0, 1)
	sl.SetValue(0.5)
	m.Append(sl)

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
	if m == target {
		m.MouseLDown = true
		m.LastMouseX, m.LastMouseY = ebiten.CursorPosition()
		m.Options.ColorScale.ScaleAlpha(0.8)
	}
	m.Node.MouseLeftDown(target)
}

func (m *Module) MouseLeftUp(target INode) {
	if m == target {
		m.Options.ColorScale.Reset()
		m.MouseLDown = false
	}
	m.Node.MouseLeftUp(target)
}
