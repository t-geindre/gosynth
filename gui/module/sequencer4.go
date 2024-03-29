package module

import (
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type Sequencer4 struct {
	*Module
}

func NewSequencer4(seq audio.IModule) *Sequencer4 {
	s := &Sequencer4{}
	s.Module = NewModule("SQR", 1)

	s.Append(widget.NewLabel("CLK", widget.LabelPositionTop))
	s.Append(connection.NewPlug(connection.PlugDirectionIn, seq, audio.PortInTrigger))

	l := widget.NewLine(true, 1)
	l.GetLayout().SetFill(25)
	s.Append(l)

	s.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	s.Append(widget.NewKnob(seq, audio.PortIn1, 0))

	s.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	s.Append(widget.NewKnob(seq, audio.PortIn2, 0))

	s.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	s.Append(widget.NewKnob(seq, audio.PortIn3, 0))

	s.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	s.Append(widget.NewKnob(seq, audio.PortIn4, 0))

	l2 := widget.NewLine(true, 1)
	l2.GetLayout().SetFill(25)
	s.Append(l2)

	s.Append(widget.NewLabel("V/OCT", widget.LabelPositionTop))
	s.Append(connection.NewPlug(connection.PlugDirectionOut, seq, audio.PortOutCv))

	s.Append(widget.NewLabel("GATE", widget.LabelPositionTop))
	s.Append(connection.NewPlug(connection.PlugDirectionOut, seq, audio.PortOutGate))

	return s
}
