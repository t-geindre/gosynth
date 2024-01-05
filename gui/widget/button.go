package widget

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
)

type Button struct {
	*component.Component
}

func NewButton(label string) *Button {
	b := &Button{
		Component: component.NewComponent(),
	}

	b.Append(NewTitle(label, TitlePositionCenter))

	b.GetGraphic().AddListener(&b, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {

	})

	return b
}
