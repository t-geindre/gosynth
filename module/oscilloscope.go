package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Oscilloscope struct {
	*Module
	time   time.Duration
	buffer []float64
}

func (o *Oscilloscope) Init(rate beep.SampleRate) {
	o.Module.Init(rate, o)
	o.time = time.Second * 2
}

func (o *Oscilloscope) Write(port Port, value float64) {
	switch port {
	case PortIn:
		o.buffer = append(o.buffer, value)
		//if len(o.buffer) > int(o.time.Seconds()*float64(o.SampleRate())) {
		//	o.buffer = o.buffer[1:]
		//}
	}

	o.Module.Write(port, value)
}

func (o *Oscilloscope) Update(time time.Duration) {

}
