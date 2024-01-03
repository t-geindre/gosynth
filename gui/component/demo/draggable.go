package demo

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/graphic"
	"image/color"
)

type Draggable struct {
	*component.Component
	mouseDrag *control.MouseDelta
	outerType component.IComponent
	mouseOver bool
}

func NewDraggable(outerType component.IComponent) *Draggable {
	d := &Draggable{
		Component: component.NewComponent(),
		outerType: outerType,
		mouseDrag: control.NewMouseDelta(),
	}

	if outerType == nil {
		d.outerType = d
	}

	l := d.GetLayout()
	l.SetAbsolutePositioning(true)

	c := randomColor()
	bCol := colorInverse(c)
	d.GetGraphic().AddListener(&d, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := d.GetGraphic().GetImage()
		image.Fill(c)
		vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 2, color.Black, false)
		if d.mouseOver {
			vector.StrokeRect(image, 0, 0, float32(image.Bounds().Dx()), float32(image.Bounds().Dy()), 20, bCol, false)
		}
	})

	d.	AddListener(&d, control.LeftMouseDownEvent, func(e event.IEvent) {
		d.mouseDrag.Start()
		d.GetGraphic().GetOptions().ColorScale.ScaleAlpha(0.95)
		e.StopPropagation()
	})

	d.	AddListener(&d, control.LeftMouseUpEvent, func(e event.IEvent) {
		d.mouseDrag.Stop()
		d.GetGraphic().GetOptions().ColorScale.Reset()
		e.StopPropagation()
	})

	d.	AddListener(&d, control.MouseEnterEvent, func(e event.IEvent) {
		d.mouseOver = true
		d.GetGraphic().ScheduleUpdate()
	})

	d.	AddListener(&d, control.MouseLeaveEvent, func(e event.IEvent) {
		d.mouseOver = false
		d.GetGraphic().ScheduleUpdate()
	})

	d.	AddListener(&d, control.FocusEvent, func(e event.IEvent) {
		if p := d.GetParent(); p != nil {
			p.MoveFront(d.outerType)
		}
	})

	return d
}

func (d *Draggable) Update() {
	if d.mouseDrag.IsActive() {
		dx, dy := d.mouseDrag.GetDelta()
		d.GetLayout().GetPosition().MoveBy(float64(dx), float64(dy))
	}

	d.Component.Update()
}
