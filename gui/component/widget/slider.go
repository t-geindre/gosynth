package widget

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/graphic"
	"gosynth/gui/theme"
)

type Slider struct {
	*component.Component
	from, to float64
	marks    int
}

func NewSlider(from, to float64, marks int) *Slider {
	s := &Slider{
		Component: component.NewComponent(),
		from:      from,
		to:        to,
		marks:     marks,
	}

	s.GetGraphic().GetDispatcher().AddListener(&s, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := s.GetGraphic().GetImage()
		img.Fill(theme.Colors.BackgroundInverted)
	})

	s.GetLayout().GetPadding().SetAll(5)

	for i := 0; i < marks; i++ {
		m := NewContainer()

		ml := m.GetLayout()
		fmt.Println(int(100 / float64(marks)))
		ml.SetFill(100 / float64(marks))

		if i > 0 {
			ml.GetMargin().SetTop(5)
		}

		mg := m.GetGraphic()
		mg.GetDispatcher().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
			img := mg.GetImage()
			img.Fill(theme.Colors.Off)
		})

		s.Append(m)
	}

	s.GetDispatcher().AddListener(&s, control.LeftMouseDownEvent, func(e event.IEvent) {
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
		e.StopPropagation()
	})

	s.GetDispatcher().AddListener(&s, control.LeftMouseUpEvent, func(e event.IEvent) {
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		e.StopPropagation()
	})

	return s
}
