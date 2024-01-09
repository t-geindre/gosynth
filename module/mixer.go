package module

import "github.com/gopxl/beep"

type Mixer struct {
	*Module
}

func NewMixer(sr beep.SampleRate) *Mixer {
	m := &Mixer{}
	m.Module = NewModule(sr, m)
	return m
}

func (m *Mixer) Write(port Port, value float64) {
	m.Module.Write(port, value)
	switch port {
	case PortIn1, PortIn2, PortIn3, PortIn4:
		m.ConnectionWrite(PortOut, value)
	}
}
