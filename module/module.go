package module

import (
	"github.com/gopxl/beep"
	"time"
)

const chanPortBuffering = 3

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
	Connections    []Connection
	CommandChan    chan Command
	WrittenInputs  map[Port]chan float64
	WrittenOutputs map[Port]chan float64
	IModule        IModule
}

func (m *Module) Init(_ beep.SampleRate, imodule IModule) {
	m.Connections = make([]Connection, 0)
	m.CommandChan = make(chan Command, 1)
	m.IModule = imodule
	m.WrittenInputs = make(map[Port]chan float64)
	m.WrittenOutputs = make(map[Port]chan float64)
}

func (m *Module) Connect(srcPort Port, destModule IModule, destPort Port) {
	m.Connections = append(m.Connections, Connection{
		SrcPort:    srcPort,
		DestModule: destModule,
		DestPort:   destPort,
	})
}

func (m *Module) Disconnect(srcPort Port, destModule IModule, destPort Port) {
	for i, con := range m.Connections {
		if con.SrcPort == srcPort && con.DestModule == destModule && con.DestPort == destPort {
			m.Connections = append(m.Connections[:i], m.Connections[i+1:]...)
			return
		}
	}
}

func (m *Module) ConnectionWrite(srcPort Port, value float64) {
	if _, ok := m.WrittenOutputs[srcPort]; !ok {
		m.WrittenOutputs[srcPort] = make(chan float64, chanPortBuffering)
	}

	if len(m.WrittenOutputs[srcPort]) < chanPortBuffering {
		m.WrittenOutputs[srcPort] <- value
	}

	for _, con := range m.Connections {
		if con.SrcPort == srcPort {
			con.DestModule.Write(con.DestPort, value)
		}
	}
}

// SendInput Thread-safe way to send a command to a module
func (m *Module) SendInput(port Port, value float64) {
	m.CommandChan <- Command{Port: port, Value: value}
}

// ReceiveInput Thread-safe way to read data sent to a module
func (m *Module) ReceiveInput(port Port) *float64 {
	if _, ok := m.WrittenInputs[port]; !ok {
		return nil
	}

	if len(m.WrittenInputs[port]) > 0 {
		val := <-m.WrittenInputs[port]
		return &val
	}

	return nil
}

// ReceiveOutput Thread-safe way to get data output from a module
func (m *Module) ReceiveOutput(port Port) *float64 {
	if _, ok := m.WrittenOutputs[port]; !ok {
		return nil
	}

	if len(m.WrittenOutputs[port]) > 0 {
		val := <-m.WrittenOutputs[port]
		return &val
	}

	return nil
}

func (m *Module) Write(port Port, value float64) {
	if _, ok := m.WrittenInputs[port]; !ok {
		m.WrittenInputs[port] = make(chan float64, chanPortBuffering)
	}

	if len(m.WrittenInputs[port]) < chanPortBuffering {
		m.WrittenInputs[port] <- value
	}

}

func (m *Module) Read(port Port) float64 {
	return 0
}

func (m *Module) Dispose() {
}

func (m *Module) Update(time time.Duration) {
	select {
	case cmd := <-m.CommandChan:
		m.GetIModule().Write(cmd.Port, cmd.Value)
	default:
	}
}

func (m *Module) GetIModule() IModule {
	return m.IModule
}
