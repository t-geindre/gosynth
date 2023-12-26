package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Limiter struct {
	Module
	Sample float64
	Thresh float64
}

func (l *Limiter) Init(rate beep.SampleRate) {
	l.Module.Init(rate)
	l.AddInput("in", PortIn)
	l.AddOutput("out", PortOut)
}

func (l *Limiter) SetThreshold(thresh float64) {
	l.Thresh = thresh
}

func (l *Limiter) Write(port Port, value float64) {
	switch port {
	case PortIn:
		l.Sample += value
	}
}

func (l *Limiter) Update(time time.Duration) {
	if l.Sample > l.Thresh {
		l.Sample = l.Thresh
	} else if l.Sample < -l.Thresh {
		l.Sample = -l.Thresh
	}

	l.ConnectionWrite(PortOut, l.Sample)
	l.Sample = 0
}
