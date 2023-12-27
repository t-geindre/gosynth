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

	oscA := &module.Oscillator{}
	rck.AddModule(oscA)
	oscA.SetShape(module.OscillatorShapeTriangle)
	oscA.SetOctaveShift(1)

	oscB := &module.Oscillator{}
	rck.AddModule(oscB)
	oscB.SetAmplitude(.1)
	oscB.SetShape(module.OscillatorShapeSquare)
	oscB.SetOctaveShift(-2)

	vca := &module.VCA{}
	vca.SetGain(.5)
	rck.AddModule(vca)

	sqr := &module.Sequencer{}
	rck.AddModule(sqr)
	AddTetrisSequence(sqr, time.Millisecond*10, time.Millisecond*100)
	sqr.SetLoop(true)

	adsr := &module.Adsr{}
	rck.AddModule(adsr)

	delay := &module.Delay{}
	rck.AddModule(delay)
	delay.SetDelay(time.Millisecond * 200)
	delay.SetFeedback(.15)

	lmt := &module.Limiter{}
	lmt.SetThreshold(1)
	rck.AddModule(lmt)

	sqr.Connect(module.PortOutFreq, oscA, module.PortInFreq)
	sqr.Connect(module.PortOutFreq, oscB, module.PortInFreq)
	sqr.Connect(module.PortOutGate, adsr, module.PortInGate)

	oscA.Connect(module.PortOut, adsr, module.PortIn)
	oscB.Connect(module.PortOut, adsr, module.PortIn)

	adsr.Connect(module.PortOut, delay, module.PortIn)

	delay.Connect(module.PortOut, vca, module.PortIn)

	vca.Connect(module.PortOut, lmt, module.PortIn)

	lmt.Connect(module.PortOut, rck, module.PortInL)
	lmt.Connect(module.PortOut, rck, module.PortInR)

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
	sqr.AppendAfter(note.E5, interval+duration, duration*2)
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
	sqr.AppendAfter(0, interval+duration*4, 0)
}
