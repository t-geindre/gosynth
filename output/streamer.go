package output

import (
	"gosynth/module"
	"gosynth/time"
)

type Streamer struct {
	clock *time.Clock
	rack  *module.Rack
}

func NewStreamer(clock *time.Clock, rack *module.Rack) *Streamer {
	s := &Streamer{}
	s.clock = clock
	s.rack = rack

	return s
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0], samples[i][1] = s.rack.GetSamples()
		s.clock.Tick()
	}

	return len(samples), true
}

func (s *Streamer) Err() error {
	return nil
}
