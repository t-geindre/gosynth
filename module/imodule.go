package module

import (
	"github.com/gopxl/beep"
	"time"
)

type IModule interface {
	Init(rate beep.SampleRate)
	Write(port Port, value float64)
	Connect(srcPort Port, destModule IModule, destPort Port)
	Update(time time.Duration)
	Dispose()
}
