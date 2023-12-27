package node

import "gosynth/event"

var LeftMouseDownEvent event.Id
var LeftMouseUpEvent event.Id

func init() {
	LeftMouseDownEvent = event.Register()
	LeftMouseUpEvent = event.Register()
}