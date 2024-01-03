package widget

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/theme"
)

func NewLine(horizontal bool, width float32) *component.Component {
	l := component.NewComponent()

	l.GetGraphic().AddListener(&l, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := l.GetGraphic().GetImage()
		img.Clear()

		if horizontal {
			y := float32(l.GetLayout().GetSize().GetHeight()) / 2
			vector.StrokeLine(img, 0, y, float32(l.GetLayout().GetSize().GetWidth()), y, width, theme.Colors.BackgroundInverted, false)
		} else {
			x := float32(l.GetLayout().GetSize().GetWidth()) / 2
			vector.StrokeLine(img, x, 0, x, float32(l.GetLayout().GetSize().GetHeight()), width, theme.Colors.BackgroundInverted, false)
		}
	})

	return l
}
