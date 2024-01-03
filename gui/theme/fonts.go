package theme

import (
	_ "embed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

//go:embed LEMONMILK-Medium.otf
var fontLemonMilkMedium []byte

type fontList struct {
	Title  font.Face
	Large  font.Face
	Medium font.Face
	Small  font.Face
}

var Fonts fontList

func init() {
	Fonts = fontList{
		Title:  getFontFace(fontLemonMilkMedium, 18),
		Large:  getFontFace(fontLemonMilkMedium, 14),
		Medium: getFontFace(fontLemonMilkMedium, 12),
		Small:  getFontFace(fontLemonMilkMedium, 10),
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
