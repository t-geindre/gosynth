package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type Clock struct {
	*Module
}

func NewClock(audioClk audio.IModule) *Clock {
	c := &Clock{
		Module: NewModule("CLK", 1),
	}

	c.Append(component.NewFiller(100 / 6))
	c.Append(widget.NewLabel("Time", widget.LabelPositionTop))
	c.Append(widget.NewKnob(audioClk, audio.PortInCV, 0))

	c.Append(component.NewFiller(100 / 6))
	c.Append(widget.NewLabel("1/1", widget.LabelPositionTop))
	c.Append(connection.NewPlug(connection.PlugDirectionOut, audioClk, audio.PortOut1))

	c.Append(component.NewFiller(100 / 6))
	c.Append(widget.NewLabel("1/2", widget.LabelPositionTop))
	c.Append(connection.NewPlug(connection.PlugDirectionOut, audioClk, audio.PortOut2))

	c.Append(component.NewFiller(100 / 6))
	c.Append(widget.NewLabel("1/4", widget.LabelPositionTop))
	c.Append(connection.NewPlug(connection.PlugDirectionOut, audioClk, audio.PortOut3))

	c.Append(component.NewFiller(100 / 6))
	c.Append(widget.NewLabel("1/8", widget.LabelPositionTop))
	c.Append(connection.NewPlug(connection.PlugDirectionOut, audioClk, audio.PortOut4))

	c.Append(component.NewFiller(100 / 6))

	return c
}
