package main

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/gui"
	"gosynth/module"
	"gosynth/output"
	clock "gosynth/time"
	"log"
	"runtime"
	"time"
)

func main() {
	SampleRate := beep.SampleRate(44100)
	clk := clock.NewClock(SampleRate.D(1))
	rck := module.NewRack(clk, SampleRate)
	str := output.NewStreamer(clk, rck)

	go func() {
		runtime.LockOSThread()

		// Todo check if 50ms is enough (especially on Windows)
		// higher value would mean more latency, might be an issue for live performance (MIDI keyboard)
		err := speaker.Init(SampleRate, SampleRate.N(time.Millisecond*50))
		if err != nil {
			panic(err)
		}

		speaker.Play(str)
	}()

	app := gui.NewApp(rck)
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
