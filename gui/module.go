package gui

import (
	"image/color"
	"math/rand"
)

type Module struct {
	*Node
}

func NewModule() *Module {
	m := &Module{}
	m.Node = NewNode(nil, 300, 500)
	col := color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
	m.Node.Image.Fill(col)

	return m
}
