package module

import (
	"github.com/gopxl/beep"
	"time"
)

type VCA struct {
	*Module
	Gain       float64
	MasterGain float64
	Sample     float64
}

func (g *VCA) Init(rate beep.SampleRate) {
	g.Module = &Module{}
	g.Module.Init(rate, g)

	g.Gain = 1
}

func (g *VCA) SetGain(gain float64) {
	g.Gain = gain
}

func (g *VCA) Write(port Port, value float64) {
	switch port {
	case PortCvIn:
		g.SetGain((value + 1) / 2) // Normalize to 0-1
	case PortIn:
		g.Sample += value
	}
}

func (g *VCA) GetGain() float64 {
	return g.Gain*2 - 1
}

func (g *VCA) Update(t time.Duration) {
	g.Module.Update(t)
	g.ConnectionWrite(PortOut, g.Sample*g.Gain)
	g.Sample = 0
}
