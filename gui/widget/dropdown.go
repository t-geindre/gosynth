package widget

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"gosynth/gui/theme"
	"math"
)

type Dropdown struct {
	*component.Component
	options map[string]func()
}

func NewDropdown() *Dropdown {
	d := &Dropdown{
		Component: component.NewComponent(),
	}

	d.GetLayout().SetAbsolutePositioning(true)
	d.GetLayout().ScheduleUpdate()

	d.GetLayout().AddListener(&d, layout.UpdateStartsEvent, func(e event.IEvent) {
		if p := d.GetParent(); p != nil {
			p.Remove(d)
			p.GetRoot().Append(d)
		}

		w, h := float64(0), float64(0)
		for _, c := range d.GetChildren() {
			cw, ch := c.GetLayout().GetWantedSize()
			if d.GetLayout().GetContentOrientation() == layout.Horizontal {
				w += cw
				h = math.Max(h, ch)
			} else {
				w = math.Max(w, cw)
				h += ch
			}
		}
		d.GetLayout().SetSize(w, h)
		d.GetLayout().SetPosition(300, 0)
	})

	d.GetGraphic().AddListener(&d, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		d.GetGraphic().GetImage().Fill(theme.Colors.Background)
	})

	return d
}

func (d *Dropdown) AddOption(label string, callback func()) {
	l := NewTitle(label, TitlePositionTop)
	l.AddListener(&d, control.LeftMouseUpEvent, func(e event.IEvent) {
		callback()
	})
	d.Append(l)
}
