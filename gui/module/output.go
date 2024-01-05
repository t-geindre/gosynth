package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type Output struct {
	*Module
}

func NewOutput(rack *audio.Rack) *Module {
	v := &Output{
		Module: NewModule("Output", 2),
	}

	v.Append(widget.NewLabel("L/Mono", widget.LabelPositionTop))
	inLPlug := connection.NewPlug(connection.PlugDirectionIn)
	inLPlug.Bind(rack, audio.PortInL)
	v.Append(inLPlug)

	v.Append(widget.NewLabel("R", widget.LabelPositionTop))
	inRPlug := connection.NewPlug(connection.PlugDirectionIn)
	inRPlug.Bind(rack, audio.PortInR)
	v.Append(inRPlug)

	return v.Module
}
