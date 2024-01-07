package ramp

import (
	"time"
)

type Linear struct {
	From   float64
	Target float64

	Duration time.Duration
	Start    time.Duration

	Finished bool
}

func NewLinear(value float64) *Linear {
	return &Linear{
		From:     value,
		Target:   value,
		Duration: time.Duration(0),
		Start:    time.Duration(0),
		Finished: true,
	}
}

func (l *Linear) GoTo(target float64, duration, time time.Duration) {
	l.GoFromTo(l.Value(time), target, duration, time)
}

func (l *Linear) GoFromTo(from, target float64, duration, time time.Duration) {
	l.From = from
	l.Target = target
	l.Duration = duration
	l.Start = time
	l.Finished = false
}

func (l *Linear) Value(time time.Duration) float64 {
	elapsed := time - l.Start

	if elapsed >= l.Duration {
		l.Finished = true
		return l.Target
	}

	return l.From + (l.Target-l.From)*float64(elapsed)/float64(l.Duration)
}

func (l *Linear) IsFinished() bool {
	return l.Finished
}
