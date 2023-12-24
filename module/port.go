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
