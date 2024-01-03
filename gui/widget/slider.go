package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/control"
	"gosynth/gui/graphic"
	layout2 "gosynth/gui/layout"
	"gosynth/gui/theme"
)

type Slider struct {
	*component.Component
	from, to   float64
	marksCount int
	marksOn    int
	value      float64
	mouseDown  bool
}

func NewSlider(from, to float64, marks int) *Slider {
	s := &Slider{
		Component:  component.NewComponent(),
		from:       from,
		to:         to,
		marksCount: marks,
	}

	s.GetGraphic().AddListener(&s, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := s.GetGraphic().GetImage()
		img.Fill(theme.Colors.BackgroundInverted)
	})

	s.GetLayout().GetPadding().SetAll(5)
	s.addMarks()

	s.AddListener(&s, control.LeftMouseDownEvent, s.onMouseDown)
	s.AddListener(&s, control.LeftMouseUpEvent, s.onMouseUp)

	return s
}

func (s *Slider) addMarks() {
	for i := 0; i < s.marksCount; i++ {
		m := NewContainer()
		m.GetLayout().SetFill(100/float64(s.marksCount) - 1)

		mg := m.GetGraphic()
		mg.AddListener(&m, graphic.DrawUpdateRequiredEvent, func(index int) func(e event.IEvent) {
			return func(e event.IEvent) {
				i := index
				if s.GetLayout().GetContentOrientation() == layout2.Vertical {
					i = s.marksCount - index - 1
				}
				img := mg.GetImage()
				if i < s.marksOn {
					img.Fill(theme.Colors.On)
				} else {
					img.Fill(theme.Colors.Off)
				}
			}
		}(i))

		s.Append(m)
	}

	s.GetLayout().AddListener(&s, layout2.UpdateStartsEvent, func(e event.IEvent) {
		for i, m := range s.GetChildren() {
			if i > 0 {
				if s.GetLayout().GetContentOrientation() == layout2.Horizontal {
					m.GetLayout().GetMargin().SetLeft(2)
					continue
				}
				m.GetLayout().GetMargin().SetTop(2)
			}
		}
	})
}

func (s *Slider) updateMarks() {
	for _, m := range s.GetChildren() {
		m.GetGraphic().ScheduleUpdate()
	}
}

func (s *Slider) Update() {
	if s.mouseDown {
		mx, my := ebiten.CursorPosition()
		sx, sy := s.GetLayout().GetAbsolutePosition().Get()
		if s.GetLayout().GetContentOrientation() == layout2.Horizontal {
			s.SetValue(s.from + (s.to-s.from)*(float64(mx)-sx)/s.GetLayout().GetSize().GetWidth())
		} else {
			s.SetValue(s.from + (s.to-s.from)*(sy+s.GetLayout().GetSize().GetHeight()-float64(my))/s.GetLayout().GetSize().GetHeight())
		}
	}

	s.Component.Update()
}

func (s *Slider) onMouseDown(e event.IEvent) {
	ebiten.SetCursorShape(ebiten.CursorShapePointer)
	s.mouseDown = true
	e.StopPropagation()
}

func (s *Slider) onMouseUp(e event.IEvent) {
	ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	s.mouseDown = false
	e.StopPropagation()
}

func (s *Slider) SetValue(value float64) {
	if value < s.from {
		value = s.from
	} else if value > s.to {
		value = s.to
	}

	s.value = value
	s.marksOn = int(value / (s.to - s.from) * float64(s.marksCount))

	s.updateMarks()
}