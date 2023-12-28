package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/theme"
	"time"
)

type Module struct {
	*Node
	MouseLDown             bool
	LastMouseX, LastMouseY int
}

func NewModule(width, height int, inode INode) *Module {
	m := &Module{}
	m.Node = NewNode(width, height, inode)

	tl := NewLabel("VCA", theme.Fonts.Title)
	m.AppendWithOptions(tl, NewAppendOptions().Margins(5, 10, 0, 0).HorizontallyCentered())

	m.Dispatcher.AddListener(&m, LeftMouseDownEvent, func(e event.IEvent) {
		m.OnMouseLeftDown(e.GetSource().(INode))
	})

	m.Dispatcher.AddListener(&m, LeftMouseUpEvent, func(e event.IEvent) {
		m.MouseLeftUp(e.GetSource().(INode))
	})

	return m
}

func (m *Module) Clear() {
	if m.Dirty {
		m.Image.Fill(theme.Colors.Background)
		vector.StrokeRect(m.Image, 0, 0, float32(m.Width), float32(m.Height), 2, theme.Colors.Off, false)

		m.Dirty = false
	}
	m.Node.Clear()
}

func (m *Module) Update(time time.Duration) error {
	if m.MouseLDown {
		x, y := ebiten.CursorPosition()
		m.MoveBy(x-m.LastMouseX, y-m.LastMouseY)
		m.LastMouseX, m.LastMouseY = x, y
	}

	return m.Node.Update(time)
}
func (m *Module) OnMouseLeftDown(target INode) {
	if m.GetParent() != nil {
		m.GetParent().GetINode().MoveFront(m.GetINode())
	}

	if m.GetINode() == target {
		m.MouseLDown = true
		m.LastMouseX, m.LastMouseY = ebiten.CursorPosition()
		m.Options.ColorScale.ScaleAlpha(0.8)
	}
}

func (m *Module) MouseLeftUp(target INode) {
	if m.GetINode() == target {
		m.Options.ColorScale.Reset()
		m.MouseLDown = false
	}
}

func (m *Module) SetParent(parent INode) {
	m.Node.SetParent(parent)
	m.Node.SetAppendOptions(NewAppendOptions().Padding(10))
}
