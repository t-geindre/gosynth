package module

import (
	"gosynth/gui/theme"
	widget2 "gosynth/gui/widget"
)

type VCA struct {
	*Module
}

func NewVCA() *VCA {
	v := &VCA{}
	v.Module = NewModule("VCA", 1, v)

	slider := widget2.NewSlider(0, 1, 25)
	slider.GetLayout().SetFill(100)
	slider.GetLayout().GetMargin().SetBottom(5)
	slider.SetValue(0.5)
	v.Append(slider)

	vLine := widget2.NewLine(false, float32(2))
	vLine.GetLayout().GetWantedSize().SetHeight(10)
	vLine.GetLayout().GetMargin().SetBottom(2)
	v.Append(vLine)

	cvInPlug := widget2.NewPlug()
	cvInPlug.GetLayout().GetMargin().SetBottom(5)
	v.Append(cvInPlug)
	v.Append(widget2.NewText("CV", theme.Fonts.Small))

	hLine := widget2.NewLine(true, float32(2))
	hLine.GetLayout().GetWantedSize().SetHeight(10)
	hLine.GetLayout().GetMargin().Set(5, 5, 0, 0)
	v.Append(hLine)

	v.Append(widget2.NewText("IN", theme.Fonts.Small))

	inPlug := widget2.NewPlug()
	inPlug.GetLayout().GetMargin().SetTop(2)
	v.Append(inPlug)

	vLineOut := widget2.NewLine(false, float32(2))
	vLineOut.GetLayout().GetWantedSize().SetHeight(10)
	vLineOut.GetLayout().GetMargin().SetBottom(2)
	v.Append(vLineOut)

	outContainer := widget2.NewContainer()
	outContainer.SetInverted(true)
	outContainer.GetLayout().GetWantedSize().SetHeight(55)
	outContainer.GetLayout().GetPadding().Set(5, 5, 5, 5)
	v.Append(outContainer)

	outLabel := widget2.NewText("OUT", theme.Fonts.Small)
	outLabel.SetInverted(true)
	outContainer.Append(outLabel)

	outPlug := widget2.NewPlug()
	outPlug.SetInverted(true)
	outPlug.GetLayout().GetMargin().SetTop(2)
	outContainer.Append(outPlug)

	return v
}
