package time

import "gosynth/event"

var TickEvent event.Id

func init() {
	TickEvent = event.Register()
}
