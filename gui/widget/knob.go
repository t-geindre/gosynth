package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/behavior"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/control"
	"gosynth/gui/connection"
	"gosynth/gui/theme"
	audio "gosynth/module"
)

type Knob struct {
	*component.Image
	remoteValue *connection.Value
}

func NewKnob(module audio.IModule, port audio.Port) *Knob {
	k := &Knob{
		Image:       component.NewImage(theme.Images.Knob),
		remoteValue: connection.NewValue(-90, 90, module, port),
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
		drag := float64(dragEvent.DeltaX)
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			drag /= 10
		}
		k.Image.Rotate(drag)

		if k.Image.GetRotation() > 90 {
			k.Image.SetRotation(90)
		}

		if k.Image.GetRotation() < -90 {
			k.Image.SetRotation(-90)
		}

		k.remoteValue.SendGuiValue(k.Image.GetRotation())

		e.StopPropagation()
	})

	k.remoteValue.SendGuiValue(k.Image.GetRotation())

	return k
}
