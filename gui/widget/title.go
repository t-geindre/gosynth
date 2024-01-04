package widget

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/theme"
)

type titlePosition uint8

const (
	TitlePositionLeft titlePosition = iota
	TitlePositionRight
	TitlePositionTop
	TitlePositionBottom
	TitlePositionCenter
)

type Title struct {
	*component.Text
}

func NewTitle(text string, position titlePosition) *Title {
	t := &Title{
		Text: component.NewText(text, theme.Fonts.Title, theme.Colors.Text, theme.Colors.Background),
	}

	switch position {
	case TitlePositionLeft:
		t.GetLayout().SetMargin(0, 0, 10, 0)
	case TitlePositionRight:
		t.GetLayout().SetMargin(0, 0, 0, 10)
	case TitlePositionTop:
		t.GetLayout().SetMargin(0, 10, 0, 0)
	case TitlePositionBottom:
		t.GetLayout().SetMargin(10, 0, 0, 0)
	}

	return t
}
