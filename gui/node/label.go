package node

import (
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"gosynth/gui/theme"
)

type Label struct {
	*Node
	Text     string
	Font     font.Face
	Inverted bool
}

func NewLabel(text string, font font.Face) *Label {
	t := &Label{}
	t.Font = font
	t.Inverted = false
	width, height := t.computeSize(text)
	t.Node = NewNode(width, height, t)
	t.SetText(text)

	return t
}

func (t *Label) Clear() {
	if t.Dirty {
		t.Inverted = IsNodeInverted(t)
		t.Image.Clear()

		// Draw a rect behind the text to allow correct blending
		bgColor := theme.Colors.Background
		if t.Inverted {
			bgColor = theme.Colors.BackgroundInverted
		}
		vector.DrawFilledRect(t.Image, 0, 0, float32(t.Width), float32(t.Height), bgColor, false)

		color := theme.Colors.Text
		if t.Inverted {
			color = theme.Colors.TextInverted
		}
		text.Draw(t.Image, t.Text, t.Font, 0, t.Height, color)
	}

	t.Node.Clear()
}

func (t *Label) SetText(text string) {
	t.Text = text
	t.Node.Dirty = true
}

func (t *Label) Targetable() bool {
	return false
}

func (t *Label) computeSize(text string) (int, int) {
	tw := font.MeasureString(t.Font, text).Round()
	th := t.Font.Metrics().CapHeight.Round()

	return tw, th
}
