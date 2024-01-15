package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

type VCF struct {
	*Module
}

func NewVCF(audioVcf *audio.VCF) *VCF {
	vcf := &VCF{
		Module: NewModule("VCF", 1),
	}

	vcf.Append(component.NewFiller(25))
	vcf.Append(widget.NewLabel("Cutoff", widget.LabelPositionTop))
	vcf.Append(widget.NewKnob(audioVcf, audio.PortInCV, 0))

	vcf.Append(component.NewFiller(25))
	vcf.Append(widget.NewLabel("IN", widget.LabelPositionTop))
	vcf.Append(connection.NewPlug(connection.PlugDirectionIn, audioVcf, audio.PortIn))

	vcf.Append(component.NewFiller(25))
	vcf.Append(widget.NewLabel("OUT", widget.LabelPositionTop))
	vcf.Append(connection.NewPlug(connection.PlugDirectionOut, audioVcf, audio.PortOut))

	vcf.Append(component.NewFiller(25))

	return vcf
}
