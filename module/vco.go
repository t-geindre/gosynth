package module

import (
	"github.com/gopxl/beep"
	"math"
)

type VCO struct {
	*Module
	// Todo phase shift not on all osc shapes
	// Could also be a ramp to avoid clapping on phase shifting
	phaseAcc, phaseShift float64
	freqRef, freqRange   float64
	freq, octShift       float64
	freqMod, freqModFact float64
	sDuration            float64
	pwm                  float64
}

func NewVCO(sr beep.SampleRate) *VCO {
	v := &VCO{}
	v.Module = NewModule(sr, v)
	v.freqRef = 440 // std A4
	v.freqRange = 4 // +-4 octaves
	v.sDuration = v.GetSampleRate().D(1).Seconds()

	v.AddInput(PortInVOct)
	v.AddInput(PortInPw)
	v.AddInput(PortInSync)
	v.AddInput(PortInPwmFact)
	v.AddInput(PortInPwm)
	v.AddInput(PortInFmFact)
	v.AddInput(PortInFm)
	v.AddInput(PortInOctShift)
	v.AddInput(PortInPhaseShift)

	v.AddOutput(PortOutSin)
	v.AddOutput(PortOutSquare)
	v.AddOutput(PortOutSaw)
	v.AddOutput(PortOutTriangle)

	v.Write(PortInVOct, 0)
	v.Write(PortInPw, 0)
	v.Write(PortInFmFact, 0)
	v.Write(PortInOctShift, 0)
	v.Write(PortInPwmFact, 0)

	return v
}

func (v *VCO) Write(port Port, value float64) {
	switch port {
	case PortInVOct:
		v.freq = v.freqRef * math.Pow(2, value*v.freqRange)
	case PortInPw:
		v.pwm = (value+1)/2*0.45 + 0.05
	case PortInSync:
		v.phaseAcc = 0
	case PortInPwmFact:
		// Todo
	case PortInPwm:
		// Todo
	case PortInFmFact:
		v.freqModFact = (value + 1) / 2
	case PortInFm:
		v.freqMod = v.freqRef * math.Pow(2, value*v.freqRange)
	case PortInOctShift:
		v.octShift = math.Round(value * 8)
	case PortInPhaseShift:
		v.phaseShift = (value + 1) / 2
	}

	v.Module.Write(port, value)
}

func (v *VCO) Update() {
	v.Module.Update()

	freq := v.freq * math.Pow(2, v.octShift)
	if v.freqMod != 0 {
		freq += v.freqModFact * v.freqMod
	}

	v.phaseAcc += freq * v.sDuration
	if v.phaseAcc >= 1 {
		v.phaseAcc -= 1
	}

	v.ConnectionWrite(PortOutSin, v.oscSin(v.phaseAcc))
	v.ConnectionWrite(PortOutSquare, v.oscSquare(v.phaseAcc))
	v.ConnectionWrite(PortOutSaw, v.oscSaw(v.phaseAcc))
	v.ConnectionWrite(PortOutTriangle, v.oscTriangle(v.phaseAcc))

	v.freqMod = 0
}

func (v *VCO) oscSin(phase float64) float64 {
	radiantPhase := 2 * math.Pi * (phase + v.phaseShift)
	return math.Sin(radiantPhase)
}

func (v *VCO) oscSquare(phase float64) float64 {
	if phase > v.phaseShift && phase < v.phaseShift+v.pwm {
		return 1
	} else {
		return -1
	}
}

func (v *VCO) oscSaw(phase float64) float64 {
	phase = math.Mod(phase+v.phaseShift, 1)
	return 2*phase - 1
}

func (v *VCO) oscTriangle(phase float64) float64 {
	if phase > v.phaseShift && phase < v.phaseShift+0.5 {
		return 4*phase - 1
	} else {
		return -4*phase + 3
	}
}
