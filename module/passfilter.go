package module

import (
	"github.com/gopxl/beep"
	"math"
)

type FilterMode uint8

const (
	PassFilterModeLow FilterMode = iota
	PassFilterModeHigh
)

type PassFilter struct {
	*Module
	sample float64
	buffer float64
	alpha  float64
	mode   FilterMode
}

func NewPassFilter(sr beep.SampleRate) *PassFilter {
	p := &PassFilter{}
	p.Module = NewModule(sr, p)

	p.AddInput(PortIn)

	return p
}
func (p *PassFilter) Init(SampleRate beep.SampleRate) {
}

func (p *PassFilter) SetCutOff(cutoff float64) {
	tan := math.Tan(math.Pi * cutoff / float64(p.GetSampleRate()))
	p.alpha = (tan - 1) / (tan + 1)
}

func (p *PassFilter) SetMode(mode FilterMode) {
	p.mode = mode
}

func (p *PassFilter) Write(port Port, value float64) {
	switch port {
	case PortIn:
		p.sample += value
	}

	p.Module.Write(port, value)
}

func (p *PassFilter) Update() {
	p.Module.Update()

	pass := p.alpha*p.sample + p.buffer
	p.buffer = p.sample - p.alpha*pass

	if p.mode == PassFilterModeHigh {
		pass *= -1
	}

	p.ConnectionWrite(PortOut, p.sample+pass)
	p.sample = 0
}
