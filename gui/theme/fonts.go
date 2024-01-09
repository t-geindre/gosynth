package theme

import (
	_ "embed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"os"
)

type fontList struct {
	Title  font.Face
	Large  font.Face
	Medium font.Face
	Small  font.Face
}

var Fonts fontList

func init() {
	fontMedium, err := os.ReadFile("assets/fonts/LEMONMILK-Medium.otf")

	if err != nil {
		panic(err)
	}

	fontLight, err := os.ReadFile("assets/fonts/LEMONMILK-Light.otf")

	if err != nil {
		panic(err)
	}

	Fonts = fontList{
		Title:  getFontFace(fontMedium, 18),
		Large:  getFontFace(fontLight, 18),
		Medium: getFontFace(fontLight, 12),
		Small:  getFontFace(fontLight, 10),
	}
}

func getFontFace(src []byte, size float64) font.Face {
	ft, err := sfnt.Parse(src)

	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}

	return face
}
