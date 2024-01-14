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
		Module: NewModule("CLOCK", 2),
	}

	// Top
	top := component.NewVContainter()
	top.GetLayout().SetFill(20)
	c.Append(top)

	top.Append(component.NewFiller(50))
	top.Append(widget.NewLabel("Time", widget.LabelPositionTop))
	top.Append(widget.NewMediumKnob(audioClk, audio.PortInCV, 0))
	top.Append(component.NewFiller(50))

	// Separator
	sep := widget.NewHLine(1)
	sep.GetLayout().SetFill(4)
	c.Append(sep)

	// Clocks
	bottom := component.NewVContainter()
	bottom.GetLayout().SetFill(62)
	c.Append(bottom)

	for _, port := range [...][2]audio.Port{
		{audio.PortInCV1, audio.PortOut1},
		{audio.PortInCV2, audio.PortOut2},
		{audio.PortInCV3, audio.PortOut3},
		{audio.PortInCV4, audio.PortOut4},
		{audio.PortInCV5, audio.PortOut5},
		{audio.PortInCV6, audio.PortOut6},
	} {
		line := component.NewHContainter()
		line.Append(widget.NewKnob(audioClk, port[0], 0))
		con := widget.NewHLine(1)
		con.GetLayout().SetFill(100)
		line.Append(con)
		line.Append(connection.NewPlug(connection.PlugDirectionOut, audioClk, port[1]))
		line.GetLayout().SetFill(100 / 7)
		bottom.Append(line)
	}

	// Separator
	sSep := widget.NewHLine(1)
	sSep.GetLayout().SetFill(4)
	c.Append(sSep)

	// Reset
	// Todo add a reset button/output, and a reset input on the sequencer
	reset := component.NewHContainter()
	reset.GetLayout().SetFill(10)
	c.Append(reset)

	return c
}
