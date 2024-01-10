package module

import "github.com/gopxl/beep"

const OscBufferSize = 1024

type Scope struct {
	*Module
	valChan chan float64
}

func NewScope(sr beep.SampleRate) *Scope {
	s := &Scope{}
	s.Module = NewModule(sr, s)
	s.valChan = make(chan float64, OscBufferSize)

	s.AddInput(PortIn)
	s.AddInput(PortInSync)
	s.AddInput(PortInTime)

	return s
}

func (s *Scope) Write(port Port, value float64) {
	switch port {
	case PortIn:
		if len(s.valChan) < OscBufferSize {
			s.valChan <- value
		}
		s.ConnectionWrite(PortOut, value)
		// Drop value if channel full; avoid blocking
	case PortInSync:
	//clear everything, we received a sync signal from gui
	case PortInTime:
		//set time scale

	}
}

func (s *Scope) ReceiveValues() []float64 {
	values := make([]float64, len(s.valChan))
	for i := range values {
		values[i] = <-s.valChan
	}
	return values
}
