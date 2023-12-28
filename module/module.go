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
	Connections []Connection
	CommandChan chan Command
	IModule     IModule
}

func (m *Module) Init(_ beep.SampleRate, imodule IModule) {
	m.Connections = make([]Connection, 0)
	m.CommandChan = make(chan Command, 3)
	m.IModule = imodule
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

// SendInput Thread-safe way to send a command to a module
func (m *Module) SendInput(port Port, value float64) {
	m.CommandChan <- Command{Port: port, Value: value}
}

func (m *Module) Write(port Port, value float64) {
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
