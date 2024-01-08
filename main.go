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

		err := speaker.Init(SampleRate, SampleRate.N(time.Millisecond*16))
		if err != nil {
			panic(err)
		}

		speaker.Play(str)
	}()

	app := gui.NewApp(rck)
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
	/*
		op := &oto.NewContextOptions{}
		op.SampleRate = 44100
		op.ChannelCount = 2
		op.Format = oto.FormatUnsignedInt8
		const max = 1<<7 - 1
		freq := 440.0
		b := int(math.Sin(2*math.Pi*freq*float64(time.Now().Second())) * 0.3 * max)
		for ch := 0; ch < op.ChannelCount; ch++ {
			buf[num*i+ch] = byte(b + 128)
		}
	*/
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
