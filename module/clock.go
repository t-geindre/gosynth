package module

import (
	"github.com/gopxl/beep"
	"time"
)

type Clock struct {
	*Module
	ticksTrigger int
	ticks1       int
	ticks2       int
	ticks4       int
	ticks8       int
}

func NewClock(sr beep.SampleRate) *Clock {
	c := &Clock{}
	c.Module = NewModule(sr, c)

	c.AddInput(PortInCV)
	c.AddOutput(PortOut1) // 1/1
	c.AddOutput(PortOut2) // 1/2
	c.AddOutput(PortOut3) // 1/4
	c.AddOutput(PortOut4) // 1/8

	c.Write(PortInCV, 0)

	return c
}

func (c *Clock) Write(port Port, value float64) {
	if port == PortInCV {
		c.ticksTrigger = c.GetSampleRate().N(time.Duration(int((1-(value+1)/2)*1000)) * time.Millisecond)
	}
}

func (c *Clock) Update() {
	c.Module.Update()

	c.ticks1++
	c.ticks2++
	c.ticks4++
	c.ticks8++

	if c.ticks8 > c.ticksTrigger/8 {
		c.ticks8 = 0
		c.ConnectionWrite(PortOut4, 1)
	}

	if c.ticks4 > c.ticksTrigger/4 {
		c.ticks4 = 0
		c.ConnectionWrite(PortOut3, 1)
	}
	if c.ticks2 > c.ticksTrigger/2 {
		c.ticks2 = 0
		c.ConnectionWrite(PortOut2, 1)
	}

	if c.ticks1 > c.ticksTrigger {
		c.ticks1 = 0
		c.ConnectionWrite(PortOut1, 1)
	}
}
