package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/behavior"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui/theme"
	"math"
)

type Knob struct {
	*component.Image
}

func NewKnob() *Knob {
	k := &Knob{
		Image: component.NewImage(theme.Images.Knob),
	}

	behavior.NewDraggable(k)

	k.AddListener(&k, behavior.DragStartEvent, func(e event.IEvent) {
		control.Cursor.Push(ebiten.CursorShapeEWResize)
	})

	k.AddListener(&k, behavior.DragStopEvent, func(e event.IEvent) {
		control.Cursor.Pop()
	})

	k.AddListener(&k, behavior.DragEvent, func(e event.IEvent) {
		dragEvent := e.(*behavior.DragEventDetails)
		k.Image.Rotate(float64(dragEvent.DeltaX) * math.Pi / 180)
		e.StopPropagation()
	})

	return k
}
