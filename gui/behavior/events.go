package behavior

import "gosynth/event"

var DragStartEvent event.Id
var DragStopEvent event.Id
var DragEvent event.Id

func init() {
	DragStartEvent = event.Register()
	DragStopEvent = event.Register()
	DragEvent = event.Register()
}
