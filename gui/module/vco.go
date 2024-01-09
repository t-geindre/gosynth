package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/layout"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type VCO struct {
	*Module
}

func NewVCO(vco audio.IModule) *Module {
	v := &VCO{
		Module: NewModule("VCO", 3),
	}

	// MAIN CONTROLS
	mainControls := component.NewContainer()

	mainControls.Append(component.NewFiller(50))
	mainControls.Append(widget.NewLargeLabel("Frequency", widget.LabelPositionTop))
	mainControls.Append(widget.NewLargeKnob(vco, audio.PortInVOct, 0))

	mainControls.Append(component.NewFiller(50))
	mainControls.Append(widget.NewMediumLabel("Pulse width", widget.LabelPositionTop))
	mainControls.Append(widget.NewMediumKnob(vco, audio.PortInPw, 0))

	mainControls.GetLayout().SetFill(55)
	v.Append(mainControls)

	// SEPARATION LINE
	l := widget.NewLine(true, float32(2))
	l.GetLayout().SetFill(8)
	v.Append(l)

	// INPUTS
	inputs := component.NewContainer()

	inputs.GetLayout().SetContentOrientation(layout.Horizontal)
	inputs.Append(v.addInput("FM", vco, audio.PortInFm, audio.PortInFmFact, 25))
	inputs.Append(v.addInput("V/OCT", vco, audio.PortInVOct, audio.PortInOctShift, 25))
	inputs.Append(v.addInput("SYNC", vco, audio.PortInSync, audio.PortInPhaseShift, 25))
	inputs.Append(v.addInput("PWM", vco, audio.PortInPwm, audio.PortInPwmFact, 25))

	inputs.GetLayout().SetFill(25)
	inputs.GetLayout().SetPadding(0, 10, 0, 0)
	v.Append(inputs)

	// OUTPUTS
	outputs := component.NewContainer()

	outputs.GetLayout().SetContentOrientation(layout.Horizontal)

	outputs.Append(v.addOutput("SIN", vco, audio.PortOutSin, 25))
	outputs.Append(v.addOutput("TRI", vco, audio.PortOutTriangle, 25))
	outputs.Append(v.addOutput("SAW", vco, audio.PortOutSaw, 25))
	outputs.Append(v.addOutput("SQR", vco, audio.PortOutSquare, 25))

	outputs.GetLayout().SetFill(12)
	v.Append(outputs)

	return v.Module
}

func (v *VCO) addInput(label string, module audio.IModule, pPort, kPort audio.Port, fill float64) component.IComponent {
	c := component.NewContainer()
	c.GetLayout().SetFill(fill)
	c.Append(widget.NewKnob(module, kPort, 0))
	l := widget.NewLine(false, 2)
	l.GetLayout().SetFill(100)
	c.Append(l)

	c.GetLayout().SetContentOrientation(layout.Vertical)
	c.Append(widget.NewLabel(label, widget.LabelPositionTop))
	c.Append(connection.NewPlug(connection.PlugDirectionIn, module, pPort))

	return c
}

func (v *VCO) addOutput(label string, module audio.IModule, port audio.Port, fill float64) component.IComponent {
	c := component.NewContainer()
	c.Append(component.NewFiller(100))
	c.GetLayout().SetFill(fill)

	c.GetLayout().SetContentOrientation(layout.Vertical)
	c.Append(widget.NewLabel(label, widget.LabelPositionTop))
	c.Append(connection.NewPlug(connection.PlugDirectionOut, module, port))

	return c
}
