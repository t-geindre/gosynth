package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type VCO struct {
	*Module
}

func NewVCO(vco *audio.Oscillator) *Module {
	v := &VCO{
		Module: NewModule("VCO", 2),
	}

	v.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	outPlug := connection.NewPlug(connection.PlugDirectionOut)
	outPlug.Bind(vco, audio.PortOut)
	v.Append(outPlug)

	return v.Module
}
