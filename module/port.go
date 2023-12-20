package module

type Port uint8

// Inputs
const (
	PortInGain Port = iota
	PortInFreq
	PortInGate
	PortInTime
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