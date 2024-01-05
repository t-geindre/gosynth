package connection

import "image/color"

var colors = [...]color.RGBA{
	{R: 0, G: 20, B: 219, A: 255},
	{R: 219, G: 201, B: 0, A: 255},
	{R: 219, G: 85, B: 0, A: 255},
	{R: 0, G: 219, B: 44, A: 255},
	{R: 219, G: 0, B: 207, A: 255},
	{R: 219, G: 0, B: 0, A: 255},
}

var index = 0

func CycleColor() color.RGBA {
	index = (index + 1) % len(colors)
	return colors[index]
}
