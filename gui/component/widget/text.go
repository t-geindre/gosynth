package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/theme"
	"image/color"
)

type Text struct {
	*component.Component
	str        string
	strw, strh int
	font       font.Face
	color      color.RGBA
}

func NewText(str string, fontFace font.Face, color color.RGBA) *Text {
	t := &Text{
		Component: component.NewComponent(),
		font:      fontFace,
		color:     color,
	}

	t.SetText(str)

	t.GetGraphic().GetDispatcher().AddListener(&t, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		// Write on a new image to allow scale down if required
		img := ebiten.NewImage(t.strw, t.strh)

		// Draw a rect behind the str to allow correct blending
		img.Fill(theme.Colors.Background)

		text.Draw(img, t.str, t.font, 0, t.strh, t.color)

		destImg := t.GetGraphic().GetImage()

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

	t.strw = font.MeasureString(t.font, t.str).Round()
	t.strh = t.font.Metrics().CapHeight.Round()

	t.GetLayout().GetWantedSize().Set(float64(t.strw), float64(t.strh))

	t.GetGraphic().ScheduleUpdate()
}
