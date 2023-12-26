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
	GetTargetNodeAt(x, y int) INode
	Targetable() bool
	SetPosition(x, y int)
	MoveBy(x, y int)
	GetPosition() (int, int)
	GetAbsolutePosition() (int, int)
	GetSize() (int, int)
	Dispose()
	SetParent(parent INode)
	GetINode() INode
	MoveFront(child INode)
	MouseLeftDown(target INode)
	MouseLeftUp(target INode)
	MoveChildrenBy(x, y int)
	HCenter()
}
