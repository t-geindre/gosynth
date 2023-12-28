package output

import (
	"gosynth/module"
	"gosynth/time"
)

type Streamer struct {
	clock    *time.Clock
	rack     *module.Rack
	silenced bool
	command  chan bool
}

func NewStreamer(clock *time.Clock, rack *module.Rack) *Streamer {
	s := &Streamer{}
	s.clock = clock
	s.rack = rack
	s.silenced = false
	s.command = make(chan bool, 3)

	return s
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	select {
	case cmd := <-s.command:
		s.silenced = cmd
	default:
	}

	for i := range samples {
		if !s.silenced {
			samples[i][0], samples[i][1] = s.rack.GetSamples()
		} else {
			samples[i][0], samples[i][1] = 0, 0
		}
		s.clock.Tick()
	}

	return len(samples), true
}

func (s *Streamer) Err() error {
	return nil
}

func (s *Streamer) Silence() chan bool {
	return s.command
}

func (s *Streamer) IsSilenced() bool {
	return s.silenced
}

func (s *Streamer) GetRack() *module.Rack {
	return s.rack
}
