package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/graphic"
	"gosynth/gui/theme"
)

type Plug struct {
	*component.Component
	inverted bool
}

func NewPlug() *Plug {
	p := &Plug{
		Component: component.NewComponent(),
	}

	l := p.GetLayout()
	l.SetWantedSize(
		float64(theme.Images.Plug.Bounds().Dx()),
		float64(theme.Images.Plug.Bounds().Dy()),
	)

	p.GetGraphic().AddListener(&p, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		w, h := p.GetLayout().GetSize()
		x := (w - float64(theme.Images.Plug.Bounds().Dx())) / 2
		y := (h - float64(theme.Images.Plug.Bounds().Dy())) / 2

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(x, y)

		if p.inverted {
			p.GetGraphic().GetImage().DrawImage(theme.Images.PlugInverted, options)
			return
		}

		p.GetGraphic().GetImage().DrawImage(theme.Images.Plug, options)
	})

	return p
}

func (p *Plug) SetInverted(inverted bool) {
	p.inverted = inverted
}
