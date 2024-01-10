package module

import (
	"github.com/gopxl/beep"
	"gosynth/math/ramp"
	"time"
)

type VCA struct {
	*Module
	gain   *ramp.Linear
	sample float64
}

func NewVCA(sr beep.SampleRate) *VCA {
	v := &VCA{}
	v.Module = NewModule(sr, v)
	v.gain = ramp.NewLinear(sr, 0)

	v.AddInput(PortIn)
	v.AddInput(PortInCV)
	v.AddOutput(PortOut)

	return v
}

func (g *VCA) Write(port Port, value float64) {
	switch port {
	case PortInCV:
		g.gain.GoTo((value+1)/2, time.Millisecond*5)
	case PortIn:
		g.sample += value
	}
	g.Module.Write(port, value)
}

func (g *VCA) Update() {
	g.Module.Update()
	g.ConnectionWrite(PortOut, g.sample*g.gain.Value())
	g.sample = 0
}
