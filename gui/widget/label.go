package widget

import (
	"golang.org/x/image/font"
	"gosynth/gui-lib/component"
	"gosynth/gui/theme"
)

type labelPosition uint8

const (
	LabelPositionLeft labelPosition = iota
	LabelPositionRight
	LabelPositionTop
	LabelPositionBottom
)

type Label struct {
	*component.Text
}

func NewLabel(text string, position labelPosition) *Label {
	return newLabel(text, position, theme.Fonts.Small)
}

func NewMediumLabel(text string, position labelPosition) *Label {
	return newLabel(text, position, theme.Fonts.Medium)
}

func NewLargeLabel(text string, position labelPosition) *Label {
	return newLabel(text, position, theme.Fonts.Large)
}

func newLabel(text string, position labelPosition, font font.Face) *Label {
	t := &Label{
		Text: component.NewText(text, font, theme.Colors.Text, theme.Colors.Background),
	}

	switch position {
	case LabelPositionLeft:
		t.GetLayout().SetMargin(0, 0, 3, 0)
	case LabelPositionRight:
		t.GetLayout().SetMargin(0, 0, 0, 3)
	case LabelPositionTop:
		t.GetLayout().SetMargin(0, 3, 0, 0)
	case LabelPositionBottom:
		t.GetLayout().SetMargin(3, 0, 0, 0)
	}

	return t
}
