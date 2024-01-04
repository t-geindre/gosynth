package demo

import (
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/control"
	"gosynth/gui/graphic"
	"gosynth/gui/layout"
	"image/color"
	"math"
)

type Resize struct {
	*component.Component
	mouseDrag           *control.MouseDelta
	minWidth, minHeight float64
}

func NewResize(minWidth, minHeight float64) *Resize {
	r := &Resize{
		Component: component.NewComponent(),
		mouseDrag: control.NewMouseDelta(),
		minWidth:  minWidth,
		minHeight: minHeight,
	}

	r.GetGraphic().AddListener(&r, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := r.GetGraphic().GetImage()
		image.Fill(color.Black)
	})

	r.AddListener(&r, control.LeftMouseDownEvent, func(e event.IEvent) {
		r.mouseDrag.Start()
		e.StopPropagation()
	})

	r.AddListener(&r, control.LeftMouseUpEvent, func(e event.IEvent) {
		r.mouseDrag.Stop()
		e.StopPropagation()
	})

	r.GetLayout().SetSize(10, 10)
	r.GetLayout().SetAbsolutePositioning(true)

	return r
}

func (r *Resize) SetParent(parent component.IComponent) {
	r.Component.SetParent(parent)

	// Todo we should make sure listener is removed on the previous parent

	if parent != nil {
		parent.GetLayout().AddListener(&r, layout.UpdatedEvent, func(e event.IEvent) {
			r.UpdatePosition()
		})
	}

	r.UpdatePosition()
}

func (r *Resize) Update() {
	if p := r.GetParent(); r.mouseDrag.IsActive() && p != nil {
		dx, dy := r.mouseDrag.GetDelta()
		pw, ph := p.GetLayout().GetSize()

		nw := math.Max(r.minWidth, pw+float64(dx))
		nh := math.Max(r.minHeight, ph+float64(dy))

		p.GetLayout().SetSize(nw, nh)
	}

	r.Component.Update()
}

func (r *Resize) UpdatePosition() {
	if r.GetParent() == nil {
		return
	}

	w, h := r.GetLayout().GetSize()
	pw, ph := r.GetParent().GetLayout().GetSize()

	r.GetLayout().SetPosition(pw-w, ph-h)
}
