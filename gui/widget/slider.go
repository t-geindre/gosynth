package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"gosynth/gui/connection"
	"gosynth/gui/theme"
	audio "gosynth/module"
)

type Slider struct {
	*component.Component
	from, to    float64
	marksCount  int
	marksOn     int
	value       float64
	mouseDown   bool
	remoteValue *connection.Value
}

func NewSlider(marks int, module audio.IModule, port audio.Port) *Slider {
	s := &Slider{
		Component:   component.NewComponent(),
		marksCount:  marks,
		remoteValue: connection.NewValue(0, 10, module, port),
	}

	s.GetGraphic().AddListener(&s, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := s.GetGraphic().GetImage()
		img.Fill(theme.Colors.BackgroundInverted)
	})

	s.GetLayout().SetPadding(5, 5, 5, 5)
	s.addMarks()

	s.AddListener(&s, control.LeftMouseDownEvent, s.onMouseDown)
	s.AddListener(&s, control.LeftMouseUpEvent, s.onMouseUp)
	s.AddListener(&s, component.UpdateEvent, s.onUpdate)

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

func (s *Slider) onUpdate(_ event.IEvent) {
	v := s.remoteValue.ReceiveAudioValue()

	if v != nil {
		s.SetValue(*v)
	}

	if s.mouseDown {
		mx, my := ebiten.CursorPosition()
		sx, sy := s.GetLayout().GetAbsolutePosition()
		w, h := s.GetLayout().GetSize()
		if s.GetLayout().GetContentOrientation() == layout.Horizontal {
			s.SetValue(0 + (10-0)*(float64(mx)-sx)/w)
		} else {
			s.SetValue(0 + (10-0)*(sy+h-float64(my))/h)
		}
	}
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
	if value == s.value {
		return
	}

	if value < 0 {
		value = 0
	} else if value > 10 {
		value = 10
	}

	s.value = value
	s.marksOn = int(value / (10 - 0) * float64(s.marksCount))

	s.remoteValue.SendGuiValue(value)

	s.updateMarks()
}
