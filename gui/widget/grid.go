package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/behavior"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
)

type Grid struct {
	*component.Component
	cellW, cellH   float64
	shiftX, shiftY float64
}

func NewGrid(cellW, cellH float64) *Grid {
	r := &Grid{
		Component: component.NewComponent(),
		cellW:     cellW,
		cellH:     cellH,
	}

	r.GetLayout().SetFill(100)
	r.GetGraphic().AddListener(&r, graphic.DrawStartEvent, func(e event.IEvent) {
		r.GetGraphic().GetImage().Clear()
	})

	behavior.NewDraggable(r)
	r.AddListener(&r, behavior.DragEvent, func(e event.IEvent) {
		dragEvent := e.(*behavior.DragEventDetails)
		r.Shift(float64(dragEvent.DeltaX), float64(dragEvent.DeltaY))
	})

	return r
}

func (r *Grid) Shift(x, y float64) {
	if r.shiftX == x && r.shiftY == y {
		return
	}

	r.shiftX += x
	r.shiftY += y

	for _, c := range r.Component.GetChildren() {
		l := c.GetLayout()
		cx, cy := l.GetPosition()
		l.SetPosition(cx+x, cy+y)
	}
}

func (r *Grid) Append(c component.IComponent) {
	r.Component.Append(c)

	behavior.NewFocusable(c)
	behavior.NewDraggable(c)

	c.AddListener(&r, behavior.DragEvent, func(e event.IEvent) {
		r.onChildDrag(c, e)
	})

	c.AddListener(&r, behavior.DragStopEvent, func(e event.IEvent) {
		r.setComponentPosition(c)
	})

	c.GetLayout().SetAbsolutePositioning(true)
	r.setComponentPosition(c)

}

func (r *Grid) Remove(c component.IComponent) {
	c.RemoveListener(&r, behavior.DragEvent)
	c.RemoveListener(&r, behavior.DragStopEvent)
	r.Component.Remove(c)
}

func (r *Grid) onChildDrag(c component.IComponent, e event.IEvent) {
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		dEv := e.(*behavior.DragEventDetails)
		px, py := c.GetLayout().GetPosition()
		c.GetLayout().SetPosition(float64(dEv.DeltaX)+px, float64(dEv.DeltaY)+py)
		e.StopPropagation()
	}
}

func (r *Grid) setComponentPosition(c component.IComponent) {
	l := c.GetLayout()
	w, _ := l.GetSize()

	x, y := r.GetLayout().GetAbsolutePosition()
	mx, my := ebiten.CursorPosition()
	x, y = x+float64(mx), y+float64(my)

	shiftX := float64(int(r.shiftX) % int(r.cellW))
	shiftY := float64(int(r.shiftY) % int(r.cellH))

	// Compute cell position where the mouse is inside
	cx := float64(int(x/r.cellW))*r.cellW + shiftX - w + r.cellW
	cy := float64(int(y/r.cellH))*r.cellH + shiftY

	if x < cx {
		cx -= r.cellW
	}

	if y < cy {
		cy -= r.cellH
	}

	if y > cy+r.cellH {
		cy += r.cellH
	}

	if x > cx+r.cellW {
		cx += r.cellW
	}

	l.SetPosition(cx, cy)
}
