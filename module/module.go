package module

import "github.com/gopxl/beep"

type Module struct {
	Inputs      []IO
	Outputs     []IO
	Connections []Connection
}

type Connection struct {
	SrcPort    Port
	DestModule IModule
	DestPort   Port
}

func (m *Module) Init(_ beep.SampleRate) {
	m.Inputs = make([]IO, 0)
	m.Outputs = make([]IO, 0)
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

func (m *Module) Write(port Port, value float64) {
}

func (m *Module) Dispose() {
}
