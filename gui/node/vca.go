package node

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gosynth/gui/theme"
)

type VCA struct {
	Module
}

func NewVCA() *VCA {
	v := &VCA{}
	width, height := 65, 500
	v.Module = *NewModule(width, height, v)

	// Components
	sl := NewSlider(width-20, 200)
	sl.SetPosition(10, 35)
	sl.SetRange(0, 1)
	sl.SetValue(0.5)
	v.Append(sl)

	vector.StrokeLine(v.Image, float32(width/2), float32(237), float32(width/2), 243, 1, theme.Colors.Off, false)

	cvPl := NewPlug()
	cvPl.SetPosition(0, 245)
	v.Append(cvPl)
	cvPl.HCenter()

	lb := NewLabel(width, 10, "CV", theme.Fonts.Small)
	lb.SetPosition(0, 287)
	v.Append(lb)

	inLb := NewLabel(width-20, 10, "IN", theme.Fonts.Small)
	inLb.SetPosition(0, height-122)
	v.Append(inLb)
	inLb.HCenter()

	inPl := NewPlug()
	inPl.SetPosition(0, height-110)
	v.Append(inPl)
	inPl.HCenter()

	vector.StrokeLine(v.Image, float32(width/2), float32(height-60), float32(width/2), float32(height-90), 1, theme.Colors.Off, false)

	iv := NewInverted(width-20, 48)
	iv.SetPosition(10, height-58)
	v.Append(iv)

	outPl := NewPlug()
	outPl.SetPosition(0, 5)
	iv.Append(outPl)
	outPl.HCenter()

	return v

}
