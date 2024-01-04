package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/control"
	"gosynth/gui/graphic"
	"gosynth/gui/layout"
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

	s.GetLayout().SetPadding(5, 5, 5, 5)
	s.addMarks()

	s.AddListener(&s, control.LeftMouseDownEvent, s.onMouseDown)
	s.AddListener(&s, control.LeftMouseUpEvent, s.onMouseUp)

	return s
}

func (s *Slider) addMarks() {
	for i := 0; i < s.marksCount; i++ {
		m := component.NewContainer()
		m.GetLayout().SetFill(100/float64(s.marksCount) - 1)

		mg := m.GetGraphic()
		mg.AddListener(&m, graphic.DrawUpdateRequiredEvent, func(index int) func(e event.IEvent) {
			return func(e event.IEvent) {
				i := index
				if s.GetLayout().GetContentOrientation() == layout.Vertical {
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

	s.GetLayout().AddListener(&s, layout.UpdateStartsEvent, func(e event.IEvent) {
		for i, m := range s.GetChildren() {
			if i > 0 {
				if s.GetLayout().GetContentOrientation() == layout.Horizontal {
					m.GetLayout().SetMargin(0, 0, 2, 0)
					continue
				}
				m.GetLayout().SetMargin(2, 0, 0, 0)
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
		sx, sy := s.GetLayout().GetAbsolutePosition()
		w, h := s.GetLayout().GetSize()
		if s.GetLayout().GetContentOrientation() == layout.Horizontal {
			s.SetValue(s.from + (s.to-s.from)*(float64(mx)-sx)/w)
		} else {
			s.SetValue(s.from + (s.to-s.from)*(sy+h-float64(my))/h)
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
