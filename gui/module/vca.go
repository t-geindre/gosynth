package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type VCA struct {
	*Module
}

func NewVCA(vca *audio.VCA) *Module {
	v := &VCA{
		Module: NewModule("VCA", 1),
	}

	slider := widget.NewSlider(0, 1, 25)
	slider.GetLayout().SetFill(80)
	slider.SetValue(0.5)
	v.Append(slider)

	vLine := widget.NewLine(false, float32(2))
	vLine.GetLayout().SetFill(5)
	v.Append(vLine)

	cvInPlug := connection.NewPlug(connection.PlugDirectionIn)
	cvInPlug.Bind(vca, audio.PortCvIn)
	v.Append(cvInPlug)
	v.Append(widget.NewLabel("CV", widget.LabelPositionBottom))

	hLine := widget.NewLine(true, float32(2))
	hLine.GetLayout().SetFill(10)
	v.Append(hLine)

	v.Append(widget.NewLabel("IN", widget.LabelPositionTop))
	inPlug := connection.NewPlug(connection.PlugDirectionIn)
	inPlug.Bind(vca, audio.PortIn)
	v.Append(inPlug)

	vLineOut := widget.NewLine(false, float32(2))
	vLineOut.GetLayout().SetFill(5)
	v.Append(vLineOut)

	v.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	outPlug := connection.NewPlug(connection.PlugDirectionOut)
	outPlug.Bind(vca, audio.PortOut)
	v.Append(outPlug)

	return v.Module
}
