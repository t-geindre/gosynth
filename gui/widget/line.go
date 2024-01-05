package widget

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
	"gosynth/gui/theme"
)

func NewLine(horizontal bool, width float32) *component.Component {
	l := component.NewComponent()

	if !horizontal {
		l.GetLayout().SetMargin(5, 5, 0, 0)
	} else {
		l.GetLayout().SetMargin(0, 0, 5, 5)
	}

	l.GetGraphic().AddListener(&l, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		img := l.GetGraphic().GetImage()
		img.Clear()

		w, h := l.GetLayout().GetSize()
		if horizontal {
			vector.StrokeLine(img, 0, float32(h/2), float32(w), float32(h/2), width, theme.Colors.BackgroundInverted, false)
		} else {
			vector.StrokeLine(img, float32(w/2), 0, float32(w/2), float32(h), width, theme.Colors.BackgroundInverted, false)
		}
	})

	return l
}
