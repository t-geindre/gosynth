package module

import (
	"github.com/gopxl/beep"
	"math"
)

type VCO struct {
	*Module
	phaseAcc             float64
	freqRef, freqRange   float64
	freq                 float64
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

	v.Write(PortInVOct, 0)
	v.Write(PortInPw, 0)
	v.Write(PortInFmFact, 0)

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
	}

	v.Module.Write(port, value)
}

func (v *VCO) Update() {
	v.Module.Update()

	freq := v.freq
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
}

func (v *VCO) oscSin(phase float64) float64 {
	radiantPhase := 2 * math.Pi * phase
	return math.Sin(radiantPhase)
}

func (v *VCO) oscSquare(phase float64) float64 {
	val := 0.0
	if phase < v.pwm {
		val = 1
	} else {
		val = -1
	}

	return val
}

func (v *VCO) oscSaw(phase float64) float64 {
	if phase < 0.5 {
		return 2 * phase
	} else {
		return 2*phase - 2
	}
}

func (v *VCO) oscTriangle(phase float64) float64 {
	if phase < 0.25 {
		return 4 * phase
	} else if phase < 0.75 {
		return -4*phase + 2
	} else {
		return 4*phase - 4
	}
}
