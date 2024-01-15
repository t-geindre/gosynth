package module

import (
	"github.com/gopxl/beep"
	"gosynth/math/normalizer"
	"math"
)

type FilterMode uint8

const (
	PassFilterModeLow FilterMode = iota
	PassFilterModeHigh
)

type VCF struct {
	*Module
	sample float64
	buffer float64
	alpha  float64
	mode   FilterMode
}

func NewVCF(sr beep.SampleRate) *VCF {
	p := &VCF{}
	p.Module = NewModule(sr, p)

	p.AddInput(PortIn)
	p.AddInput(PortInCV)
	p.AddOutput(PortOut)

	return p
}

func (p *VCF) SetMode(mode FilterMode) {
	p.mode = mode
}

func (p *VCF) Write(port Port, value float64) {
	switch port {
	case PortIn:
		p.sample += value
	case PortInCV:
		tan := math.Tan(math.Pi * normalizer.CvToFrequency(value) / float64(p.GetSampleRate()))
		p.alpha = (tan - 1) / (tan + 1)
	}

	p.Module.Write(port, value)
}

func (p *VCF) Update() {
	p.Module.Update()

	pass := p.alpha*p.sample + p.buffer
	p.buffer = p.sample - p.alpha*pass

	if p.mode == PassFilterModeHigh {
		pass *= -1
	}

	p.ConnectionWrite(PortOut, p.sample+pass)
	p.sample = 0
}
