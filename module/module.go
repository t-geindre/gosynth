package module

import (
	"github.com/gopxl/beep"
	"time"
)

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
	Inputs      []IO
	Outputs     []IO
	Connections []Connection
	CommandChan chan Command
}

func (m *Module) Init(_ beep.SampleRate) {
	m.Inputs = make([]IO, 0)
	m.Outputs = make([]IO, 0)
	m.Connections = make([]Connection, 0)
	m.CommandChan = make(chan Command, 3)
}

func (m *Module) AddInput(name string, port Port) {
	m.Inputs = append(m.Inputs, IO{
		Name: name,
		Port: port,
	})
}

func (m *Module) AddOutput(name string, port Port) {
	m.Outputs = append(m.Outputs, IO{
		Name: name,
		Port: port,
	})
}

func (m *Module) GetInputs() []IO {
	return m.Inputs
}

func (m *Module) GetOutputs() []IO {
	return m.Outputs
}

func (m *Module) Connect(srcPort Port, destModule IModule, destPort Port) {
	m.Connections = append(m.Connections, Connection{
		SrcPort:    srcPort,
		DestModule: destModule,
		DestPort:   destPort,
	})
}

func (m *Module) ConnectionWrite(srcPort Port, value float64) {
	for _, con := range m.Connections {
		if con.SrcPort == srcPort {
			con.DestModule.Write(con.DestPort, value)
		}
	}
}

// SendCommand Thread-safe way to send a command to a module
func (m *Module) SendCommand(port Port, value float64) {
	m.CommandChan <- Command{Port: port, Value: value}
}

func (m *Module) Write(port Port, value float64) {
}

func (m *Module) Dispose() {
}

func (m *Module) Update(time time.Duration) {
	select {
	case cmd := <-m.CommandChan:
		m.Write(cmd.Port, cmd.Value)
	default:
	}
}
