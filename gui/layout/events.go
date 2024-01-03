package layout

import "gosynth/event"

var ResizeEvent event.Id
var MoveEvent event.Id
var UpdatedEvent event.Id
var UpdateStartsEvent event.Id

func init() {
	ResizeEvent = event.Register()
	MoveEvent = event.Register()
	UpdatedEvent = event.Register()
	UpdateStartsEvent = event.Register()
}
