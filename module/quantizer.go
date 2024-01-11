package module

import (
	"github.com/gopxl/beep"
)

type Quantizer struct {
	*Module
	ref   []float64
	value float64
}

func NewQuantizer(sr beep.SampleRate) *Quantizer {
	q := &Quantizer{}
	q.Module = NewModule(sr, q)

	return q
}

func (q *Quantizer) Write(port Port, value float64) {
	switch port {
	case PortIn:
		q.value = value
	}
}

func (q *Quantizer) Update() {
	// move value to a one of the c chromatic scale note
	value := q.value
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	value = value * 12
	value = value - float64(int(value))
	q.value = value
	q.ConnectionWrite(PortOut, q.value)
}
