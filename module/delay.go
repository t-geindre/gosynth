package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Delay struct {
	Module
	sample     float64
	buffer     []float64
	delay      time.Duration
	feedback   float64
	sampleRate beep.SampleRate
	cursor     int
}

func (d *Delay) Init(SampleRate beep.SampleRate) {
	d.Module.Init(SampleRate, d)
	d.sampleRate = SampleRate
	d.SetDelay(0)
}

func (d *Delay) SetDelay(delay time.Duration) {
	d.delay = delay
	d.cursor = 0
	d.buffer = make([]float64, d.sampleRate.N(delay))
}

func (d *Delay) SetFeedback(feedback float64) {
	d.feedback = feedback
}

func (d *Delay) Write(port Port, value float64) {
	switch port {
	case PortIn:
		d.sample += value
	case PortDelayIn:
		delay := time.Duration((value + 1) / 2 * 1000 * float64(time.Millisecond))
		d.SetDelay(delay)
	case PortFeedbackIn:
		d.SetFeedback((value + 1) / 2)
	}

	d.Module.Write(port, value)
}

func (d *Delay) Update(time time.Duration) {
	d.Module.Update(time)
	if d.delay != 0 {
		d.sample += d.buffer[d.cursor] * d.feedback
		d.buffer[d.cursor] = d.sample
		d.cursor = (d.cursor + 1) % len(d.buffer)
	}

	d.ConnectionWrite(PortOut, d.sample)
	d.sample = 0
}
