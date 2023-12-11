package main

import (
	"github.com/gopxl/beep/speaker"
	"gosynth/osc"
	"time"
)

func main() {
	str := streamer{
		SampleRate: 44100,
		Oscillator: &osc.Sin{},
	}

	str.Init()
	err := speaker.Init(str.SampleRate, str.N(time.Second/10))
	if err != nil {
		panic(err)
	}

	speaker.Play(&str)
	select {}
}
