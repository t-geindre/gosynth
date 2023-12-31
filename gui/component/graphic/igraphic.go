package graphic

import "github.com/hajimehoshi/ebiten/v2"

type IGraphic interface {
	GetChildren() []IGraphic
	GetParent() IGraphic
	SetParent(parent IGraphic)
	Append(child IGraphic)
	Remove(child IGraphic)
	MoveFront(child IGraphic)

	Draw(dest *ebiten.Image)

	Translate(x, y float64)
	SetTranslation(x, y float64)

	SetSize(width, height int)

	GetImage() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions

	ScheduleUpdate()
	SetUpdateFunc(func())

	Disable()
	Enable()
}
