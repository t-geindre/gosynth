package primitive

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"gosynth/gui/theme"
)

func Title(title string, width, yOffset int) *ebiten.Image {
	height := theme.Fonts.Title.Metrics().CapHeight.Round() + yOffset
	img := ebiten.NewImage(width, height)
	img.Fill(theme.Colors.Background)
	w := font.MeasureString(theme.Fonts.Title, title)
	text.Draw(img, title, theme.Fonts.Title, (width-w.Round())/2, img.Bounds().Dy(), theme.Colors.Text)

	return img
}
