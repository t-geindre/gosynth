package module

import "gosynth/gui/component/widget"

type VCA struct {
	*Module
}

func NewVCA() *VCA {
	v := &VCA{}
	v.Module = NewModule("VCA", 1, v)
	slider := widget.NewSlider(0, 1, 25)
	slider.GetLayout().SetFill(100)
	v.Append(slider)
	return v
}
