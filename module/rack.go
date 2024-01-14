package module

import (
	"github.com/gopxl/beep"
	"gosynth/event"
	clock "gosynth/time"
)

type Rack struct {
	*Module
	modules        []IModule
	clock          *clock.Clock
	sampleL        float64
	sampleR        float64
	connectionChan chan connectionCmd
}

func NewRack(clk *clock.Clock, rate beep.SampleRate) *Rack {
	r := &Rack{}
	r.Module = NewModule(rate, r)
	r.clock = clk
	r.connectionChan = make(chan connectionCmd, 3)
	r.modules = make([]IModule, 0)

	r.AddInput(PortInL)
	r.AddInput(PortInR)
	r.AddInput(PortIn)

	r.clock.AddListener(&r, clock.TickEvent, func(e event.IEvent) {
		r.Update()
	})

	return r
}

func (r *Rack) AddModule(module IModule) {
	r.modules = append(r.modules, module)
}

func (r *Rack) Dispose() {
	for _, module := range r.modules {
		module.Dispose()
	}
	r.clock.RemoveListener(r, clock.TickEvent)
}

func (r *Rack) Write(port Port, value float64) {
	switch port {
	case PortInL:
		r.sampleL += value * 0.07
	case PortInR:
		r.sampleR += value * 0.07
	case PortIn:
		r.sampleL += value * 0.07
		r.sampleR += value * 0.07
	}

	r.Module.Write(port, value)
}

func (r *Rack) Update() {
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

	r.Module.Update()

	r.sampleL = 0
	r.sampleR = 0

	for _, module := range r.modules {
		module.Update()
	}
}

func (r *Rack) GetSamples() (float64, float64) {
	return r.sampleL, r.sampleR
}

func (r *Rack) GetModules() []IModule {
	return r.modules
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
