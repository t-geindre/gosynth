package behavior

import (
	"gosynth/event"
	component2 "gosynth/gui-lib/component"
	control2 "gosynth/gui-lib/control"
)

type Draggable struct {
	drag *control2.MouseDelta
	node component2.IComponent
}

func NewDraggable(node component2.IComponent) *Draggable {
	d := &Draggable{
		drag: control2.NewMouseDelta(),
		node: node,
	}

	node.AddListener(&d, control2.LeftMouseDownEvent, func(e event.IEvent) {
		d.drag.Start()
		e.StopPropagation()
		node.Dispatch(event.NewEvent(DragStartEvent, node))
	})

	node.AddListener(&d, control2.LeftMouseUpEvent, func(e event.IEvent) {
		d.drag.Stop()
		e.StopPropagation()
		node.Dispatch(event.NewEvent(DragStopEvent, node))
	})

	node.AddListener(&d, component2.UpdateEvent, func(e event.IEvent) {
		if d.drag.IsActive() {
			dx, dy := d.drag.GetDelta()
			node.Dispatch(NewDragEvent(node, dx, dy))

			if !e.IsPropagationStopped() {
				px, py := d.node.GetLayout().GetPosition()
				d.node.GetLayout().SetPosition(float64(dx)+px, float64(dy)+py)
			}
		}
	})

	return d
}

func (d *Draggable) Remove() {
	d.node.RemoveListener(&d, control2.LeftMouseDownEvent)
	d.node.RemoveListener(&d, control2.LeftMouseUpEvent)
	d.node.RemoveListener(&d, component2.UpdateEvent)
}

type DragEventDetails struct {
	*event.Event
	DeltaX, DeltaY int
}

func NewDragEvent(src any, deltaX, deltaY int) *DragEventDetails {
	e := &DragEventDetails{
		Event:  event.NewEvent(DragEvent, src),
		DeltaX: deltaX,
		DeltaY: deltaY,
	}

	return e
}
