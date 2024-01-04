package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gosynth/event"
	"gosynth/gui-lib/graphic"
	"gosynth/gui-lib/layout"
	"image/color"
)

type FPS struct {
	*Component
}

func NewFPS() *FPS {
	f := &FPS{}
	f.Component = NewComponent()

	f.GetGraphic().AddListener(&f, graphic.DrawEvent, func(e event.IEvent) {
		img := f.GetGraphic().GetImage()
		img.Fill(color.Black)
		ebitenutil.DebugPrintAt(img, fmt.Sprintf("FPS %0.2f", ebiten.ActualFPS()), 2, 2)
	})

	f.GetLayout().SetAbsolutePositioning(true)
	f.GetLayout().SetSize(60, 20)

	return f
}

func (f *FPS) SetParent(p IComponent) {
	if op := f.GetParent(); op != nil {
		op.GetLayout().RemoveListener(&f, layout.UpdatedEvent)
	}

	p.GetLayout().AddListener(&f, layout.UpdatedEvent, func(e event.IEvent) {
		f.position()
	})

	f.position()

	f.Component.SetParent(p)
}

func (f *FPS) position() {
	if p := f.GetParent(); p != nil {
		pw, ph := p.GetLayout().GetSize()
		fw, fh := f.GetLayout().GetSize()
		f.GetLayout().SetPosition(pw-fw, ph-fh)
	}
}
