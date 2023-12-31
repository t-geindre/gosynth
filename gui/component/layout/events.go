package layout

import "gosynth/event"

var ResizeEvent event.Id
var MoveEvent event.Id
var UpdatedEvent event.Id

func init() {
	ResizeEvent = event.Register()
	MoveEvent = event.Register()
	UpdatedEvent = event.Register()
}
