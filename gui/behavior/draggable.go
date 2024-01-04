package behavior

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/control"
)

type Draggable struct {
	drag *control.MouseDelta
	node component.IComponent
}

func NewDraggable(node component.IComponent) *Draggable {
	d := &Draggable{
		drag: control.NewMouseDelta(),
		node: node,
	}

	node.AddListener(&d, control.LeftMouseDownEvent, func(e event.IEvent) {
		d.drag.Start()
		e.StopPropagation()
		node.Dispatch(event.NewEvent(DragStartEvent, node))
	})

	node.AddListener(&d, control.LeftMouseUpEvent, func(e event.IEvent) {
		d.drag.Stop()
		e.StopPropagation()
		node.Dispatch(event.NewEvent(DragStopEvent, node))
	})

	node.AddListener(&d, component.UpdateEvent, func(e event.IEvent) {
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
	d.node.RemoveListener(&d, control.LeftMouseDownEvent)
	d.node.RemoveListener(&d, control.LeftMouseUpEvent)
	d.node.RemoveListener(&d, component.UpdateEvent)
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
