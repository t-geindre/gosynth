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
	time       time.Duration
}

func (g *VCA) Init(rate beep.SampleRate) {
	g.Module = &Module{}
	g.Module.Init(rate, g)
	g.Write(PortCvIn, 1)
}

func (g *VCA) Write(port Port, value float64) {
	switch port {
	case PortCvIn:
		g.Gain = (value + 1) / 2
	case PortIn:
		g.Sample += value
	}
	g.Module.Write(port, value)
}

func (g *VCA) Update(t time.Duration) {
	g.time = t
	g.Module.Update(t)
	g.ConnectionWrite(PortOut, g.Sample*g.Gain)
	g.Sample = 0
}
