package graphic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
)

type Graphic struct {
	*event.Dispatcher
	parent          IGraphic
	children        []IGraphic
	options         *ebiten.DrawImageOptions
	image           *ebiten.Image
	width, height   int
	updateScheduled bool
	imageDirty      bool
}

func NewGraphic() *Graphic {
	g := &Graphic{
		children:   make([]IGraphic, 0),
		options:    &ebiten.DrawImageOptions{},
		imageDirty: true,
		Dispatcher: event.NewDispatcher(),
	}

	g.ScheduleUpdate()

	return g
}

func (g *Graphic) GetChildren() []IGraphic {
	return g.children
}

func (g *Graphic) GetParent() IGraphic {
	return g.parent
}

func (g *Graphic) SetParent(parent IGraphic) {
	g.parent = parent
}

func (g *Graphic) Append(child IGraphic) {
	g.children = append(g.children, child)
	child.SetParent(g)
}

func (g *Graphic) Prepend(child IGraphic) {
	g.children = append([]IGraphic{child}, g.children...)
	child.SetParent(g)
}

func (g *Graphic) Remove(child IGraphic) {
	for i, c := range g.children {
		if c == child {
			g.children = append(g.children[:i], g.children[i+1:]...)
		}
	}
	child.SetParent(nil)
}

func (g *Graphic) Draw(dest *ebiten.Image) {
	if g.width <= 0 || g.height <= 0 {
		return
	}

	if g.imageDirty {
		g.imageDirty = false

		if g.image != nil {
			g.image.Dispose()
		}

		if g.width > 0 && g.height > 0 {
			g.image = ebiten.NewImage(g.width, g.height)
		} else {
			return
		}
	}

	if g.updateScheduled {
		g.Dispatch(event.NewEvent(DrawUpdateRequiredEvent, g))
		g.updateScheduled = false
	}

	g.Dispatch(event.NewEvent(DrawStartEvent, g))

	// Todo images should all be drawn to a single image, merge matrices
	for _, child := range g.children {
		child.Draw(g.image)
	}

	g.Dispatch(event.NewEvent(DrawEndEvent, g))

	dest.DrawImage(g.image, g.options)
}

func (g *Graphic) Translate(x, y float64) {
	g.options.GeoM.Translate(x, y)
}

func (g *Graphic) SetTranslation(x, y float64) {
	g.options.GeoM.Reset()
	g.Translate(x, y)
}

func (g *Graphic) SetSize(width, height int) {
	if g.width == width && g.height == height {
		return
	}

	g.width = width
	g.height = height
	g.imageDirty = true

	g.ScheduleUpdate()
}

func (g *Graphic) ScheduleUpdate() {
	g.updateScheduled = true
}

func (g *Graphic) GetImage() *ebiten.Image {
	return g.image
}

func (g *Graphic) GetOptions() *ebiten.DrawImageOptions {
	return g.options
}

func (g *Graphic) MoveFront(child IGraphic) {
	for i, c := range g.children {
		if c == child {
			g.children = append(g.children[:i], g.children[i+1:]...)
			g.children = append(g.children, child)
			return
		}
	}
}
