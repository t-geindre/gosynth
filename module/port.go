package module

type Port uint8

const (
	// INPUTS

	PortInCV Port = iota
	PortInDelay
	PortInFeedback
	PortIn
	PortInGate
	PortInL
	PortInR
	PortInVOct
	PortInValue1
	PortInValue2
	PortInValue3
	PortInValue4
	PortInPwmFact
	PortInPwm
	PortInPw
	PortInMix
	PortInSync
	PortInFmFact
	PortInFm

	// OUTPUTS

	PortOut
	PortOutCv
	PortOutSin
	PortOutSquare
	PortOutSaw
	PortOutTriangle
	PortOutTrigger
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
