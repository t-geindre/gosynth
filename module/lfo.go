package module

import (
	"github.com/gopxl/beep"
	"math"
	"time"
)

type LFO struct {
	Module
	Freq     float64
	OctShift float64
}

func (l *LFO) Init(rate beep.SampleRate) {
	l.Module.Init(rate, l)
	l.Write(PortInVOct, 0)
}

func (l *LFO) Write(port Port, value float64) {
	switch port {
	case PortInVOct:
		// CV l/oct input to frequency
		l.Freq = 10 * math.Pow(2, value*4)
	}

	l.Module.Write(port, value)
}

func (l *LFO) SetOctaveShift(octShift float64) {
	l.OctShift = octShift
}

func (l *LFO) Update(time time.Duration) {
	l.Module.Update(time)
	freq := l.Freq * math.Pow(2, l.OctShift)

	// Normalize 0-10V
	l.ConnectionWrite(PortOutSin, l.oscSin(time, freq))
	l.ConnectionWrite(PortOutSquare, l.oscSquare(time, freq))
	l.ConnectionWrite(PortOutSaw, l.oscSaw(time, freq))
	l.ConnectionWrite(PortOutTriangle, l.oscTriangle(time, freq))
}

func (l *LFO) oscSin(time time.Duration, freq float64) float64 {
	return math.Sin(2 * math.Pi * time.Seconds() * freq)
}

func (l *LFO) oscSquare(time time.Duration, freq float64) float64 {
	var val float64
	if math.Sin(2*math.Pi*time.Seconds()*freq) > 0 {
		val = 1
	} else {
		val = -1
	}

	return val
}

func (l *LFO) oscSaw(time time.Duration, freq float64) float64 {
	return 2 * (time.Seconds()*freq - math.Floor(0.5+time.Seconds()*freq))
}

func (l *LFO) oscTriangle(time time.Duration, freq float64) float64 {
	return math.Abs(2*(time.Seconds()*freq-math.Floor(0.5+time.Seconds()*freq))) - 1
}
