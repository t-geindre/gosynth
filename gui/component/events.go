package component

import "gosynth/event"

var UpdateEvent event.Id

func init() {
	UpdateEvent = event.Register()
}
