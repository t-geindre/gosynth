package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/layout"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
)

type Delay struct {
	*Module
}

func NewDelay() *Module {
	d := &Delay{
		Module: NewModule("Delay", 2),
	}

	d.appendControlLine("TIME", 70/4)
	d.appendControlLine("FDBK", 70/4)
	d.appendControlLine("COLOR", 70/4)
	d.appendControlLine("MIX", 70/4)

	hr := widget.NewLine(true, float32(2))
	hr.GetLayout().SetFill(5)
	d.Append(hr)

	cb := component.NewContainer()
	cb.GetLayout().SetFill(25)

	cb.Append(widget.NewLabel("IN", widget.LabelPositionTop))
	cb.Append(connection.NewPlug(connection.PlugDirectionIn))
	line := widget.NewLine(false, float32(2))
	line.GetLayout().SetFill(100)
	cb.Append(line)
	cb.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	cb.Append(connection.NewPlug(connection.PlugDirectionOut))
	d.Append(cb)

	return d.Module
}

func (d *Delay) appendControlLine(label string, fill float64) {
	c := component.NewContainer()
	c.GetLayout().SetContentOrientation(layout.Horizontal)
	c.GetLayout().SetFill(fill)
	c.GetLayout().SetPadding(10, 0, 0, 0)
	d.Append(c)

	pcc := component.NewContainer()
	pcc.GetLayout().SetFill(80 / 2)
	c.Append(pcc)

	pcc.Append(widget.NewLabel("CV", widget.LabelPositionTop))
	pcc.Append(connection.NewPlug(connection.PlugDirectionIn))

	hr := widget.NewLine(true, float32(2))
	hr.GetLayout().SetFill(20)
	c.Append(hr)

	kcc := component.NewContainer()
	kcc.GetLayout().SetFill(80 / 2)
	c.Append(kcc)

	kcc.Append(widget.NewLabel(label, widget.LabelPositionTop))
	kcc.Append(widget.NewKnob())
}
