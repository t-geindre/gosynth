package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Gain struct {
	Module
	gain   float64
	Sample float64
}

func (g *Gain) Init(rate beep.SampleRate) {
	g.Module.Init(rate)
	g.AddInput("gain", PortInGain)
	g.AddInput("in", PortIn)
	g.AddOutput("out", PortOut)
}

func (g *Gain) GetName() string {
	return "Gain"
}

func (g *Gain) SetGain(gain float64) {
	g.gain = gain
}

func (g *Gain) Write(port Port, value float64) {
	switch port {
	case PortInGain:
		g.gain = value
	case PortIn:
		g.Sample += value
	}
}

func (g *Gain) Update(_ time.Duration) {
	g.ConnectionWrite(PortOut, g.Sample*g.gain)
	g.Sample = 0
}
