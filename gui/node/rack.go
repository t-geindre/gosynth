package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"image/color"
	"time"
)

type Rack struct {
	*Node
	MouseLDown             bool
	LastMouseX, LastMouseY int
	BgColor                *color.RGBA
}

func NewRack(width, height int) *Rack {
	r := &Rack{}
	r.Node = NewNode(width, height, r)
	r.Node.DisableLayoutComputing()

	r.BgColor = &color.RGBA{R: 12, G: 12, B: 12, A: 255}

	r.Dispatcher.AddListener(&r, LeftMouseDownEvent, func(e event.IEvent) {
		r.MouseLeftDown(e.GetSource().(INode))
	})

	r.Dispatcher.AddListener(&r, LeftMouseUpEvent, func(e event.IEvent) {
		r.MouseLeftUp(e.GetSource().(INode))
	})

	return r
}

func (r *Rack) Update(time time.Duration) error {
	if r.MouseLDown {
		x, y := ebiten.CursorPosition()
		r.MoveChildrenBy(x-r.LastMouseX, y-r.LastMouseY)
		r.LastMouseX, r.LastMouseY = x, y
	}

	return r.Node.Update(time)
}

func (r *Rack) MouseLeftDown(target INode) {
	if r == target {
		r.MouseLDown = true
		r.LastMouseX, r.LastMouseY = ebiten.CursorPosition()
	}
}

func (r *Rack) MouseLeftUp(target INode) {
	r.MouseLDown = false
}

func (r *Rack) Clear() {
	r.Node.Clear()
	r.Node.Image.Fill(r.BgColor)
}
