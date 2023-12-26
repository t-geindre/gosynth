package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Gain struct {
	Module
	Gain       float64
	MasterGain float64
	Sample     float64
}

func (g *Gain) Init(rate beep.SampleRate) {
	g.Module.Init(rate)

	g.Gain = 1
	g.MasterGain = 1

	g.AddInput("Gain", PortInGain)
	g.AddInput("in", PortIn)
	g.AddOutput("out", PortOut)
}

func (g *Gain) SetMasterGain(gain float64) {
	g.MasterGain = gain
}

func (g *Gain) Write(port Port, value float64) {
	switch port {
	case PortInGain:
		g.Gain = (value + 1) / 2 // Normalize to 0-1
	case PortIn:
		g.Sample += value
	}
}

func (g *Gain) Update(_ time.Duration) {
	g.ConnectionWrite(PortOut, g.Sample*g.Gain*g.MasterGain)
	g.Sample = 0
}
