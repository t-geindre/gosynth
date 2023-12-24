package gui

import "github.com/hajimehoshi/ebiten/v2"

type INode interface {
	Append(child INode)
	Remove(child INode)
	Clear()
	GetParent() INode
	Draw(dest *ebiten.Image)
	Update() error
	GetNodeAt(x, y int) INode
	SetPosition(x, y int)
}
