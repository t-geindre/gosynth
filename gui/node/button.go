package node

import "github.com/hajimehoshi/ebiten/v2"

type Button struct {
	Node
}

func NewButton(width, height int, inode INode) *Button {
	b := &Button{}
	b.Children = make([]INode, 0)
	b.Options = &ebiten.DrawImageOptions{}
	b.INode = inode
	b.Resize(width, height)

	return b
}
