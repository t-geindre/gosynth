package module

import "github.com/gopxl/beep"

type Multiplier struct {
	*Module
}

func NewMultiplier(sr beep.SampleRate) *Multiplier {
	m := &Multiplier{}
	m.Module = NewModule(sr, m)

	m.AddInput(PortIn)

	m.AddOutput(PortOut1)
	m.AddOutput(PortOut2)
	m.AddOutput(PortOut3)
	m.AddOutput(PortOut4)

	return m
}

func (m *Multiplier) Write(port Port, value float64) {
	m.Module.Write(port, value)
	switch port {
	case PortIn:
		// Todo Possible stack overflow, fix this ?
		m.ConnectionWrite(PortOut1, value)
		m.ConnectionWrite(PortOut2, value)
		m.ConnectionWrite(PortOut3, value)
		m.ConnectionWrite(PortOut4, value)
	}
}
