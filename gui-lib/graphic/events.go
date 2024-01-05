package graphic

import "gosynth/event"

// DrawUpdateRequiredEvent
// Triggered when the graphic needs to be redrawn
// in most cases, it means that the size of the graphic has changed
var DrawUpdateRequiredEvent event.Id

// DrawStartEvent
// Triggered when the graphic children are about to be drawn
var DrawStartEvent event.Id

// DrawEndEvent
// Triggered when the graphic children have been drawn
var DrawEndEvent event.Id

func init() {
	DrawUpdateRequiredEvent = event.Register()
	DrawStartEvent = event.Register()
	DrawEndEvent = event.Register()
}
