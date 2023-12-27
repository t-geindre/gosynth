package module

import (
	"github.com/gopxl/beep"
	"gosynth/module/ramp"
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

type Adsr struct {
	Module
	State  EnvState
	Phases [4]EnvPhase
	On     bool
	Ramp   ramp.Linear
	Sample float64
}

func (a *Adsr) Init(rate beep.SampleRate) {
	a.Module.Init(rate)
	a.State = EnvStateOff
	a.Phases = [4]EnvPhase{
		EnvPhase{time.Millisecond * 20, 1},  // Attack
		EnvPhase{time.Millisecond * 20, .5}, // Decay
		EnvPhase{0, .5},                     // Sustain
		EnvPhase{time.Millisecond * 20, 0},  // Release
	}
	a.AddInput("gate", PortInGate)
	a.AddInput("trigger", PortIn)
	a.AddOutput("out", PortOut)
}

func (a *Adsr) Write(port Port, value float64) {
	switch port {
	case PortInGate:
		a.On = value > 0
	case PortIn:
		a.Sample += value
	}
}

func (a *Adsr) Update(time time.Duration) {
	a.Module.Update(time)
	switch a.State {
	case EnvStateAttack:
		if a.Ramp.IsFinished() {
			a.GoToState(EnvStateDecay, time)
		}
	case EnvStateDecay:
		if a.Ramp.IsFinished() {
			a.GoToState(EnvStateSustain, time)
		}
	case EnvStateSustain:
		if !a.On {
			a.GoToState(EnvStateRelease, time)
		}
	case EnvStateRelease:
		if a.Ramp.IsFinished() {
			a.GoToState(EnvStateOff, time)
		}
	case EnvStateOff:
		if a.On {
			a.GoToState(EnvStateAttack, time)
		}
	}

	a.ConnectionWrite(PortOut, a.Sample*a.Ramp.Value(time))
	a.Sample = 0
}

func (a *Adsr) GoToState(state EnvState, time time.Duration) {
	a.State = state
	if state == EnvStateOff {
		return
	}
	a.Ramp.GoTo(a.Phases[state].Target, a.Phases[state].Duration, time)
}
