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
		{time.Millisecond * 20, 0},  // Decay
		{0, 0},                      // Sustain
		{time.Millisecond * 20, -1}, // Release
	}
	a.ramp = ramp.NewLinear(rate, 0)
}

func (a *ADSR) Write(port Port, value float64) {
	switch port {
	case PortInGate:
		a.on = value > 0
	case PortInAttack:
		a.phases[EnvStateAttack].Duration = time.Duration((value+1)/2) * time.Millisecond * 1000
	case PortInDecay:
		a.phases[EnvStateDecay].Duration = time.Duration((value+1)/2) * time.Millisecond * 1000
	case PortInSustain:
		a.phases[EnvStateSustain].Target = value
		a.phases[EnvStateDecay].Target = a.phases[EnvStateSustain].Target
	case PortInRelease:
		a.phases[EnvStateRelease].Duration = time.Duration((value+1)/2) * time.Millisecond * 1000
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

	a.ConnectionWrite(PortOutCv, a.ramp.Value())
}

func (a *ADSR) GoToState(state EnvState) {
	a.state = state
	if state == EnvStateOff {
		return
	}
	a.ramp.GoTo(a.phases[state].Target, a.phases[state].Duration)
}
