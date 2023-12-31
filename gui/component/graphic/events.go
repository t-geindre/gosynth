package graphic

import "gosynth/event"

// DrawUpdateRequiredEvent
// Triggered when the graphic needs to be redrawn
// in most cases, it means that the size of the graphic has changed
var DrawUpdateRequiredEvent event.Id

// DrawEvent
// Triggered when the graphic is drawn
// I s also triggered after a DrawUpdateRequiredEvent is triggered
var DrawEvent event.Id

func init() {
	DrawUpdateRequiredEvent = event.Register()
	DrawEvent = event.Register()
}
