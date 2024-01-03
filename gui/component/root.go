package component

import (
	"gosynth/gui/control"
	"gosynth/gui/layout"
)

type Root struct {
	*Component
	mouse *control.Mouse
}

func NewRoot() *Root {
	r := &Root{}

	r.Component = NewComponent()
	r.mouse = control.NewMouse(r)

	return r
}

func (r *Root) Update() {
	r.mouse.Update()
	layout.Sync.Update()
	r.Component.Update()
}
