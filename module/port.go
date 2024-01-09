package module

type Port uint8

const (
	// INPUTS

	PortInCV Port = iota
	PortInDelay
	PortInFeedback
	PortIn
	PortIn1
	PortIn2
	PortIn3
	PortIn4
	PortInGate
	PortInL
	PortInR
	PortInVOct
	PortInOctShift
	PortInPhaseShift
	PortInPwmFact
	PortInPwm
	PortInPw
	PortInMix
	PortInSync
	PortInFmFact
	PortInFm

	// OUTPUTS

	PortOut
	PortOut1
	PortOut2
	PortOut3
	PortOut4
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
