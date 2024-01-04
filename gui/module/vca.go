package module

import (
	"gosynth/gui/component"
	"gosynth/gui/widget"
)

type VCA struct {
	*Module
}

func NewVCA() *VCA {
	v := &VCA{}
	v.Module = NewModule("VCA", 1, v)

	slider := widget.NewSlider(0, 1, 25)
	slider.GetLayout().SetFill(100)
	//slider.GetLayout().GetMargin().SetBottom(5)
	slider.SetValue(0.5)
	v.Append(slider)

	vLine := widget.NewLine(false, float32(2))
	//vLine.GetLayout().GetWantedSize().SetHeight(10)
	//vLine.GetLayout().GetMargin().SetBottom(2)
	v.Append(vLine)

	cvInPlug := widget.NewPlug()
	//cvInPlug.GetLayout().GetMargin().SetBottom(5)
	v.Append(cvInPlug)
	v.Append(widget.NewText("CV", widget.TextSizeSmall))

	hLine := widget.NewLine(true, float32(2))
	//hLine.GetLayout().GetWantedSize().SetHeight(10)
	//hLine.GetLayout().GetMargin().Set(5, 5, 0, 0)
	v.Append(hLine)

	v.Append(widget.NewText("IN", widget.TextSizeSmall))

	inPlug := widget.NewPlug()
	//inPlug.GetLayout().GetMargin().SetTop(2)
	v.Append(inPlug)

	vLineOut := widget.NewLine(false, float32(2))
	//vLineOut.GetLayout().GetWantedSize().SetHeight(10)
	//vLineOut.GetLayout().GetMargin().SetBottom(2)
	v.Append(vLineOut)

	outContainer := component.NewContainer()
	outContainer.GetLayout().SetWantedSize(0, 55)
	//outContainer.GetLayout().GetPadding().Set(5, 5, 5, 5)
	v.Append(outContainer)

	outLabel := widget.NewText("OUT", widget.TextSizeSmall)
	outContainer.Append(outLabel)

	outPlug := widget.NewPlug()
	outPlug.SetInverted(true)
	//outPlug.GetLayout().GetMargin().SetTop(2)
	outContainer.Append(outPlug)

	return v
}
