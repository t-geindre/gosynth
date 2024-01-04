package widget

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"gosynth/gui/theme"
)

type Menu struct {
	*component.Component
}

func NewMenu() *Menu {
	m := &Menu{
		Component: component.NewComponent(),
	}

	l := m.GetLayout()
	l.SetWantedSize(0, 50)
	l.SetPadding(5, 5, 5, 5)
	l.SetContentOrientation(layout.Horizontal)

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		m.GetGraphic().GetImage().Fill(theme.Colors.Background)
		w, h := m.GetLayout().GetSize()
		img := m.GetGraphic().GetImage()
		vector.StrokeLine(img, 0, float32(h), float32(w), float32(h), 1, theme.Colors.BackgroundInverted, false)
	})

	m.Append(NewTitle("Gosynth", TitlePositionCenter))

	m.Append(component.NewFiller(100))

	s := NewSlider(0, 1, 25)
	s.GetLayout().SetContentOrientation(layout.Horizontal)
	s.GetLayout().SetWantedSize(300, 0)
	s.SetValue(1)
	m.Append(s)

	return m
}
