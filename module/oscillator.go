package module

import (
	"github.com/gopxl/beep"
	"math"
	"time"
)

type OscillatorShape int

const (
	OscillatorShapeSine OscillatorShape = iota
	OscillatorShapeSquare
	OscillatorShapeSaw
	OscillatorShapeTriangle
)

type Oscillator struct {
	Module
	Freq     float64
	Shape    OscillatorShape
	OctShift float64
}

func (o *Oscillator) Init(rate beep.SampleRate) {
	o.Module.Init(rate)
	o.AddInput("freq", PortInFreq)
	o.AddOutput("out", PortOut)
}

func (o *Oscillator) GetName() string {
	return "Oscillator oscillator"
}

func (o *Oscillator) Write(port Port, value float64) {
	switch port {
	case PortInFreq:
		o.SetFreq(value)
	}
}

func (o *Oscillator) SetFreq(freq float64) {
	o.Freq = freq
}

func (o *Oscillator) SetShape(shape OscillatorShape) {
	o.Shape = shape
}

func (o *Oscillator) SetOctaveShift(octShift float64) {
	o.OctShift = octShift
}

func (o *Oscillator) Update(time time.Duration) {
	freq := o.Freq * math.Pow(2, o.OctShift)
	val := 0.0

	switch o.Shape {
	case OscillatorShapeSine:
		val = math.Sin(2 * math.Pi * time.Seconds() * freq)
	case OscillatorShapeSquare:
		if math.Sin(2*math.Pi*time.Seconds()*freq) > 0 {
			val = 1
		} else {
			val = -1
		}
	case OscillatorShapeSaw:
		val = 2 * (time.Seconds()*freq - math.Floor(0.5+time.Seconds()*freq))
	case OscillatorShapeTriangle:
		val = math.Abs(2*(time.Seconds()*freq-math.Floor(0.5+time.Seconds()*freq))) - 1
	}

	o.ConnectionWrite(PortOut, val)
}
