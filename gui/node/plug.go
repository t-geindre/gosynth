package node

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/gui/theme"
	"image/color"
)

const plugRadius = 18
const shadowOffset = 3

type Plug struct {
	*Node
	Inverted bool
	Dirty    bool
}

func NewPlug() *Plug {
	p := &Plug{}
	p.Node = NewNode(plugRadius*2+shadowOffset+3, plugRadius*2+shadowOffset+3, p)
	p.Inverted = false
	p.Dirty = true

	return p
}

func (p *Plug) SetParent(parent INode) {
	p.Node.SetParent(parent)
	p.Inverted = IsNodeInverted(p)
}

func (p *Plug) Clear() {
	if p.Dirty {
		p.Image.Clear()

		// Draw a rect behind the plug to allow correct blending
		bgColor := theme.Colors.Background
		if p.Inverted {
			bgColor = theme.Colors.BackgroundInverted
		}
		vector.DrawFilledRect(p.Image, float32(p.Width/2-plugRadius-3), float32(p.Height/2-plugRadius-3), float32(plugRadius*2+6), float32(plugRadius*2+6), bgColor, false)

		if !p.Inverted {
			shadowColor := color.RGBA{A: 40}
			vector.DrawFilledCircle(p.Image, float32(p.Width/2), float32(p.Height/2)+1, plugRadius, shadowColor, true)
		}

		vector.StrokeCircle(p.Image, float32(p.Width/2), plugRadius+1, plugRadius, 1, theme.Colors.Off, true)
		vector.DrawFilledCircle(p.Image, float32(p.Width/2), plugRadius+1, plugRadius-1, theme.Colors.Background, true)
		vector.StrokeCircle(p.Image, float32(p.Width/2), plugRadius+1, plugRadius-5, 1, theme.Colors.Off, true)
		vector.DrawFilledCircle(p.Image, float32(p.Width/2), plugRadius+1, plugRadius-10, theme.Colors.BackgroundInverted, true)

		p.Dirty = false
	}

	p.Node.Clear()
}
