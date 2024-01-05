package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
)

type Output struct {
	*Module
}

func NewOutput() *Module {
	v := &Output{
		Module: NewModule("Output", 2),
	}

	v.Append(widget.NewLabel("L/Mono", widget.LabelPositionTop))
	inLPlug := connection.NewPlug(connection.PlugDirectionIn)
	v.Append(inLPlug)

	v.Append(widget.NewLabel("R", widget.LabelPositionTop))
	inRPlug := connection.NewPlug(connection.PlugDirectionIn)
	v.Append(inRPlug)

	return v.Module
}
