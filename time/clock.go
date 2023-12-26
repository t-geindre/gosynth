package time

import (
	"gosynth/event"
	"time"
)

type Clock struct {
	time         time.Duration
	tickDuration time.Duration
	event.Dispatcher
	Events struct {
		Tick event.Id
	}
}

func NewClock(tickDuration time.Duration) *Clock {
	c := &Clock{}
	c.Init(tickDuration)
	return c
}

func (c *Clock) Init(tickDuration time.Duration) {
	c.Dispatcher.Init()
	c.Events.Tick = c.RegisterEvent()
	c.time = 0
	c.tickDuration = tickDuration
}

func (c *Clock) Tick() {
	c.time += c.tickDuration
	c.Dispatch(event.NewEvent(c.Events.Tick, c))
}

func (c *Clock) GetTime() time.Duration {
	return c.time
}
