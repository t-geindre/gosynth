package module

import (
	"github.com/gopxl/beep"
)

type Sequencer4 struct {
	*Module
	values  [4]float64
	cursor  int
	trigger bool
}

func NewSequencer4(sr beep.SampleRate) *Sequencer4 {
	s := &Sequencer4{}
	s.Module = NewModule(sr, s)

	s.AddInput(PortIn1)
	s.AddInput(PortIn2)
	s.AddInput(PortIn3)
	s.AddInput(PortIn4)
	s.AddInput(PortInTrigger)

	s.AddOutput(PortOutGate)
	s.AddOutput(PortOutCv)

	return s
}

func (s *Sequencer4) Write(port Port, value float64) {
	s.Module.Write(port, value)
	switch port {
	case PortIn1:
		s.values[0] = value
	case PortIn2:
		s.values[1] = value
	case PortIn3:
		s.values[2] = value
	case PortIn4:
		s.values[3] = value
	case PortInTrigger:
		s.trigger = value > 0
	}
}

func (s *Sequencer4) Update() {
	s.Module.Update()

	if s.trigger {
		s.trigger = false
		s.cursor = (s.cursor + 1) % 4
		s.ConnectionWrite(PortOutGate, -1)
	}

	s.ConnectionWrite(PortOutGate, 1)
	s.ConnectionWrite(PortOutCv, s.values[s.cursor])
}
