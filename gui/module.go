package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

type Module struct {
	Node
	Layout                 *ebiten.Image
	MouseLDown             bool
	LastMouseX, LastMouseY int
}

func NewModule(width, height int) *Module {
	m := &Module{}
	m.Node = *NewNode(width, height, m)
	m.Layout = ebiten.NewImage(width, height)
	col := color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
	m.Layout.Fill(col)

	return m
}

func (m *Module) Draw(dest *ebiten.Image) {
	m.Node.Draw(dest)
	dest.DrawImage(m.Layout, m.Options)
}

func (m *Module) Update() error {
	if m.MouseLDown {
		x, y := ebiten.CursorPosition()
		m.MoveBy(x-m.LastMouseX, y-m.LastMouseY)
		m.LastMouseX, m.LastMouseY = x, y
	}

	return m.Node.Update()
}
func (m *Module) MouseLeftDown() {
	if m.GetParent() != nil {
		// Todo click on sub child wont move module to front
		m.GetParent().MoveFront(m)
	}
	m.MouseLDown = true
	m.LastMouseX, m.LastMouseY = ebiten.CursorPosition()
	m.Options.ColorScale.ScaleAlpha(0.5)
}

func (m *Module) MouseLeftUp() {
	m.Options.ColorScale.Reset()
	m.MouseLDown = false
}
