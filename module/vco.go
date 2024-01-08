package module

import (
	"github.com/gopxl/beep"
	"math"
)

type VCO struct {
	*Module
	phaseCur, phaseAdv, phaseAcc float64
	octShift                     float64
	freqRef, freqRange           float64
	sDuration                    float64
}

func NewVCO(sr beep.SampleRate) *VCO {
	v := &VCO{}
	v.Module = NewModule(sr, v)
	v.freqRef = 440 // std A4
	v.freqRange = 4 // +-4 octaves
	v.sDuration = v.GetSampleRate().D(1).Seconds()
	v.Write(PortInVOct, 0)
	return v
}

func (v *VCO) Write(port Port, value float64) {
	switch port {
	case PortInVOct:
		// CV v/oct input to phase shifting
		v.phaseAdv = v.freqRef * math.Pow(2, value*v.freqRange) * math.Pow(2, v.octShift) * v.sDuration
	}

	v.Module.Write(port, value)
}

func (v *VCO) Update() {
	v.Module.Update()

	v.phaseCur += v.phaseAdv
	if v.phaseCur >= 1 {
		v.phaseCur -= 1
	}

	v.ConnectionWrite(PortOutSin, v.oscSin(v.phaseCur))
	v.ConnectionWrite(PortOutSquare, v.oscSquare(v.phaseCur))
	v.ConnectionWrite(PortOutSaw, v.oscSaw(v.phaseCur))
	v.ConnectionWrite(PortOutTriangle, v.oscTriangle(v.phaseCur))

}

func (v *VCO) oscSin(phase float64) float64 {
	radiantPhase := 2 * math.Pi * phase
	return math.Sin(radiantPhase)
}

func (v *VCO) oscSquare(phase float64) float64 {
	val := 0.0
	if phase < 0.5 {
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
