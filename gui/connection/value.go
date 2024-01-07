package connection

import (
	"gosynth/math/ramp"
	audio "gosynth/module"
)

const (
	audioFrom float64 = -1
	audioTo   float64 = 1
)

type Value struct {
	// gui
	gFrom, gTo float64
	// audio
	module audio.IModule
	port   audio.Port
	ramp   *ramp.Linear
}

func NewValue(gFrom, gTo float64, module audio.IModule, port audio.Port) *Value {
	v := &Value{
		gFrom:  gFrom,
		gTo:    gTo,
		module: module,
		port:   port,
	}

	return v
}

func (v *Value) SendGuiValue(value float64) float64 {
	value = audioFrom + (audioTo-audioFrom)*(value-v.gFrom)/(v.gTo-v.gFrom)
	v.module.SendInput(v.port, value)
	return value
}

func (v *Value) ReceiveAudioValue() *float64 {
	value := v.module.ReceiveInput(v.port)
	if value == nil {
		return nil
	}

	*value = v.gFrom + (v.gTo-v.gFrom)*(*value-audioFrom)/(audioTo-audioFrom)

	return value
}
