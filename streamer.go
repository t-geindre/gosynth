package main

import (
	"github.com/gopxl/beep"
	"gosynth/osc"
	"math/rand"
	"time"
)

type streamer struct {
	beep.SampleRate
	time float64
	osc.Oscillator
	freqChannel chan float64
	freq        float64
}

func (s *streamer) Init() {
	s.freq = 440
	s.freqChannel = make(chan float64)

	go func() {
		freq := 0.0

		for {
			s.freqChannel <- freq
			for {
				newFreq := 220 * float64(rand.Intn(4)+1)
				if newFreq != freq {
					freq = newFreq
					break
				}
			}
			time.Sleep(time.Second / 4)
		}
	}()
}

func (s *streamer) Stream(samples [][2]float64) (n int, ok bool) {
	select {
	case s.freq = <-s.freqChannel:
	default:
	}
	for i := range samples {
		s.Oscillator.Stream(s.time, s.freq, &samples[i])
		s.time += 1 / float64(s.SampleRate)
	}

	return len(samples), true
}

func (s *streamer) Err() error {
	return nil
}
