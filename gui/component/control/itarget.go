package control

import "gosynth/event"

type ITarget interface {
	GetTargetAt(x, y int) (ITarget, error)
	Dispatch(event event.IEvent)
}
