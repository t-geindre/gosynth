package output

import (
	"gosynth/module"
	"gosynth/time"
)

type Streamer struct {
	clock   *time.Clock
	input   *module.Rack
	stopped bool
	command chan bool
}

func NewStreamer(clock *time.Clock, input *module.Rack) *Streamer {
	s := &Streamer{}
	s.clock = clock
	s.input = input
	s.stopped = true
	s.command = make(chan bool, 3)

	return s
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	select {
	case cmd := <-s.command:
		s.stopped = !cmd
	default:
	}

	if s.stopped {
		return 0, true
	}

	for i := range samples {
		samples[i][0], samples[i][1] = s.input.GetSamples()
		s.clock.Tick()
	}

	return len(samples), true
}

func (s *Streamer) Err() error {
	return nil
}

func (s *Streamer) Play() chan bool {
	return s.command
}

func (s *Streamer) IsPlaying() bool {
	return !s.stopped
}
