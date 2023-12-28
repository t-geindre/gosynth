package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"time"
)

type INode interface {
	// Tree management
	Append(child INode)
	AppendWithOptions(child INode, options *AppendOptions)
	Remove(child INode)
	RemoveAll()
	GetChildren() []INode
	GetParent() INode
	MoveFront(child INode)
	MoveChildrenBy(x, y int)
	SetParent(parent INode)
	GetINode() INode

	// Life cycle
	Clear()
	Draw(dest *ebiten.Image)
	Update(time time.Duration) error
	Dispose()

	// Targetting
	GetTargetNodeAt(x, y int) INode
	Targetable() bool

	// Positioning
	SetPosition(x, y int)
	SetPositionX(x int)
	SetPositionY(y int)
	GetPosition() (int, int)
	GetPositionX() int
	GetPositionY() int
	GetAbsolutePosition() (int, int)
	MoveBy(x, y int)
	SetAppendOptions(options *AppendOptions)
	GetAppendOptions() *AppendOptions
	HCenter()

	// Sizing
	SetSize(width, height int)
	SetWidth(width int)
	SetHeight(height int)
	GetSize() (int, int)
	GetOuterSize() (int, int)
	GetWidth() int
	GetOuterWidth() int
	GetHeight() int
	GetOuterHeight() int

	event.IDispatcher
}
