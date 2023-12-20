package module

import (
	"github.com/gopxl/beep"
	"time"
)

type IO struct {
	Name string
	Port Port
}

type IModule interface {
	Init(rate beep.SampleRate)
	GetName() string
	GetInputs() []IO
	GetOutputs() []IO
	Write(port Port, value float64)
	Connect(srcPort Port, destModule IModule, destPort Port)
	Update(time time.Duration)
	Dispose()
}
