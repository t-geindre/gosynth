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

func NewModule(width, height int) *Module {
	m := &Module{}
	m.Node = *NewNode(width, height, m)

	// Layout
	m.Image.Fill(theme.Colors.Background)
	vector.StrokeRect(m.Image, 0, 0, float32(width), float32(height), 2, theme.Colors.Off, false)

	tl := NewLabel(width, 35, "VCA", theme.Fonts.Title)
	tl.SetPosition(0, 0)
	m.Append(tl)

	// Components
	sl := NewSlider(width-20, 200)
	sl.SetPosition(10, 35)
	sl.SetRange(0, 1)
	sl.SetValue(0.5)
	m.Append(sl)

	vector.StrokeLine(m.Image, float32(width/2), float32(237), float32(width/2), 243, 1, theme.Colors.Off, false)

	cvPl := NewPlug()
	cvPl.SetPosition(0, 245)
	m.Append(cvPl)
	cvPl.HCenter()

	lb := NewLabel(width, 10, "CV", theme.Fonts.Small)
	lb.SetPosition(0, 287)
	m.Append(lb)

	inLb := NewLabel(width-20, 10, "IN", theme.Fonts.Small)
	inLb.SetPosition(0, height-122)
	m.Append(inLb)
	inLb.HCenter()

	inPl := NewPlug()
	inPl.SetPosition(0, height-110)
	m.Append(inPl)
	inPl.HCenter()

	vector.StrokeLine(m.Image, float32(width/2), float32(height-60), float32(width/2), float32(height-90), 1, theme.Colors.Off, false)

	iv := NewInverted(width-20, 48)
	iv.SetPosition(10, height-58)
	m.Append(iv)

	outPl := NewPlug()
	outPl.SetPosition(0, 5)
	iv.Append(outPl)
	outPl.HCenter()

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
