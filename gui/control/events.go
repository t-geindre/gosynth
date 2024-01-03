package control

import "gosynth/event"

var LeftMouseDownEvent event.Id
var LeftMouseUpEvent event.Id
var MouseEnterEvent event.Id
var MouseLeaveEvent event.Id

func init() {
	LeftMouseDownEvent = event.Register()
	LeftMouseUpEvent = event.Register()
	MouseEnterEvent = event.Register()
	MouseLeaveEvent = event.Register()
}
