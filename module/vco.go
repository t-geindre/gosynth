package module

import (
	"github.com/gopxl/beep"
	"math"
	"time"
)

type OscillatorShape int

type VCO struct {
	Module
	Freq     float64
	OctShift float64
}

func (v *VCO) Init(rate beep.SampleRate) {
	v.Module.Init(rate, v)
	v.Write(PortInVOct, 0)
}

func (v *VCO) Write(port Port, value float64) {
	switch port {
	case PortInVOct:
		// CV v/oct input to frequency
		v.Freq = 440 * math.Pow(2, value*4)
	}

	v.Module.Write(port, value)
}

func (v *VCO) SetOctaveShift(octShift float64) {
	v.OctShift = octShift
}

func (v *VCO) Update(time time.Duration) {
	v.Module.Update(time)

	freq := v.Freq * math.Pow(2, v.OctShift)

	// Normalize 0-10V
	v.ConnectionWrite(PortOutSin, v.oscSin(time, freq))
	v.ConnectionWrite(PortOutSquare, v.oscSquare(time, freq))
	v.ConnectionWrite(PortOutSaw, v.oscSaw(time, freq))
	v.ConnectionWrite(PortOutTriangle, v.oscTriangle(time, freq))
}

func (v *VCO) oscSin(time time.Duration, freq float64) float64 {
	return math.Sin(2 * math.Pi * time.Seconds() * freq)
}

func (v *VCO) oscSquare(time time.Duration, freq float64) float64 {
	var val float64
	if math.Sin(2*math.Pi*time.Seconds()*freq) > 0 {
		val = 1
	} else {
		val = -1
	}

	return val
}

func (v *VCO) oscSaw(time time.Duration, freq float64) float64 {
	return 2 * (time.Seconds()*freq - math.Floor(0.5+time.Seconds()*freq))
}

func (v *VCO) oscTriangle(time time.Duration, freq float64) float64 {
	return math.Abs(2*(time.Seconds()*freq-math.Floor(0.5+time.Seconds()*freq))) - 1
}
