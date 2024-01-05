package connection

import "gosynth/event"

var ConnectionStartEvent event.Id
var ConnectionStopEvent event.Id
var ConnectionEnterEvent event.Id
var ConnectionLeaveEvent event.Id

func init() {
	ConnectionStartEvent = event.Register()
	ConnectionStopEvent = event.Register()
	ConnectionEnterEvent = event.Register()
	ConnectionLeaveEvent = event.Register()
}
