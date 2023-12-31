package demo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/layout"
	"image/color"
	"math"
)

type Resize struct {
	*component.Component
	mouseDrag           *control.MouseDelta
	minWidth, minHeight int
}

func NewResize(minWidth, minHeight int) *Resize {
	r := &Resize{
		Component: component.NewComponent(),
		mouseDrag: control.NewMouseDelta(),
		minWidth:  minWidth,
		minHeight: minHeight,
	}

	r.GetGraphic().SetUpdateFunc(func() {
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

		nw := int(math.Max(float64(r.minWidth), float64(pw+dx)))
		nh := int(math.Max(float64(r.minHeight), float64(ph+dy)))

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
