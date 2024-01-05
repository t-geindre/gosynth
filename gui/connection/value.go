package connection

import (
	audio "gosynth/module"
)

type Value struct {
	// gui
	gFrom, gTo float64
	// audio
	aFrom, aTo float64
	module     audio.IModule
	port       audio.Port
}

func NewValue(aFrom, aTo, gFrom, gTo float64, module audio.IModule, port audio.Port) *Value {
	v := &Value{
		aFrom:  aFrom,
		aTo:    aTo,
		gFrom:  gFrom,
		gTo:    gTo,
		module: module,
		port:   port,
	}

	return v
}

func (v *Value) SendGuiValue(value float64) float64 {
	// normalize gvalue to avalue
	value = v.aFrom + (v.aTo-v.aFrom)*(value-v.gFrom)/(v.gTo-v.gFrom)
	v.module.SendInput(v.port, value)
	return value
}

func (v *Value) ReceiveAudioValue() *float64 {
	value := v.module.ReceiveInput(v.port)
	if value == nil {
		return nil
	}
	*value = v.aFrom + (v.aTo-v.aFrom)*(*value)

	return value
}
