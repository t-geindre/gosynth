package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"gosynth/event"
	"gosynth/gui/graphic"
	"image/color"
)

type Text struct {
	*Component
	str  string
	w, h int
	font font.Face
}

func NewText(str string, fontFace font.Face, color, bgColor color.RGBA) *Text {
	t := &Text{
		Component: NewComponent(),
		font:      fontFace,
	}

	t.SetText(str)

	t.GetGraphic().AddListener(&t, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		// Write on a new image to allow scale down if required
		img := ebiten.NewImage(t.w, t.h)
		img.Fill(bgColor)
		text.Draw(img, t.str, t.font, 0, t.h, color)

		destImg := t.GetGraphic().GetImage()
		destImg.Clear()

		// Compute scaling factor
		scale := float64(1)
		scaleX := float64(destImg.Bounds().Dx()) / float64(img.Bounds().Dx())
		scaleY := float64(destImg.Bounds().Dy()) / float64(img.Bounds().Dy())

		if scaleX < 1 {
			scale = scaleX
		}

		if scaleY < 1 && scaleY < scaleX {
			scale = scaleY
		}

		// Center the temporary image according to scaling factor
		x := (destImg.Bounds().Dx() - int(float64(img.Bounds().Dx())*scale)) / 2
		y := (destImg.Bounds().Dy() - int(float64(img.Bounds().Dy())*scale)) / 2

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(x), float64(y))

		destImg.DrawImage(img, op)
	})

	return t
}

func (t *Text) SetText(str string) {
	t.str = str
	t.w = font.MeasureString(t.font, t.str).Round()
	t.h = t.font.Metrics().CapHeight.Round()
	t.GetLayout().SetWantedSize(float64(t.w), float64(t.h))
	t.GetGraphic().ScheduleUpdate()
}
