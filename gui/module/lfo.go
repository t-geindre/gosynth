package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type LFO struct {
	*Module
}

func NewLFO(lfo *audio.LFO) *Module {
	v := &LFO{
		Module: NewModule("LFO", 1),
	}

	v.Append(widget.NewLabel("FRQ", widget.LabelPositionTop))
	v.Append(widget.NewKnob(lfo, audio.PortInVOct))

	v.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionIn, lfo, audio.PortInVOct))

	l := widget.NewLine(true, float32(2))
	l.GetLayout().SetFill(25)
	v.Append(l)

	v.Append(widget.NewLabel("SIN", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, lfo, audio.PortOutSin))

	v.Append(widget.NewLabel("TRI", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, lfo, audio.PortOutTriangle))

	v.Append(widget.NewLabel("SAW", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, lfo, audio.PortOutSaw))

	v.Append(widget.NewLabel("SQR", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionOut, lfo, audio.PortOutSquare))

	return v.Module
}
