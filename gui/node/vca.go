package node

import (
	"gosynth/gui/theme"
)

type VCA struct {
	*Module
}

func NewVCA() *VCA {
	v := &VCA{}
	width, height := 65, 500
	v.Module = NewModule(width, height, v)

	slider := NewSlider()
	slider.SetRange(0, 1)
	slider.SetValue(0.5)
	v.AppendWithOptions(slider, NewAppendOptions().HorizontallyFill(100).VerticallyFill(100))

	lineToCv := NewLine(10, 1, LineOrientationVertical)
	v.AppendWithOptions(lineToCv, NewAppendOptions().HorizontallyCentered())

	cvPlug := NewPlug()
	v.AppendWithOptions(cvPlug, NewAppendOptions().HorizontallyCentered())

	cvLabel := NewLabel("CV", theme.Fonts.Small)
	v.AppendWithOptions(cvLabel, NewAppendOptions().HorizontallyCentered().Margins(3, 0, 0, 0))

	separatorLine := NewLine(10, 1, LineOrientationHorizontal)
	v.AppendWithOptions(
		separatorLine,
		NewAppendOptions().
			HorizontallyCentered().
			HorizontallyFill(100).
			Margins(20, 20, 10, 10),
	)

	inLabel := NewLabel("IN", theme.Fonts.Small)
	v.AppendWithOptions(inLabel, NewAppendOptions().HorizontallyCentered().Margins(0, 3, 0, 0))

	inPlug := NewPlug()
	v.AppendWithOptions(inPlug, NewAppendOptions().HorizontallyCentered())

	inOutLine := NewLine(10, 1, LineOrientationVertical)
	v.AppendWithOptions(inOutLine, NewAppendOptions().HorizontallyCentered())

	outPlug := NewPlug()
	outPlugContainer := NewContainer(outPlug.GetOuterWidth()+10, outPlug.GetOuterHeight()+10)
	outPlugContainer.SetInverted(true)
	v.AppendWithOptions(outPlugContainer, NewAppendOptions().HorizontallyCentered().Padding(5))

	outPlugContainer.AppendWithOptions(outPlug, NewAppendOptions().HorizontallyCentered())

	return v
}
