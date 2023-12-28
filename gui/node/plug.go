package node

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/gui/theme"
)

const plugRadius = 18

type Plug struct {
	*Node
	Led *LED
}

func NewPlug() *Plug {
	p := &Plug{}
	p.Node = NewNode(plugRadius*2+1, plugRadius*2+1, p)
	p.Led = NewLED(theme.Colors.LedInOn, theme.Colors.LedInOff)
	p.AppendWithOptions(p.Led, NewAppendOptions().HorizontallyCentered())
	return p
}

func (p *Plug) SetParent(parent INode) {
	p.Node.SetParent(parent)
}

func (p *Plug) Clear() {
	if p.Dirty {
		p.Led.SetPositionY(p.Height/2 - p.Led.Height/2)

		p.Image.Clear()

		// Draw a rect behind the plug to allow correct blending
		bgColor := theme.Colors.Background
		if IsNodeInverted(p) {
			bgColor = theme.Colors.BackgroundInverted
		}
		vector.DrawFilledRect(p.Image, float32(p.Width/2-plugRadius-3), float32(p.Height/2-plugRadius-3), float32(plugRadius*2+6), float32(plugRadius*2+6), bgColor, false)

		vector.StrokeCircle(p.Image, float32(p.Width/2), plugRadius, plugRadius, 1, theme.Colors.Off, true)
		vector.DrawFilledCircle(p.Image, float32(p.Width/2), plugRadius, plugRadius-1, theme.Colors.Background, true)
		vector.StrokeCircle(p.Image, float32(p.Width/2), plugRadius, plugRadius-5, 1, theme.Colors.Off, true)

		p.Dirty = false
	}

	p.Node.Clear()
}

func (p *Plug) On() {
	p.Led.On()
}

func (p *Plug) Off() {
	p.Led.Off()
}
