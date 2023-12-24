package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"gosynth/module"
	"gosynth/note"
	"gosynth/output"
	clock "gosynth/time"
	"image"
	"log"
	"os"
	"time"
)

func main() {
	SampleRate := beep.SampleRate(44100)

	clk := clock.NewClock(SampleRate.D(1))
	rck := module.NewRack(clk, SampleRate)
	str := output.NewStreamer(clk, rck)

	go func() {
		w := app.NewWindow()
		err := run(w, rck)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	oscA := &module.Oscillator{}
	rck.AddModule(oscA)
	oscA.SetShape(module.OscillatorShapeTriangle)
	oscA.SetOctaveShift(1)

	oscB := &module.Oscillator{}
	rck.AddModule(oscB)
	oscB.SetAmplitude(.1)
	oscB.SetShape(module.OscillatorShapeSquare)
	oscB.SetOctaveShift(-2)

	gain := &module.Gain{}
	gain.SetMasterGain(.5)
	rck.AddModule(gain)

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

	delay.Connect(module.PortOut, gain, module.PortIn)

	gain.Connect(module.PortOut, lmt, module.PortIn)

	lmt.Connect(module.PortOut, rck, module.PortInL)
	lmt.Connect(module.PortOut, rck, module.PortInR)

	err := speaker.Init(SampleRate, SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}

	speaker.Play(str)
	app.Main()
}

func run(w *app.Window, rck *module.Rack) error {
	//th := material.NewTheme()
	var ops op.Ops
	for {
		e := w.NextEvent()
		th := material.NewTheme()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			for i, md := range rck.Modules {
				offset := op.Offset(image.Pt(10, i*25+10)).Push(&ops)
				title := material.Body1(th, md.GetName())
				title.Layout(gtx)
				offset.Pop()
			}
			/*


					wg := new(widget.Clickable)
					btn := material.Button(th, wg, "click me")
					//btn.Background = red
					btn.Text = "Play"
					btn.Layout(gtx)

					offset := op.Offset(image.Pt(10, 10)).Push(&ops)
				rect := clip.Rect{Max: image.Pt(200, 400)}
				_ = rect
				line := clip.Path{}
				line.Begin(&ops)
				line.Move(f32.Pt(0, 0))
				line.Line(f32.Pt(200, 400))
				//line.End()
				clp := clip.Stroke{Path: line.End(), Width: 2}.Op().Push(&ops)
				paint.ColorOp{Color: color.NRGBA{R: 50, G: 50, B: 50, A: 255}}.Add(&ops)
				paint.PaintOp{}.Add(&ops)
				clp.Pop()
				offset.Pop()
			*/
			e.Frame(&ops)
		}
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
