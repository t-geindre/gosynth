package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui/component/control"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/layout"
)

type IComponent interface {
	control.ITarget

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
	GetDispatcher() *event.Dispatcher
}
