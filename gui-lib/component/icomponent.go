package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/control"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
)

type IComponent interface {
	control.ITarget
	event.IDispatcher

	GetChildren() []IComponent
	GetParent() IComponent
	GetRoot() IComponent
	SetParent(parent IComponent)
	Append(child IComponent)
	Prepend(child IComponent)
	Remove(child IComponent)
	MoveFront(child IComponent)

	Draw(dest *ebiten.Image)
	Update()

	GetGraphic() graphic.IGraphic
	GetLayout() layout.ILayout
}
