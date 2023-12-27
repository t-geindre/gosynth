package module

import (
	"github.com/gopxl/beep"
	"time"
)

type VCA struct {
	Module
	Gain       float64
	MasterGain float64
	Sample     float64
}

func (g *VCA) Init(rate beep.SampleRate) {
	g.Module.Init(rate)

	g.Gain = 1

	g.AddInput("VCA", PortInGain)
	g.AddInput("in", PortIn)
	g.AddOutput("out", PortOut)
}

func (g *VCA) SetGain(gain float64) {
	g.Gain = gain
}

func (g *VCA) Write(port Port, value float64) {
	switch port {
	case PortInGain:
		g.SetGain((value + 1) / 2) // Normalize to 0-1
	case PortIn:
		g.Sample += value
	}
}

func (g *VCA) Update(t time.Duration) {
	g.Module.Update(t)
	g.ConnectionWrite(PortOut, g.Sample*g.Gain)
	g.Sample = 0
}