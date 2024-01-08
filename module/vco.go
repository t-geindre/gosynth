package module

import (
	"github.com/gopxl/beep"
	"math"
	"time"
)

type VCO struct {
	Module
	Freq             float64
	currentPhase     float64
	phaseAccumulator float64
	phaseAdvance     float64
	OctShift         float64
	SampleRate       beep.SampleRate
}

func (v *VCO) Init(rate beep.SampleRate) {
	v.Module.Init(rate, v)
	v.SampleRate = rate
	v.Write(PortInVOct, 0)
}

func (v *VCO) Write(port Port, value float64) {
	switch port {
	case PortInVOct:
		// CV v/oct input to frequency
		v.Freq = 440 * math.Pow(2, value*4) * math.Pow(2, v.OctShift)
		v.phaseAdvance = v.Freq * v.SampleRate.D(1).Seconds()
	}

	v.Module.Write(port, value)
}

func (v *VCO) SetOctaveShift(octShift float64) {
	v.OctShift = octShift
}

func (v *VCO) Update(time time.Duration) {
	v.Module.Update(time)

	v.currentPhase += v.phaseAdvance
	if v.currentPhase >= 1 {
		v.currentPhase -= 1
	}

	// Normalize 0-10V
	v.ConnectionWrite(PortOutSin, v.oscSin(v.currentPhase))
	v.ConnectionWrite(PortOutSquare, v.oscSquare(v.currentPhase))
	v.ConnectionWrite(PortOutSaw, v.oscSaw(v.currentPhase))
	v.ConnectionWrite(PortOutTriangle, v.oscTriangle(v.currentPhase))

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
