package widget

import (
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
	t := &Label{
		Text: component.NewText(text, theme.Fonts.Small, theme.Colors.Text, theme.Colors.Background),
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
