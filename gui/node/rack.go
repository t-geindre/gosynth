package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
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
	r.BgColor = &color.RGBA{R: 12, G: 12, B: 12, A: 255}

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
