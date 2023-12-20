package output

import (
	"gosynth/module"
	"gosynth/time"
)

type Streamer struct {
	clock *time.Clock
	input *module.Rack
}

func NewStreamer(clock *time.Clock, input *module.Rack) *Streamer {
	s := &Streamer{}
	s.clock = clock
	s.input = input

	return s
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0], samples[i][1] = s.input.GetSamples()
		s.clock.Tick()
	}

	return len(samples), true
}

func (s *Streamer) Err() error {
	return nil
}
