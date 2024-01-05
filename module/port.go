package module

type Port uint8

const (
	// INPUTS

	PortCvIn Port = iota
	PortIn
	PortInGate
	PortInL
	PortInR

	PortInGain
	PortInFreq

	// OUTPUTS

	PortOut
	PortCvOut

	PortOutFreq
	PortOutGate
)

/*
	TODO:
	Use a 0-10 Volt scale for CV
	1 Volt = 1 Octave = 12 Semitones
	Audio 0-10V as well, normalized to -1 to 1 in the end

	Add a quantizer module
		- Quantize CV to a scale
*/
