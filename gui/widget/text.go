package widget

import (
	"golang.org/x/image/font"
	"gosynth/gui/component"
	"gosynth/gui/theme"
)

type TextSize uint8

const (
	TextSizeSmall TextSize = iota
	TextSizeMedium
	TextSizeLarge
	TextSizeTitle
)

type Text struct {
	*component.Text
}

func NewText(text string, size TextSize) *Text {
	var fontFace font.Face

	switch size {
	case TextSizeSmall:
		fontFace = theme.Fonts.Small
	case TextSizeMedium:
		fontFace = theme.Fonts.Medium
	case TextSizeLarge:
		fontFace = theme.Fonts.Large
	case TextSizeTitle:
		fontFace = theme.Fonts.Title
	default:
		panic("invalid text size")
	}

	t := &Text{
		Text: component.NewText(text, fontFace, theme.Colors.Text, theme.Colors.Background),
	}

	return t
}
