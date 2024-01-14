package module

import (
	"github.com/gopxl/beep"
	"math"
	"time"
)

type clockTrigger struct {
	port       Port
	ticks      int
	trigger    int
	multiplier float64
}

type Clock struct {
	*Module
	triggers map[Port]*clockTrigger
	ticksRef int
}

func NewClock(sr beep.SampleRate) *Clock {
	c := &Clock{}
	c.Module = NewModule(sr, c)

	c.triggers = make(map[Port]*clockTrigger)
	c.triggers[PortInCV1] = &clockTrigger{port: PortOut1}
	c.triggers[PortInCV2] = &clockTrigger{port: PortOut2}
	c.triggers[PortInCV3] = &clockTrigger{port: PortOut3}
	c.triggers[PortInCV4] = &clockTrigger{port: PortOut4}
	c.triggers[PortInCV5] = &clockTrigger{port: PortOut5}
	c.triggers[PortInCV6] = &clockTrigger{port: PortOut6}

	for portIn, trigger := range c.triggers {
		c.AddInput(portIn)
		c.AddOutput(trigger.port)
		c.Write(portIn, 0)
	}

	c.AddInput(PortInCV)
	c.Write(PortInCV, 0)

	return c
}

func (c *Clock) Write(port Port, value float64) {
	c.Module.Write(port, value)

	if trigger, ok := c.triggers[port]; ok {
		trigger.multiplier = math.Pow(2, math.Round(value*3))
		if trigger.multiplier > 0 {
			trigger.multiplier = 1 / trigger.multiplier
		}
		trigger.multiplier = math.Abs(trigger.multiplier)
		trigger.trigger = int(trigger.multiplier * float64(c.ticksRef))
		return
	}

	if port == PortInCV {
		c.ticksRef = c.GetSampleRate().N(time.Duration(int((1-(value+1)/2)*1000)) * time.Millisecond)
		for _, trigger := range c.triggers {
			trigger.trigger = int(trigger.multiplier * float64(c.ticksRef))
		}
	}
}

func (c *Clock) Update() {
	c.Module.Update()
	for _, trigger := range c.triggers {
		trigger.ticks++
		if trigger.ticks >= trigger.trigger {
			trigger.ticks = 0
			c.ConnectionWrite(trigger.port, 1)
		}
	}
}
