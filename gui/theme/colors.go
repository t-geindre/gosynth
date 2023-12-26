package theme

import "image/color"

type colorsList struct {
	Text               *color.RGBA
	TextInverted       *color.RGBA
	Background         *color.RGBA
	BackgroundInverted *color.RGBA
	On                 *color.RGBA
	Off                *color.RGBA
}

var Colors colorsList

func init() {
	Colors = colorsList{
		Text:               &color.RGBA{R: 3, G: 3, B: 3, A: 255},
		TextInverted:       &color.RGBA{R: 232, G: 232, B: 232, A: 255},
		Background:         &color.RGBA{R: 230, G: 230, B: 230, A: 255},
		BackgroundInverted: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		On:                 &color.RGBA{R: 33, G: 95, B: 169, A: 255},
		Off:                &color.RGBA{R: 80, G: 80, B: 80, A: 255},
	}
}
