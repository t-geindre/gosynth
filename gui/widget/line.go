package widget

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/graphic"
	"gosynth/gui/theme"
)

func NewLine(horizontal bool, width float32) *component.Component {
	l := component.NewComponent()

	l.GetGraphic().AddListener(&l, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := l.GetGraphic().GetImage()
		img.Clear()

		w, h := l.GetLayout().GetSize()
		if horizontal {
			vector.StrokeLine(img, 0, float32(h/2), float32(h/2), float32(w), width, theme.Colors.BackgroundInverted, false)
		} else {
			vector.StrokeLine(img, float32(w/2), 0, float32(w/2), float32(h), width, theme.Colors.BackgroundInverted, false)
		}
	})

	return l
}
