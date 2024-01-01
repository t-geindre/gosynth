package demo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/layout"
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

	r.GetGraphic().GetDispatcher().AddListener(&r, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		image := r.GetGraphic().GetImage()
		image.Fill(color.Black)
	})

	r.GetDispatcher().AddListener(&r, control.LeftMouseDownEvent, func(e event.IEvent) {
		r.mouseDrag.Start()
		ebiten.SetCursorShape(ebiten.CursorShapeMove)
		e.StopPropagation()
	})

	r.GetDispatcher().AddListener(&r, control.LeftMouseUpEvent, func(e event.IEvent) {
		r.mouseDrag.Stop()
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)

		e.StopPropagation()
	})

	r.GetLayout().GetSize().Set(10, 10)
	r.GetLayout().SetAbsolutePositioning(true)

	return r
}

func (r *Resize) SetParent(parent component.IComponent) {
	r.Component.SetParent(parent)

	// Todo we should make sure listener is removed on the previous parent

	if parent != nil {
		parent.GetLayout().GetDispatcher().AddListener(&r, layout.UpdatedEvent, func(e event.IEvent) {
			r.UpdatePosition()
		})
	}

	r.UpdatePosition()
}

func (r *Resize) Update() {
	if p := r.GetParent(); r.mouseDrag.IsActive() && p != nil {
		dx, dy := r.mouseDrag.GetDelta()
		pw, ph := p.GetLayout().GetSize().Get()

		nw := math.Max(r.minWidth, pw+float64(dx))
		nh := math.Max(r.minHeight, ph+float64(dy))

		p.GetLayout().GetSize().Set(nw, nh)
	}

	r.Component.Update()
}

func (r *Resize) UpdatePosition() {
	if r.GetParent() == nil {
		return
	}

	w, h := r.GetLayout().GetSize().Get()
	pw, ph := r.GetParent().GetLayout().GetSize().Get()

	r.GetLayout().GetPosition().Set(pw-w, ph-h)
}
