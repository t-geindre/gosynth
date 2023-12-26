package node

import (
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"gosynth/gui/theme"
)

type Label struct {
	Node
	Text     string
	Font     font.Face
	Inverted bool
	Dirty    bool
}

func NewLabel(width, height int, text string, font font.Face) *Label {
	t := &Label{}
	t.Node = *NewNode(width, height, t)
	t.Font = font
	t.Inverted = false
	t.SetText(text)

	return t
}

func (t *Label) SetParent(parent INode) {
	t.Node.SetParent(parent)
	t.Inverted = IsNodeInverted(t)
}

func (t *Label) Clear() {
	if t.Dirty {
		t.Image.Clear()
		tw := font.MeasureString(t.Font, t.Text).Round()
		th := t.Font.Metrics().CapHeight.Round()

		tx, ty := (t.Width-tw)/2, (t.Height+th)/2

		// Draw a rect behind the text to allow correct blending
		bgColor := theme.Colors.Background
		if t.Inverted {
			bgColor = theme.Colors.BackgroundInverted
		}
		vector.DrawFilledRect(t.Image, float32(tx-3), float32(ty-th-3), float32(tw+6), float32(th+6), bgColor, false)

		color := theme.Colors.Text
		if t.Inverted {
			color = theme.Colors.TextInverted
		}
		text.Draw(t.Image, t.Text, t.Font, tx, ty, color)
	}

	t.Node.Clear()
}

func (t *Label) SetText(text string) {
	t.Text = text
	t.Dirty = true
}

func (t *Label) Targetable() bool {
	return false
}
