package widget

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gosynth/event"
	"gosynth/gui/component"
	"gosynth/gui/component/graphic"
	"gosynth/gui/component/layout"
	"image/color"
)

type FPS struct {
	*component.Component
}

func NewFPS() *FPS {
	f := &FPS{}
	f.Component = component.NewComponent()

	f.GetGraphic().GetDispatcher().AddListener(&f, graphic.DrawEvent, func(e event.IEvent) {
		img := f.GetGraphic().GetImage()
		img.Fill(color.Black)
		ebitenutil.DebugPrintAt(img, fmt.Sprintf("FPS %0.2f", ebiten.ActualFPS()), 2, 2)
	})

	f.GetLayout().SetAbsolutePositioning(true)
	f.GetLayout().GetSize().Set(60, 20)

	return f
}

func (f *FPS) SetParent(p component.IComponent) {
	if op := f.GetParent(); op != nil {
		op.GetLayout().GetDispatcher().RemoveListener(&f, layout.UpdatedEvent)
	}

	p.GetLayout().GetDispatcher().AddListener(&f, layout.UpdatedEvent, func(e event.IEvent) {
		f.position()
	})

	f.position()

	f.Component.SetParent(p)
}

func (f *FPS) position() {
	if p := f.GetParent(); p != nil {
		f.GetLayout().GetPosition().Set(
			p.GetLayout().GetSize().GetWidth()-f.GetLayout().GetSize().GetWidth(),
			p.GetLayout().GetSize().GetHeight()-f.GetLayout().GetSize().GetHeight(),
		)
	}
}
