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

	v.Append(connection.NewPlug(connection.PlugDirectionIn, vca, audio.PortCvIn))
	v.Append(widget.NewLabel("CV", widget.LabelPositionBottom))

	hLine := widget.NewLine(true, float32(2))
	hLine.GetLayout().SetFill(10)
	v.Append(hLine)

	v.Append(widget.NewLabel("IN", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionIn, vca, audio.PortIn))

	vLineOut := widget.NewLine(false, float32(2))
	vLineOut.GetLayout().SetFill(5)
	v.Append(vLineOut)

	v.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, vca, audio.PortOut))

	return v.Module
}
