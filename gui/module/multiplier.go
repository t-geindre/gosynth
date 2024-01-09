package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type Multiplier struct {
	*Module
}

func NewMultiplier(multiplier audio.IModule) *Multiplier {
	m := &Multiplier{
		Module: NewModule("MLR", 1),
	}

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("IN", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionIn, multiplier, audio.PortIn))
	m.Append(component.NewFiller(100 / 7))

	l := widget.NewLine(true, 2)
	l.GetLayout().SetFill(100 / 7)
	m.Append(l)

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("OUT 1", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionOut, multiplier, audio.PortOut1))

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("OUT 2", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionOut, multiplier, audio.PortOut2))

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("OUT 3", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionOut, multiplier, audio.PortOut3))

	m.Append(component.NewFiller(100 / 7))
	m.Append(widget.NewLabel("OUT 4", widget.LabelPositionTop))
	m.Append(connection.NewPlug(connection.PlugDirectionOut, multiplier, audio.PortOut4))

	return m
}
