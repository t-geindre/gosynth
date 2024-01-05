package theme

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Images imageList

type imageList struct {
	Plug         *ebiten.Image
	PlugInverted *ebiten.Image
	Logo         *ebiten.Image
	Knob         *ebiten.Image
}

func init() {
	Images = imageList{
		Plug:         loadImage("assets/images/plug.png"),
		PlugInverted: loadImage("assets/images/plug-inverted.png"),
		Logo:         loadImage("assets/images/512px-Go_Logo_Blue.svg.png"),
		Knob:         loadImage("assets/images/knob.png"),
	}
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		panic(err)
	}

	return img
}
