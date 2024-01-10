package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Sequencer4 struct {
	*Module
	ticks        int
	ticksTrigger int
	values       [4]float64
	cursor       int
}

func NewSequencer4(sr beep.SampleRate) *Sequencer4 {
	s := &Sequencer4{}
	s.Module = NewModule(sr, s)

	s.AddInput(PortIn1)
	s.AddInput(PortIn2)
	s.AddInput(PortIn3)
	s.AddInput(PortIn4)
	s.AddInput(PortInCV)

	s.AddOutput(PortOutGate)
	s.AddOutput(PortOutTrigger)
	s.AddOutput(PortOutCv)

	s.Write(PortInCV, 0)

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
	case PortInCV:
		s.ticksTrigger = s.GetSampleRate().N(time.Duration(int((value+1)/2*1000)) * time.Millisecond)
	}
}

func (s *Sequencer4) Update() {
	s.Module.Update()

	if s.ticks > s.ticksTrigger {
		s.ticks = 0
		s.cursor = (s.cursor + 1) % 4
		s.ConnectionWrite(PortOutGate, -1)
		s.ConnectionWrite(PortOutTrigger, 1)
	}

	s.ConnectionWrite(PortOutGate, 1)
	s.ConnectionWrite(PortOutCv, s.values[s.cursor])

	s.ticks++

}
