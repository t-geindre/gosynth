package osc

import "math"

type Sin struct {
}

func (s *Sin) Stream(time, freq float64, samples *[2]float64) {
	sample := math.Sin(2 * math.Pi * freq * time)
	samples[0] = sample
	samples[1] = sample
}
