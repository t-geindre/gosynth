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
	cursor   int
}

func NewDelay(sr beep.SampleRate) *Delay {
	d := &Delay{}
	d.Module = NewModule(sr, d)
	return d
}

func (d *Delay) Write(port Port, value float64) {
	switch port {
	case PortIn:
		d.sample += value
	case PortInDelay:
		delay := time.Duration((value + 1) / 2 * 3000 * float64(time.Millisecond))
		d.cursor = 0
		d.buffer = make([]float64, d.sampleRate.N(delay))
	case PortInFeedback:
		d.feedback = (value + 1) / 2
	}

	d.Module.Write(port, value)
}

func (d *Delay) Update() {
	d.Module.Update()

	if sLen := len(d.buffer); sLen > 0 {
		d.sample += d.buffer[d.cursor] * d.feedback
		d.buffer[d.cursor] = d.sample
		d.cursor = (d.cursor + 1) % sLen
	}

	d.ConnectionWrite(PortOut, d.sample)
	d.sample = 0
}
