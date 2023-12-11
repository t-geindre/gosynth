package osc

type Square struct {
	*Sin
}

func (s *Square) Stream(time, freq float64, samples *[2]float64) {
	s.Sin.Stream(time, freq, samples)
	samples[0] = s.sign(samples[0])
	samples[1] = s.sign(samples[1])
}

func (s *Square) sign(sample float64) float64 {
	if sample > 0 {
		return 1
	}
	return -1
}
