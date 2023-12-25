package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Rack struct {
	Node
	MouseLDown             bool
	LastMouseX, LastMouseY int
}

func NewRack(width, height int) *Rack {
	r := &Rack{}
	r.Node = *NewNode(width, height, r)

	return r
}

func (r *Rack) Update() error {
	if r.MouseLDown {
		x, y := ebiten.CursorPosition()
		r.MoveChildrenBy(x-r.LastMouseX, y-r.LastMouseY)
		r.LastMouseX, r.LastMouseY = x, y
	}

	return r.Node.Update()
}
func (r *Rack) MouseLeftDown() {
	r.MouseLDown = true
	r.LastMouseX, r.LastMouseY = ebiten.CursorPosition()
}

func (r *Rack) MouseLeftUp() {
	r.MouseLDown = false
}
