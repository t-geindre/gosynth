package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Delay struct {
	Module
	Sample     float64
	Queue      []float64
	Delay      time.Duration
	Feedback   float64
	SampleRate beep.SampleRate
	Cursor     int
}

func (d *Delay) Init(SampleRate beep.SampleRate) {
	d.Module.Init(SampleRate)
	d.SampleRate = SampleRate

	d.Queue = make([]float64, 0)

	d.AddInput("in", PortIn)
	d.AddOutput("out", PortOut)
}

func (d *Delay) GetName() string {
	return "Delay"
}

func (d *Delay) SetDelay(delay time.Duration) {
	d.Delay = delay
	d.Queue = make([]float64, d.SampleRate.N(delay))
	d.Cursor = 0
}

func (d *Delay) SetFeedback(feedback float64) {
	d.Feedback = feedback
}

func (d *Delay) Write(port Port, value float64) {
	switch port {
	case PortIn:
		d.Sample += value
	}
}

func (d *Delay) Update(time time.Duration) {
	if d.Delay != 0 {
		d.Sample += d.Queue[d.Cursor] * d.Feedback
		d.Queue[d.Cursor] = d.Sample
		d.Cursor = (d.Cursor + 1) % len(d.Queue)
	}

	d.ConnectionWrite(PortOut, d.Sample)
	d.Sample = 0
}
