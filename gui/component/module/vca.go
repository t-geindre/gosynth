package module

import (
	"gosynth/gui/component/demo"
)

type VCA struct {
	*Module
}

func NewVCA() *VCA {
	v := &VCA{}
	v.Module = NewModule(v)

	for i := 0; i < 4; i++ {
		btn := demo.NewButton()
		btn.GetLayout().GetSize().Set(50, 50)
		btn.GetLayout().GetMargin().SetBottom(10)
		v.Append(btn)
	}

	v.GetLayout().GetSize().Set(ModuleUWidth*3, ModuleHeight)

	return v
}
