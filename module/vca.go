package module

import (
	"github.com/gopxl/beep"
	"gosynth/math/ramp"
	"time"
)

type VCA struct {
	*Module
	Gain       *ramp.Linear
	MasterGain float64
	Sample     float64
	time       time.Duration
}

func (g *VCA) Init(rate beep.SampleRate) {
	g.Module = &Module{}
	g.Module.Init(rate, g)
	g.Gain = ramp.NewLinear(1) // Todo find a better way to smooth the gain
	g.Write(PortCvIn, -1)
}

func (g *VCA) Write(port Port, value float64) {
	switch port {
	case PortCvIn:
		g.Gain.GoTo((value+1)/2, time.Millisecond*5, g.time)
	case PortIn:
		g.Sample += value
	}
	g.Module.Write(port, value)
}

func (g *VCA) Update(t time.Duration) {
	g.time = t
	g.Module.Update(t)
	g.ConnectionWrite(PortOut, g.Sample*g.Gain.Value(t))
	g.Sample = 0
}
