package control

import "github.com/hajimehoshi/ebiten/v2"

type cursor struct {
	stack []ebiten.CursorShapeType
}

func (c *cursor) Push(shape ebiten.CursorShapeType) {
	c.stack = append(c.stack, shape)
	ebiten.SetCursorShape(shape)
}

func (c *cursor) Pop() {
	if len(c.stack) == 0 {
		return
	}
	c.stack = c.stack[:len(c.stack)-1]
	if len(c.stack) == 0 {
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		return
	}
	ebiten.SetCursorShape(c.stack[len(c.stack)-1])
}

var Cursor = &cursor{}
