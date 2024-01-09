package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type Mixer struct {
	*Module
}

func NewMixer(mixer audio.IModule) *Mixer {
	m := &Mixer{
		Module: NewModule("MXR", 1),
	}

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("IN 1", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionIn, mixer, audio.PortIn1))

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("IN 2", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionIn, mixer, audio.PortIn2))

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("IN 3", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionIn, mixer, audio.PortIn3))

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("IN 4", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionIn, mixer, audio.PortIn4))

	l := widget.NewLine(true, 2)
	l.GetLayout().SetFill(100 / 7)
	m.Append(l)

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionOut, mixer, audio.PortOut))
	m.Append(component.NewFiller(100 / 7))

	return m
}
