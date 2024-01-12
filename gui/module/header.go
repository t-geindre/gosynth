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

type Header struct {
	*component.Component
	rack *connection.Rack
}

func NewHeader() *Header {
	h := &Header{
		Component: component.NewComponent(),
	}

	l := h.GetLayout()
	l.SetWantedSize(0, 50)
	l.SetPadding(10, 10, 10, 10)
	l.SetContentOrientation(layout.Horizontal)

	h.GetGraphic().AddListener(&h, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		h.GetGraphic().GetImage().Fill(theme.Colors.Background)
		wi, he := h.GetLayout().GetSize()
		img := h.GetGraphic().GetImage()
		vector.StrokeLine(img, 0, float32(he), float32(wi), float32(he), 1, theme.Colors.BackgroundInverted, false)
	})

	logo := component.NewImage(theme.Images.Logo)
	logo.GetLayout().SetWantedSize(50, 0)
	logo.GetLayout().SetMargin(0, 0, 0, 3)
	h.Append(logo)
	h.Append(widget.NewTitle("Synth", widget.TitlePositionCenter))

	h.Append(component.NewFiller(100))

	return h
}
