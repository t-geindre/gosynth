package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/theme"
)

type Text struct {
	*component.Component
	str        string
	strw, strh int
	font       font.Face
	inverted   bool
}

func NewText(str string, fontFace font.Face) *Text {
	t := &Text{
		Component: component.NewComponent(),
		font:      fontFace,
	}

	t.SetText(str)

	t.GetGraphic().AddListener(&t, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		// Write on a new image to allow scale down if required
		img := ebiten.NewImage(t.strw, t.strh)

		// Draw a rect behind the str to allow correct blending
		bgColor := theme.Colors.Background
		if t.inverted {
			bgColor = theme.Colors.BackgroundInverted
		}
		img.Fill(bgColor)

		color := theme.Colors.Text
		if t.inverted {
			color = theme.Colors.TextInverted
		}

		text.Draw(img, t.str, t.font, 0, t.strh, color)

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

	bound, _ := font.BoundString(t.font, t.str)
	t.strw = (bound.Max.X - bound.Min.X).Round()
	t.strh = (bound.Max.Y - bound.Min.Y).Round()

	t.GetLayout().GetWantedSize().Set(float64(t.strw), float64(t.strh))

	t.GetGraphic().ScheduleUpdate()
}

func (t *Text) SetInverted(inverted bool) {
	t.inverted = inverted
	t.GetGraphic().ScheduleUpdate()
}
