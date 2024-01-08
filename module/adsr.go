package module

import (
	"github.com/gopxl/beep"
	"gosynth/math/ramp"
	"time"
)

type EnvState uint8

const (
	EnvStateAttack EnvState = iota
	EnvStateDecay
	EnvStateSustain
	EnvStateRelease
	EnvStateOff
)

type EnvPhase struct {
	time.Duration
	Target float64
}

type ADSR struct {
	*Module
	state  EnvState
	phases [4]EnvPhase
	on     bool
	ramp   *ramp.Linear
}

func NewAdsr(rate beep.SampleRate) {
	a := &ADSR{}
	a.Module = NewModule(rate, a)
	a.state = EnvStateOff
	a.phases = [4]EnvPhase{
		{time.Millisecond * 40, 1},  // Attack
		{time.Millisecond * 20, .5}, // Decay
		{0, .5},                     // Sustain
		{time.Millisecond * 20, 0},  // Release
	}
	a.ramp = ramp.NewLinear(rate, 0)
}

func (a *ADSR) Write(port Port, value float64) {
	switch port {
	case PortInGate:
		a.on = value > 0
	}
	a.Module.Write(port, value)
}

func (a *ADSR) Update() {
	a.Module.Update()

	switch a.state {
	case EnvStateAttack:
		if a.ramp.IsFinished() {
			a.GoToState(EnvStateDecay)
		}
	case EnvStateDecay:
		if a.ramp.IsFinished() {
			a.GoToState(EnvStateSustain)
		}
	case EnvStateSustain:
		if !a.on {
			a.GoToState(EnvStateRelease)
		}
	case EnvStateRelease:
		if a.ramp.IsFinished() {
			a.GoToState(EnvStateOff)
		}
	case EnvStateOff:
		if a.on {
			a.GoToState(EnvStateAttack)
		}
	}

	a.ConnectionWrite(PortOutCv, a.ramp.Value()*2-1)
}

func (a *ADSR) GoToState(state EnvState) {
	a.state = state
	if state == EnvStateOff {
		return
	}
	a.ramp.GoTo(a.phases[state].Target, a.phases[state].Duration)
}
