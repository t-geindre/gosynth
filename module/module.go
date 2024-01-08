package module

import (
	"github.com/gopxl/beep"
)

const chanPortBuffering = 1

type Command struct {
	Port  Port
	Value float64
}

type Connection struct {
	SrcPort    Port
	DestModule IModule
	DestPort   Port
}

type Module struct {
	connections    []Connection
	commandChan    chan Command
	writtenInputs  map[Port]chan float64
	writtenOutputs map[Port]chan float64
	sampleRate     beep.SampleRate
	IModule        IModule
}

func NewModule(sr beep.SampleRate, imodule IModule) *Module {
	m := &Module{}
	m.connections = make([]Connection, 1)
	m.commandChan = make(chan Command, 1)
	m.writtenInputs = make(map[Port]chan float64)
	m.writtenOutputs = make(map[Port]chan float64)
	m.sampleRate = sr
	m.IModule = imodule
	return m
}

func (m *Module) GetSampleRate() beep.SampleRate {
	return m.sampleRate
}

func (m *Module) Connect(srcPort Port, destModule IModule, destPort Port) {
	m.connections = append(m.connections, Connection{
		SrcPort:    srcPort,
		DestModule: destModule,
		DestPort:   destPort,
	})
}

func (m *Module) Disconnect(srcPort Port, destModule IModule, destPort Port) {
	for i, con := range m.connections {
		if con.SrcPort == srcPort && con.DestModule == destModule && con.DestPort == destPort {
			m.connections = append(m.connections[:i], m.connections[i+1:]...)
			return
		}
	}
}

func (m *Module) ConnectionWrite(srcPort Port, value float64) {
	if _, ok := m.writtenOutputs[srcPort]; !ok {
		m.writtenOutputs[srcPort] = make(chan float64, chanPortBuffering)
	}

	if len(m.writtenOutputs[srcPort]) < chanPortBuffering {
		m.writtenOutputs[srcPort] <- value
	}

	for _, con := range m.connections {
		if con.SrcPort == srcPort {
			con.DestModule.Write(con.DestPort, value)
		}
	}
}

// SendInput Thread-safe way to send a command to a module
func (m *Module) SendInput(port Port, value float64) {
	m.commandChan <- Command{Port: port, Value: value}
}

// ReceiveInput Thread-safe way to read data sent to a module
func (m *Module) ReceiveInput(port Port) *float64 {
	if _, ok := m.writtenInputs[port]; !ok {
		return nil
	}

	if len(m.writtenInputs[port]) > 0 {
		val := <-m.writtenInputs[port]
		return &val
	}

	return nil
}

// ReceiveOutput Thread-safe way to get data output from a module
func (m *Module) ReceiveOutput(port Port) *float64 {
	if _, ok := m.writtenOutputs[port]; !ok {
		return nil
	}

	if len(m.writtenOutputs[port]) > 0 {
		val := <-m.writtenOutputs[port]
		return &val
	}

	return nil
}

func (m *Module) Write(port Port, value float64) {
	if _, ok := m.writtenInputs[port]; !ok {
		m.writtenInputs[port] = make(chan float64, chanPortBuffering)
	}

	if len(m.writtenInputs[port]) < chanPortBuffering {
		m.writtenInputs[port] <- value
	}
}

func (m *Module) Read(port Port) float64 {
	return 0
}

func (m *Module) Dispose() {
}

func (m *Module) Update() {
	select {
	case cmd := <-m.commandChan:
		m.GetIModule().Write(cmd.Port, cmd.Value)
	default:
	}
}

func (m *Module) GetIModule() IModule {
	// Todo get rid of this method (and the IModule ref)
	return m.IModule
}
