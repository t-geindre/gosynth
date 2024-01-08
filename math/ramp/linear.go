package ramp

import (
	"github.com/gopxl/beep"
	"time"
)

type Linear struct {
	from      float64
	target    float64
	sDuration int // samples
	sDone     int // samples done
	sRate     beep.SampleRate
}

func NewLinear(sr beep.SampleRate, value float64) *Linear {
	return &Linear{
		from:   value,
		target: value,
		sRate:  sr,
	}
}

func (l *Linear) GoTo(target float64, duration time.Duration) {
	l.GoFromTo(l.Value(), target, duration)
}

func (l *Linear) GoFromTo(from, target float64, duration time.Duration) {
	l.from = from
	l.target = target
	l.sDuration = l.sRate.N(duration)
	l.sDone = 0
}

func (l *Linear) Value() float64 {
	if l.IsFinished() {
		return l.target
	}

	value := l.from + (l.target-l.from)*float64(l.sDone)/float64(l.sDuration)
	l.sDone++

	return value
}

func (l *Linear) IsFinished() bool {
	return l.sDone >= l.sDuration
}
