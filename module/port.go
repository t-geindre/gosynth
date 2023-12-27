package module

type Port uint8

// Inputs
const (
	PortInGain Port = iota
	PortInFreq
	PortInGate
	PortIn
	PortInL
	PortInR
)

// Outputs
const (
	PortOut Port = iota
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
