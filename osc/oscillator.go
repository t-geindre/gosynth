package osc

type Oscillator interface {
	Stream(time, freq float64, samples *[2]float64)
}
