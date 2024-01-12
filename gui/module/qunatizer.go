package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type Quantizer struct {
	*Module
}

func NewQuantizer(audioQtz *audio.Quantizer) *Quantizer {
	q := &Quantizer{}
	q.Module = NewModule("QTZ", 1)

	q.Append(component.NewFiller(30))
	q.Append(widget.NewLabel("IN", widget.LabelPositionTop))
	q.Append(connection.NewPlug(connection.PlugDirectionIn, audioQtz, audio.PortIn))
	q.Append(component.NewFiller(30))
	q.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	q.Append(connection.NewPlug(connection.PlugDirectionOut, audioQtz, audio.PortOut))
	q.Append(component.NewFiller(30))

	return q
}
