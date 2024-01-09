package graphic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
)

type IGraphic interface {
	event.IDispatcher

	GetChildren() []IGraphic
	GetParent() IGraphic
	SetParent(parent IGraphic)
	Append(child IGraphic)
	Prepend(child IGraphic)
	Remove(child IGraphic)
	MoveFront(child IGraphic)

	Draw(dest *ebiten.Image)

	Translate(x, y float64)
	SetTranslation(x, y float64)

	SetSize(width, height int)

	GetImage() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions

	// ScheduleUpdate
	// Will trigger the DrawUpdateRequiredEvent on next graphic update
	ScheduleUpdate()
}
