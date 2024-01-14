package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/behavior"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
)

type Grid struct {
	*component.Component
	cellW, cellH   float64
	shiftX, shiftY float64
	dragGroup      []component.IComponent
}

func NewGrid(cellW, cellH float64) *Grid {
	g := &Grid{
		Component: component.NewComponent(),
		cellW:     cellW,
		cellH:     cellH,
		dragGroup: make([]component.IComponent, 0),
	}

	g.GetLayout().SetFill(100)
	g.GetGraphic().AddListener(&g, graphic.DrawUpdateRequiredEvent, func(e event.IEvent) {
		g.GetGraphic().GetImage().Clear()
	})

	behavior.NewDraggable(g)
	g.AddListener(&g, behavior.DragEvent, func(e event.IEvent) {
		dragEvent := e.(*behavior.DragEventDetails)
		g.Shift(float64(dragEvent.DeltaX), float64(dragEvent.DeltaY))
		g.GetGraphic().ScheduleUpdate()
	})

	return g
}

func (g *Grid) Shift(x, y float64) {
	if g.shiftX == x && g.shiftY == y {
		return
	}

	g.shiftX += x
	g.shiftY += y

	for _, c := range g.Component.GetChildren() {
		l := c.GetLayout()
		cx, cy := l.GetPosition()
		l.SetPosition(cx+x, cy+y)
	}
}

func (g *Grid) Append(c component.IComponent) {
	g.Component.Append(c)

	behavior.NewFocusable(c)
	behavior.NewDraggable(c)

	c.AddListener(&g, behavior.DragStartEvent, func(e event.IEvent) {
		g.dragGroup = append(g.dragGroup, c)

		if !ebiten.IsKeyPressed(ebiten.KeyControl) {
			return
		}

		x, y := c.GetLayout().GetPosition()
		w, _ := c.GetLayout().GetSize()
		x += w

		for found := true; found; {
			found = false
			for _, c := range g.Component.GetChildren() {
				l := c.GetLayout()
				cx, cy := l.GetPosition()
				if cx == x && cy == y {
					found = true
					g.dragGroup = append(g.dragGroup, c)
					w, _ := c.GetLayout().GetSize()
					x += w
					break
				}
			}
		}
	})

	c.AddListener(&g, behavior.DragEvent, g.onChildDrag)

	c.AddListener(&g, behavior.DragStopEvent, func(e event.IEvent) {
		for _, c := range g.dragGroup {
			g.setComponentPosition(c)
		}
		g.dragGroup = make([]component.IComponent, 0)
	})

	x, y := ebiten.CursorPosition()
	c.GetLayout().SetAbsolutePositioning(true)
	c.GetLayout().SetPosition(float64(x), float64(y))

	g.setComponentPosition(c)
}

func (g *Grid) Remove(c component.IComponent) {
	c.RemoveListener(&g, behavior.DragEvent)
	c.RemoveListener(&g, behavior.DragStopEvent)
	g.Component.Remove(c)
}

func (g *Grid) onChildDrag(e event.IEvent) {
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		dEv := e.(*behavior.DragEventDetails)
		for _, c := range g.dragGroup {
			px, py := c.GetLayout().GetPosition()
			c.GetLayout().SetPosition(float64(dEv.DeltaX)+px, float64(dEv.DeltaY)+py)
		}
		e.StopPropagation()
		g.GetGraphic().ScheduleUpdate()
	}
}

func (g *Grid) setComponentPosition(c component.IComponent) {
	l := c.GetLayout()

	shiftX := float64(int(g.shiftX) % int(g.cellW))
	shiftY := float64(int(g.shiftY) % int(g.cellH))

	x, y := l.GetPosition()
	if x < 0 {
		x -= g.cellW / 2
	} else {
		x += g.cellW / 2
	}
	if y < 0 {
		y -= g.cellH / 2
	} else {
		y += g.cellH / 2
	}

	x -= shiftX
	y -= shiftY

	cx := float64(int(x/g.cellW))*g.cellW + shiftX
	cy := float64(int(y/g.cellH))*g.cellH + shiftY

	l.SetPosition(cx, cy)

	g.GetGraphic().ScheduleUpdate()
}
