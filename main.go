package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"gosynth/module"
	"gosynth/output"
	clock "gosynth/time"
	"image/color"
	"log"
	"os"
	"time"
)

func main() {
	go func() {
		w := app.NewWindow()
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	SampleRate := beep.SampleRate(44100)

	clk := clock.NewClock(SampleRate.D(1))
	rck := module.NewRack(clk, SampleRate)
	str := output.NewStreamer(clk, rck)

	oscA := &module.Oscillator{}
	oscA.SetShape(module.OscillatorShapeSquare)
	rck.AddModule(oscA)

	oscB := &module.Oscillator{}
	oscB.SetShape(module.OscillatorShapeSaw)
	oscB.SetOctaveShift(-2)
	rck.AddModule(oscB)

	gain := &module.Gain{}
	gain.SetGain(.5)
	rck.AddModule(gain)

	sqr := &module.Sequencer{}
	rck.AddModule(sqr)
	sqr.AppendAfter(220, time.Second/8, time.Second/10)
	sqr.AppendAfter(440, time.Second/8, time.Second/10)
	sqr.AppendAfter(880, time.Second/8, time.Second/10)
	sqr.AppendAfter(440, time.Second/8, time.Second/10)
	sqr.SetLoop(true)

	adsr := &module.Adsr{}
	rck.AddModule(adsr)

	delay := &module.Delay{}
	rck.AddModule(delay)
	delay.SetDelay(time.Millisecond * 300)
	delay.SetFeedback(.1)

	lmt := &module.Limiter{}
	lmt.SetThreshold(1)
	rck.AddModule(lmt)

	pfi := &module.PassFilter{}
	rck.AddModule(pfi)
	pfi.SetMode(module.PassFilterModeLow)
	pfi.SetCutOff(880)

	sqr.Connect(module.PortOutFreq, oscA, module.PortInFreq)
	sqr.Connect(module.PortOutFreq, oscB, module.PortInFreq)
	sqr.Connect(module.PortOutGate, adsr, module.PortInGate)

	oscA.Connect(module.PortOut, adsr, module.PortIn)
	oscB.Connect(module.PortOut, adsr, module.PortIn)

	adsr.Connect(module.PortOut, delay, module.PortIn)

	delay.Connect(module.PortOut, pfi, module.PortIn)

	pfi.Connect(module.PortOut, lmt, module.PortIn)

	lmt.Connect(module.PortOut, gain, module.PortIn)

	gain.Connect(module.PortOut, rck, module.PortInL)
	gain.Connect(module.PortOut, rck, module.PortInR)

	err := speaker.Init(SampleRate, SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}

	speaker.Play(str)
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	for {
		e := w.NextEvent()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			title := material.H1(th, "bip")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
