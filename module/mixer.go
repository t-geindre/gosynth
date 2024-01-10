package module

import "github.com/gopxl/beep"

type Mixer struct {
	*Module
}

func NewMixer(sr beep.SampleRate) *Mixer {
	m := &Mixer{}
	m.Module = NewModule(sr, m)

	m.AddInput(PortIn1)
	m.AddInput(PortIn2)
	m.AddInput(PortIn3)
	m.AddInput(PortIn4)

	m.AddOutput(PortOut)

	return m
}

func (m *Mixer) Write(port Port, value float64) {
	m.Module.Write(port, value)
	switch port {
	case PortIn1, PortIn2, PortIn3, PortIn4:
		// Todo Possible stack overflow, fix this ?
		m.ConnectionWrite(PortOut, value)
	}
}
