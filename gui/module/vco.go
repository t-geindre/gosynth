package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
)

type VCO struct {
	*Module
}

func NewVCO() *Module {
	v := &VCO{
		Module: NewModule("VCO", 2),
	}

	v.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	outPlug := connection.NewPlug(connection.PlugDirectionOut)
	v.Append(outPlug)

	return v.Module
}
