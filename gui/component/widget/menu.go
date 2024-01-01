package widget

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/layout"
	"gosynth/gui/theme"
)

type Menu struct {
	*component.Component
}

func NewMenu() *Menu {
	m := &Menu{
		Component: component.NewComponent(),
	}

	m.GetLayout().GetWantedSize().SetHeight(50)
	m.GetLayout().GetPadding().SetAll(5)
	m.GetLayout().SetContentOrientation(layout.Horizontal)

	m.GetGraphic().GetDispatcher().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		m.GetGraphic().GetImage().Fill(theme.Colors.Background)
		w, h := m.GetLayout().GetSize().Get()
		img := m.GetGraphic().GetImage()
		vector.StrokeLine(img, 0, float32(h), float32(w), float32(h), 1, theme.Colors.BackgroundInverted, false)
	})

	m.Append(NewText("Synth", theme.Fonts.Title, theme.Colors.Text))
	m.Append(NewFiller(100))

	ctr := NewContainer()
	ctr.GetLayout().GetWantedSize().SetWidth(300)
	ctr.GetLayout().SetContentOrientation(layout.Vertical)

	m.Append(ctr)

	//ctr.Append(NewText("Master volume", theme.Fonts.Small, theme.Colors.Text))

	s := NewSlider(0, 1, 20)
	s.GetLayout().SetContentOrientation(layout.Horizontal)
	s.GetLayout().GetWantedSize().SetHeight(30)
	s.SetValue(1)
	ctr.Append(s)

	return m
}
