package widget

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/graphic"
	"gosynth/gui/layout"
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
	l.GetWantedSize().SetHeight(50)
	l.GetPadding().SetAll(5)
	l.SetContentOrientation(layout.Horizontal)

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		m.GetGraphic().GetImage().Fill(theme.Colors.Background)
		w, h := m.GetLayout().GetSize().Get()
		img := m.GetGraphic().GetImage()
		vector.StrokeLine(img, 0, float32(h), float32(w), float32(h), 1, theme.Colors.BackgroundInverted, false)
	})

	m.Append(NewText("Synth", theme.Fonts.Title))

	m.Append(NewFiller(100))

	s := NewSlider(0, 1, 25)
	s.GetLayout().SetContentOrientation(layout.Horizontal)
	s.GetLayout().GetWantedSize().SetWidth(300)
	s.SetValue(1)
	m.Append(s)

	return m
}
