package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type VCO struct {
	*Module
}

func NewVCO(vco audio.IModule) *Module {
	v := &VCO{
		Module: NewModule("VCO", 1),
	}

	v.Append(widget.NewLabel("FRQ", widget.LabelPositionTop))
	v.Append(widget.NewKnob(vco, audio.PortInVOct))

	v.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionIn, vco, audio.PortInVOct))

	l := widget.NewLine(true, float32(2))
	l.GetLayout().SetFill(25)
	v.Append(l)

	v.Append(widget.NewLabel("SIN", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, vco, audio.PortOutSin))

	v.Append(widget.NewLabel("TRI", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, vco, audio.PortOutTriangle))

	v.Append(widget.NewLabel("SAW", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, vco, audio.PortOutSaw))

	v.Append(widget.NewLabel("SQR", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, vco, audio.PortOutSquare))

	return v.Module
}
