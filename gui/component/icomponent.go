package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/control"
	"gosynth/gui/graphic"
	"gosynth/gui/layout"
)

type IComponent interface {
	control.ITarget
	event.IDispatcher

	GetChildren() []IComponent
	GetParent() IComponent
	SetParent(parent IComponent)
	Append(child IComponent)
	Remove(child IComponent)
	MoveFront(child IComponent)

	Draw(dest *ebiten.Image)
	Update()

	GetGraphic() graphic.IGraphic
	GetLayout() layout.ILayout
}
