package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
	"gosynth/gui-lib/control"
	graphic2 "gosynth/gui-lib/graphic"
	layout2 "gosynth/gui-lib/layout"
	"math"
)

type Component struct {
	*event.Dispatcher
	Layout   *layout2.Layout
	Graphic  *graphic2.Graphic
	Children []IComponent
	Parent   IComponent
}

func NewComponent() *Component {
	c := &Component{
		Layout:     layout2.NewLayout(),
		Graphic:    graphic2.NewGraphic(),
		Dispatcher: event.NewDispatcher(),
		Children:   make([]IComponent, 0),
	}

	c.GetLayout().AddListener(&c, layout2.ResizeEvent, func(e event.IEvent) {
		w, h := c.GetLayout().GetSize()
		c.GetGraphic().SetSize(int(math.Round(w)), int(math.Round(h)))
	})

	c.GetLayout().AddListener(&c, layout2.MoveEvent, func(e event.IEvent) {
		x, y := c.GetLayout().GetPosition()
		c.GetGraphic().SetTranslation(x, y)
	})

	return c
}

func (c *Component) GetChildren() []IComponent {
	return c.Children
}

func (c *Component) GetParent() IComponent {
	return c.Parent
}

func (c *Component) SetParent(parent IComponent) {
	c.Parent = parent
}

func (c *Component) Append(child IComponent) {
	if cComp, ok := child.(*Component); ok {
		if cComp == c {
			panic("cannot append component to itself")
		}
	}

	c.Children = append(c.Children, child)
	c.Graphic.Append(child.GetGraphic())
	c.Layout.Append(child.GetLayout())

	child.SetParent(c)
}

func (c *Component) Remove(child IComponent) {
	for i, ch := range c.Children {
		if ch == child {
			c.Children = append(c.Children[:i], c.Children[i+1:]...)
		}
	}

	c.Graphic.Remove(child.GetGraphic())
	c.Layout.Remove(child.GetLayout())

	child.SetParent(nil)
}

func (c *Component) GetLayout() layout2.ILayout {
	return c.Layout
}

func (c *Component) GetGraphic() graphic2.IGraphic {
	return c.Graphic
}

func (c *Component) Draw(dest *ebiten.Image) {
	c.Graphic.Draw(dest)
}

func (c *Component) Update() {
	c.Dispatch(event.NewEvent(UpdateEvent, c))

	for _, child := range c.Children {
		child.Update()
	}
}

func (c *Component) Dispatch(e event.IEvent) {
	c.Dispatcher.Dispatch(e)

	if !e.IsPropagationStopped() && c.Parent != nil {
		c.Parent.Dispatch(e)
	}
}

func (c *Component) GetTargetAt(x, y int) (control.ITarget, error) {
	var target *Component = nil
	if c.GetLayout().PointCollides(float64(x), float64(y)) {
		target = c

		cX, cY := c.GetLayout().GetPosition()

		x -= int(cX)
		y -= int(cY)

		// Range in reverse order so that the top-most child is checked first
		for i := len(c.Children) - 1; i >= 0; i-- {
			child := c.Children[i]
			if ct, err := child.GetTargetAt(x, y); err == nil {
				return ct, nil
			}
		}
	}

	if target == nil {
		return nil, fmt.Errorf("no target found")
	}

	return target, nil
}

func (c *Component) MoveFront(child IComponent) {
	for i, ch := range c.Children {
		if ch == child {
			c.Children = append(c.Children[:i], c.Children[i+1:]...)
			c.Children = append(c.Children, child)
		}
	}
	c.GetGraphic().MoveFront(child.GetGraphic())
}
