package demo

import (
	"image/color"
	"math/rand"
)

func randomColor() color.RGBA {
	c := color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}

	return c
}

func colorInverse(c color.RGBA) color.RGBA {
	return color.RGBA{
		R: 255 - c.R,
		G: 255 - c.G,
		B: 255 - c.B,
		A: 255,
	}
}
