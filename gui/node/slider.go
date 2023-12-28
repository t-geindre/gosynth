package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/theme"
	"gosynth/math/ramp"
	"math"
	"time"
)

type Slider struct {
	*Node
	Marks              int
	PaddingX, PaddingY int
	MarkMargin         int
	Value              float64
	ValueMin, ValueMax float64
	MouseLDown         bool
	Ramp               *ramp.Linear
}

func NewSlider() *Slider {
	s := &Slider{}
	s.Node = NewNode(1, 1, s)
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

func (s *Slider) Update(time time.Duration) error {
	if s.MouseLDown {
		_, py := s.GetAbsolutePosition()
		_, my := ebiten.CursorPosition()
		s.SetValue(s.ValueMin + (s.ValueMax-s.ValueMin)*(1-(float64(my-py-s.PaddingY)/float64(s.Height-s.PaddingY*2))))
		if s.Dirty {
			s.Dispatch(event.NewEvent(ValueChangedEvent, s))
		}
	}

	return s.Node.Update(time)
}

func (s *Slider) Clear() {
	if s.Dirty {
		s.Dirty = false
		s.Image.Clear()
		vector.DrawFilledRect(s.Image, 0, 0, float32(s.Width), float32(s.Height), theme.Colors.BackgroundInverted, false)

		// Draw marks
		markWidth := float32(s.Width - s.PaddingX*2)
		markHeight := float32(s.Height-s.PaddingY*2-(s.Marks-1)*s.MarkMargin) / float32(s.Marks)
		markOnTrigger := float32(s.Marks) - float32(s.Value-s.ValueMin)/float32(s.ValueMax-s.ValueMin)*float32(s.Marks)
		for i := 0; i < s.Marks; i++ {
			col := theme.Colors.Off
			if float32(i) >= markOnTrigger {
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
	value = math.Min(math.Max(value, s.ValueMin), s.ValueMax)
	if value != s.Value {
		s.Value = value
		s.Dirty = true
	}
}

func (s *Slider) GetValue() float64 {
	return s.Value
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
