package time

import (
	"gosynth/event"
	"time"
)

type Clock struct {
	time         time.Duration
	tickDuration time.Duration
	event.Dispatcher
}

func NewClock(tickDuration time.Duration) *Clock {
	c := &Clock{}
	c.Init(tickDuration)
	return c
}

func (c *Clock) Init(tickDuration time.Duration) {
	c.Dispatcher.Init()
	c.time = 0
	c.tickDuration = tickDuration
}

func (c *Clock) Tick() {
	c.time += c.tickDuration
	c.Dispatch(event.NewEvent(TickEvent, c))
}

func (c *Clock) GetTime() time.Duration {
	return c.time
}
