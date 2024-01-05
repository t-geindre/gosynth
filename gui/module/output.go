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
		Module: NewModule("OUT", 1),
	}

	v.Append(widget.NewLabel("L/M", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionIn, rack, audio.PortInL))

	v.Append(widget.NewLabel("R", widget.LabelPositionTop))
	v.Append(connection.NewPlug(connection.PlugDirectionIn, rack, audio.PortInR))

	return v.Module
}
