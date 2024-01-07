package module

import (
	"github.com/gopxl/beep"
	"time"
)

type IModule interface {
	Init(rate beep.SampleRate)
	Write(port Port, value float64)
	Connect(srcPort Port, destModule IModule, destPort Port)
	Disconnect(srcPort Port, destModule IModule, destPort Port)
	Update(time time.Duration)
	Dispose()
	// SendInput Thread-safe way to send a command to a module
	SendInput(port Port, value float64)
	// ReceiveInput Thread-safe way to read data sent to a module
	ReceiveInput(port Port) *float64
	// ReceiveOutput Thread-safe way to get data output from a module
	ReceiveOutput(port Port) *float64
}
