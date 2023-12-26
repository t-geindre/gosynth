package node

import "github.com/hajimehoshi/ebiten/v2"

type INode interface {
	Append(child INode)
	Remove(child INode)
	RemoveAll()
	GetParent() INode
	Clear()
	Draw(dest *ebiten.Image)
	Update() error
	GetNodeAt(x, y int) INode
	SetPosition(x, y int)
	MoveBy(x, y int)
	Dispose()
	SetParent(parent INode)
	GetINode() INode
	MoveFront(child INode)
	MouseLeftDown(target INode)
	MouseLeftUp(target INode)
	MoveChildrenBy(x, y int)
}
