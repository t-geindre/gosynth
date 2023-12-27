package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/theme"
)

type Slider struct {
	*Node
	Dirty              bool
	Marks              int
	PaddingX, PaddingY int
	MarkMargin         int
	Value              float64
	ValueMin, ValueMax float64
	MouseLDown         bool
}

func NewSlider(width, height int) *Slider {
	s := &Slider{}
	s.Node = NewNode(width, height, s)
	s.Dirty = true
	s.Marks = 25
	s.PaddingX = 5
	s.PaddingY = 5
	s.MarkMargin = 2

	s.Dispatcher.AddListener(&s, LeftMouseDownEvent, func(e event.IEvent) {
		s.MouseLeftDown()
	})

	s.Dispatcher.AddListener(&s, LeftMouseUpEvent, func(e event.IEvent) {
		s.MouseLeftUp()
	})

	return s
}

func (s *Slider) Update() error {
	if s.MouseLDown {
		_, py := s.GetAbsolutePosition()
		_, my := ebiten.CursorPosition()
		my -= py
		s.SetValue(float64(my) / float64(s.Height-s.PaddingY*2))
	}

	return s.Node.Update()
}

func (s *Slider) Clear() {
	if s.Dirty {
		s.Dirty = false
		s.Image.Clear()
		vector.DrawFilledRect(s.Image, 0, 0, float32(s.Width), float32(s.Height), theme.Colors.BackgroundInverted, false)

		// Draw marks
		markWidth := float32(s.Width - s.PaddingX*2)
		markHeight := float32(s.Height-s.PaddingY*2-(s.Marks-1)*s.MarkMargin) / float32(s.Marks)
		markOnTrigger := float32(s.Value-s.ValueMin) / float32(s.ValueMax-s.ValueMin) * float32(s.Marks)
		for i := 0; i < s.Marks; i++ {
			col := theme.Colors.Off
			if float32(i) > markOnTrigger {
				col = theme.Colors.On
			}
			x := float32(s.PaddingX)
			y := (markHeight+float32(s.MarkMargin))*float32(i) + float32(s.PaddingY)
			vector.DrawFilledRect(s.Image, x, y, markWidth, markHeight, col, false)
		}
	}

	s.Node.Clear()
}

func (s *Slider) SetValue(value float64) {
	s.Value = value
	s.Dirty = true
}

func (s *Slider) SetRange(min, max float64) {
	s.ValueMin = min
	s.ValueMax = max
	s.Dirty = true
}

func (s *Slider) MouseLeftDown() {
	s.MouseLDown = true
}

func (s *Slider) MouseLeftUp() {
	s.MouseLDown = false
}
