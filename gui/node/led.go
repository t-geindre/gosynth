package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/gui/theme"
	"image/color"
)

const LedRadius = 8

type LED struct {
	*Node
	IsOff        bool
	StateChanged bool
	ColorOn      *color.RGBA
	ImageOn      *ebiten.Image
	ColorOff     *color.RGBA
	ImageOff     *ebiten.Image
}

func NewLED(colorOn *color.RGBA, colorOff *color.RGBA) *LED {
	l := &LED{}
	l.Node = NewNode(LedRadius*2, LedRadius*2, l)
	l.ImageOn = ebiten.NewImage(LedRadius*2, LedRadius*2)
	l.ImageOff = ebiten.NewImage(LedRadius*2, LedRadius*2)
	l.StateChanged = true
	l.IsOff = true
	l.ColorOn = colorOn
	l.ColorOff = colorOff

	return l
}

func (l *LED) Clear() {
	if l.Dirty {
		l.ImageOn.Clear()
		l.ImageOff.Clear()

		// Draw a rect behind the plug to allow correct blending
		bgColor := theme.Colors.Background

		vector.DrawFilledCircle(l.ImageOff, float32(l.Width/2), float32(l.Height/2), LedRadius+1, bgColor, true)
		vector.DrawFilledCircle(l.ImageOn, float32(l.Width/2), float32(l.Height/2), LedRadius+1, bgColor, true)

		vector.DrawFilledCircle(l.ImageOff, float32(l.Width/2), float32(l.Height/2), LedRadius, l.ColorOff, true)
		vector.DrawFilledCircle(l.ImageOn, float32(l.Width/2), float32(l.Height/2), LedRadius, l.ColorOn, true)

		l.Dirty = false
	}

	if l.StateChanged {
		l.Image.Clear()
		if l.IsOff {
			l.Image.DrawImage(l.ImageOff, nil)
		} else {
			l.Image.DrawImage(l.ImageOn, nil)
		}
		l.StateChanged = false
	}

	l.Node.Clear()
}

func (l *LED) On() {
	l.IsOff = false
	l.StateChanged = true
}

func (l *LED) Off() {
	l.IsOff = true
	l.StateChanged = true
}
