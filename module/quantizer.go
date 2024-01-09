package module

import (
	"github.com/gopxl/beep"
)

type Quantizer struct {
	*Module
	ref []float64
}

func NewQuantizer(sr beep.SampleRate) *Quantizer {
	q := &Quantizer{}
	q.Module = NewModule(sr, q)

	return q
}

func (q *Quantizer) Write(port Port, value float64) {

}
