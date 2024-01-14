package module

type Port uint8

const (
	// INPUTS

	PortInCV Port = iota
	PortInCV1
	PortInCV2
	PortInCV3
	PortInCV4
	PortInCV5
	PortInCV6
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
	PortInAttack
	PortInDecay
	PortInSustain
	PortInRelease
	PortInTime
	PortInTrigger

	// OUTPUTS

	PortOut
	PortOut1
	PortOut2
	PortOut3
	PortOut4
	PortOut5
	PortOut6
	PortOut7
	PortOutCv
	PortOutSin
	PortOutSquare
	PortOutSaw
	PortOutTriangle
	PortOutTrigger
	PortOutGate
)
