package main

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/gui"
	"gosynth/module"
	"gosynth/note"
	"gosynth/output"
	clock "gosynth/time"
	"log"
	"time"
)

func main() {
	SampleRate := beep.SampleRate(44100)

	clk := clock.NewClock(SampleRate.D(1))
	rck := module.NewRack(clk, SampleRate)
	str := output.NewStreamer(clk, rck)

	err := speaker.Init(SampleRate, SampleRate.N(time.Millisecond*10))
	if err != nil {
		panic(err)
	}

	speaker.Play(str)

	app := gui.NewApp(str)
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}

}

func AddTetrisSequence(sqr *module.Sequencer, interval, duration time.Duration) {
	sqr.AppendAfter(note.E5, interval+duration*2, duration*2)
	sqr.AppendAfter(note.B4, interval+duration*2, duration)
	sqr.AppendAfter(note.C5, interval+duration, duration)
	sqr.AppendAfter(note.D5, interval+duration, duration*2)
	sqr.AppendAfter(note.C5, interval+duration*2, duration)
	sqr.AppendAfter(note.B4, interval+duration, duration)
	sqr.AppendAfter(note.A4, interval+duration, duration*2)
	sqr.AppendAfter(note.A4, interval+duration*2, duration)
	sqr.AppendAfter(note.C5, interval+duration, duration)
	sqr.AppendAfter(note.E5, interval+duration, duration*2)
	sqr.AppendAfter(note.D5, interval+duration*2, duration)
	sqr.AppendAfter(note.C5, interval+duration, duration)
	sqr.AppendAfter(note.B4, interval+duration, duration*3)
	sqr.AppendAfter(note.C5, interval+duration*3, duration)
	sqr.AppendAfter(note.D5, interval+duration, duration*2)
	sqr.AppendAfter(note.E5, interval+duration*2, duration*2)
	sqr.AppendAfter(note.C5, interval+duration*2, duration*2)
	sqr.AppendAfter(note.A4, interval+duration*2, duration*2)
	sqr.AppendAfter(note.A4, interval+duration*2, duration*3)
	sqr.AppendAfter(0, interval+duration*2, duration*20)
}
