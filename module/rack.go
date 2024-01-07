package module

import (
	"github.com/gopxl/beep"
	"gosynth/event"
	clock "gosynth/time"
	"time"
)

type Rack struct {
	Module
	Modules        []IModule
	Clock          *clock.Clock
	SampleL        float64
	SampleR        float64
	SampleRate     beep.SampleRate
	connectionChan chan connectionCmd
}

func NewRack(clock *clock.Clock, rate beep.SampleRate) *Rack {
	r := &Rack{}
	r.Clock = clock
	r.connectionChan = make(chan connectionCmd, 3)
	r.Init(rate)
	return r
}

func (r *Rack) Init(SampleRate beep.SampleRate) {
	r.Module.Init(SampleRate, r)

	r.SampleRate = SampleRate
	r.Modules = make([]IModule, 0)

	r.Clock.AddListener(r, clock.TickEvent, func(e event.IEvent) {
		r.Update(e.GetSource().(*clock.Clock).GetTime())
	})
}

func (r *Rack) AddModule(module IModule) {
	module.Init(r.SampleRate)
	r.Modules = append(r.Modules, module)
}

func (r *Rack) Dispose() {
	for _, module := range r.Modules {
		module.Dispose()
	}
	r.Clock.RemoveListener(r, clock.TickEvent)
}

func (r *Rack) Write(port Port, value float64) {
	switch port {
	case PortInL:
		r.SampleL += value
	case PortInR:
		r.SampleR += value
	case PortIn:
		r.SampleL += value
		r.SampleR += value
	}

	r.Module.Write(port, value)
}

func (r *Rack) Update(time time.Duration) {
	select {
	case cmd := <-r.connectionChan:
		switch cmd.Action {
		case ConnectionCreate:
			cmd.Scr.Connect(cmd.PSrc, cmd.Dst, cmd.PDst)
		case ConnectionDelete:
			cmd.Scr.Disconnect(cmd.PSrc, cmd.Dst, cmd.PDst)
		}
	default:
	}

	r.Module.Update(time)

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

func (r *Rack) CreateModuleConnection(scr IModule, pSrc Port, dst IModule, pDst Port) {
	r.connectionChan <- connectionCmd{scr, dst, pSrc, pDst, ConnectionCreate}
}

func (r *Rack) DeleteModuleConnection(scr IModule, pSrc Port, dst IModule, pDst Port) {
	r.connectionChan <- connectionCmd{scr, dst, pSrc, pDst, ConnectionDelete}
}

type connectionCmd struct {
	Scr    IModule
	Dst    IModule
	PSrc   Port
	PDst   Port
	Action connectionCmAction
}

type connectionCmAction uint8

const (
	ConnectionCreate connectionCmAction = iota
	ConnectionDelete
)
