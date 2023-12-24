package module

import (
	"github.com/gopxl/beep"
	"gosynth/event"
	clock "gosynth/time"
	"time"
)

type Rack struct {
	Module
	Modules    []IModule
	Clock      *clock.Clock
	SampleL    float64
	SampleR    float64
	SampleRate beep.SampleRate
}

func NewRack(clock *clock.Clock, rate beep.SampleRate) *Rack {
	r := &Rack{}
	r.Clock = clock
	r.Init(rate)
	return r
}

func (r *Rack) Init(SampleRate beep.SampleRate) {
	r.Module.Init(SampleRate)

	r.SampleRate = SampleRate
	r.Modules = make([]IModule, 0)

	r.Clock.AddListener(r, r.Clock.Events.Tick, func(e event.ListenerArgs) {
		r.Update(e.(time.Duration))
	})

	r.AddInput("In L", PortInL)
	r.AddInput("In R", PortInR)
}

func (r *Rack) AddModule(module IModule) {
	module.Init(r.SampleRate)
	r.Modules = append(r.Modules, module)
}

func (r *Rack) Dispose() {
	for _, module := range r.Modules {
		module.Dispose()
	}
	r.Clock.RemoveListener(r, r.Clock.Events.Tick)
}

func (r *Rack) GetName() string {
	return "Rack"
}

func (r *Rack) Write(port Port, value float64) {
	switch port {
	case PortInL:
		r.SampleL += value
	case PortInR:
		r.SampleR += value
	}
}

func (r *Rack) Update(time time.Duration) {
	r.SampleL = 0
	r.SampleR = 0

	for _, module := range r.Modules {
		module.Update(time)
	}
}

func (r *Rack) GetSamples() (float64, float64) {
	return r.SampleL, r.SampleR
}

func (r *Rack) GetModules() []IModule {
	return r.Modules
}
