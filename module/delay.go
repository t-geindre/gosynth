package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Delay struct {
	*Module
	sample   float64
	buffer   []float64
	feedback float64
	mix      float64
	cursor   int
}

func NewDelay(sr beep.SampleRate) *Delay {
	d := &Delay{}
	d.Module = NewModule(sr, d)

	d.AddInput(PortIn)
	d.AddInput(PortInDelay)
	d.AddInput(PortInFeedback)
	d.AddInput(PortInMix)

	d.AddOutput(PortOut)

	d.Write(PortInDelay, -1)
	d.Write(PortInFeedback, -1)
	d.Write(PortInMix, 0)

	return d
}

func (d *Delay) Write(port Port, value float64) {
	switch port {
	case PortIn:
		d.sample += value
	case PortInDelay:
		delay := time.Duration((value + 1) / 2 * 3000 * float64(time.Millisecond))
		newLen := d.sampleRate.N(delay)
		if newLen > len(d.buffer) {
			d.buffer = append(d.buffer, make([]float64, newLen-len(d.buffer))...)
		} else {
			d.buffer = d.buffer[:newLen]
		}
		if d.cursor >= len(d.buffer) {
			d.cursor = 0
		}
	case PortInFeedback:
		d.feedback = (value + 1) / 2
	case PortInMix:
		d.mix = (value + 1) / 2
	}

	d.Module.Write(port, value)
}

func (d *Delay) Update() {
	d.Module.Update()

	if sLen := len(d.buffer); sLen > 0 {
		wetSample := d.sample + d.buffer[d.cursor]*d.feedback
		d.sample = wetSample*d.mix + d.sample*(1-d.mix)
		d.buffer[d.cursor] = d.sample
		d.cursor = (d.cursor + 1) % sLen
	}

	d.ConnectionWrite(PortOut, d.sample)
	d.sample = 0
}
