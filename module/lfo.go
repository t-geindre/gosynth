package module

import "github.com/gopxl/beep"

func NewLFO(sr beep.SampleRate) *VCO {
	l := NewVCO(sr)
	l.freqRange = 4
	l.freqRef = 10
	l.Write(PortInVOct, 0)
	return l
}
