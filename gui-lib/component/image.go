package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/graphic"
)

type Image struct {
	*Component
	img *ebiten.Image
}

func NewImage(img *ebiten.Image) *Image {
	i := &Image{
		Component: NewComponent(),
	}

	i.GetGraphic().AddListener(&i, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		// Todo this behavior could be moved to the graphic component,
		// allowing to factorize the code with the text component
		destImg := i.GetGraphic().GetImage()

		// Compute scaling factor
		scale := float64(1)
		scaleX := float64(destImg.Bounds().Dx()) / float64(i.img.Bounds().Dx())
		scaleY := float64(destImg.Bounds().Dy()) / float64(i.img.Bounds().Dy())

		if scaleX < 1 {
			scale = scaleX
		}

		if scaleY < 1 && scaleY < scaleX {
			scale = scaleY
		}

		// Center the temporary image according to scaling factor
		x := (destImg.Bounds().Dx() - int(float64(i.img.Bounds().Dx())*scale)) / 2
		y := (destImg.Bounds().Dy() - int(float64(i.img.Bounds().Dy())*scale)) / 2

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(x), float64(y))

		destImg.DrawImage(i.img, op)
	})

	i.SetImage(img)

	return i
}

func (i *Image) SetImage(img *ebiten.Image) {
	i.img = img
	// Todo this does not take scaling in account
	i.GetLayout().SetWantedSize(
		float64(img.Bounds().Dx()),
		float64(img.Bounds().Dy()),
	)
	i.GetGraphic().ScheduleUpdate()
}
