package module

import (
	"github.com/gopxl/beep"
	"math"
	"time"
)

type FilterMode uint8

const (
	PassFilterModeLow FilterMode = iota
	PassFilterModeHigh
)

type PassFilter struct {
	Module
	SampleRate beep.SampleRate
	Sample     float64
	Buffer     float64
	Alpha      float64
	Mode       FilterMode
}

func (h *PassFilter) Init(SampleRate beep.SampleRate) {
	h.Module.Init(SampleRate)
	h.SampleRate = SampleRate
	h.Alpha = 0
	h.AddInput("in", PortIn)
	h.AddOutput("out", PortOut)
}

func (h *PassFilter) SetCutOff(cutoff float64) {
	tan := math.Tan(math.Pi * cutoff / float64(h.SampleRate))
	h.Alpha = (tan - 1) / (tan + 1)
}

func (h *PassFilter) SetMode(mode FilterMode) {
	h.Mode = mode
}

func (h *PassFilter) Write(port Port, value float64) {
	switch port {
	case PortIn:
		h.Sample += value
	}
}

func (h *PassFilter) Update(time.Duration) {
	pass := h.Alpha*h.Sample + h.Buffer
	h.Buffer = h.Sample - h.Alpha*pass

	if h.Mode == PassFilterModeHigh {
		pass *= -1
	}

	h.ConnectionWrite(PortOut, h.Sample+pass)
	h.Sample = 0
}
