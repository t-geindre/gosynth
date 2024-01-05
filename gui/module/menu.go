package module

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"gosynth/gui/connection"
	"gosynth/gui/theme"
	"gosynth/gui/widget"
)

type Menu struct {
	*component.Component
	rack *connection.Rack
}

func NewMenu(rack *connection.Rack) *Menu {
	m := &Menu{
		Component: component.NewComponent(),
	}

	l := m.GetLayout()
	l.SetWantedSize(0, 50)
	l.SetPadding(10, 10, 10, 10)
	l.SetContentOrientation(layout.Horizontal)

	m.GetGraphic().AddListener(&m, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		m.GetGraphic().GetImage().Fill(theme.Colors.Background)
		w, h := m.GetLayout().GetSize()
		img := m.GetGraphic().GetImage()
		vector.StrokeLine(img, 0, float32(h), float32(w), float32(h), 1, theme.Colors.BackgroundInverted, false)
	})

	logo := component.NewImage(theme.Images.Logo)
	logo.GetLayout().SetWantedSize(50, 0)
	logo.GetLayout().SetMargin(0, 0, 0, 3)
	m.Append(logo)
	m.Append(widget.NewTitle("Synth", widget.TitlePositionCenter))

	m.Append(component.NewFiller(100))

	btn := widget.NewButton("Add")
	btn.GetLayout().SetWantedSize(50, 50)
	m.Append(btn)

	menu := widget.NewDropdown()
	for _, op := range Registry.GetNames() {
		menu.AddOption(op, func(op string, rack *connection.Rack) func() {
			return func() {
				Registry.Get(op)(rack)
			}
		}(op, rack))
	}
	m.GetRoot().Append(menu)

	return m
}
