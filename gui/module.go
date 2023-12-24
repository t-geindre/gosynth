package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

type Module struct {
	Node
	Layout *ebiten.Image
}

func NewModule() *Module {
	m := &Module{}
	m.Node = *NewNode(300, 500, m)
	m.Layout = ebiten.NewImage(300, 500)
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
