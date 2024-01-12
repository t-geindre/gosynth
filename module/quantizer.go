package module

import (
	"github.com/gopxl/beep"
	"gosynth/math/normalizer"
	"gosynth/note"
)

type Quantizer struct {
	*Module
	ref   []float64
	value float64
}

func NewQuantizer(sr beep.SampleRate) *Quantizer {
	// Todo : faster implementation please
	q := &Quantizer{}
	q.Module = NewModule(sr, q)
	q.ref = []float64{note.C0, note.D0, note.E0, note.F0, note.G0, note.A0, note.B0, note.C1, note.D1, note.E1, note.F1, note.G1, note.A1, note.B1, note.C2, note.D2, note.E2, note.F2, note.G2, note.A2, note.B2, note.C3, note.D3, note.E3, note.F3, note.G3, note.A3, note.B3, note.C4, note.D4, note.E4, note.F4, note.G4, note.A4, note.B4, note.C5, note.D5, note.E5, note.F5, note.G5, note.A5, note.B5, note.C6, note.D6, note.E6, note.F6, note.G6, note.A6, note.B6, note.C7, note.D7, note.E7, note.F7, note.G7, note.A7, note.B7, note.C8, note.D8, note.E8, note.F8, note.G8, note.A8, note.B8}
	for i := 0; i < len(q.ref); i++ {
		q.ref[i] = normalizer.FrequencyToCv(q.ref[i])
	}

	q.AddInput(PortIn)
	q.AddOutput(PortOut)

	q.value = -2

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
	if q.value < -1 {
		return
	}
	value := q.ref[0]
	for i := 0; i < len(q.ref); i++ {
		if q.value > q.ref[i] {
			value = q.ref[i]
		}
		if q.value < q.ref[i] {
			break
		}
	}
	q.ConnectionWrite(PortOut, value)
	q.value = -2
}
