package node

import (
	"gosynth/gui/theme"
)

type Container struct {
	*Node
	Inverted bool
}

func NewContainer(width, height int) *Container {
	i := &Container{}
	i.Node = NewNode(width, height, i)
	i.Dirty = true

	return i
}

func (i *Container) Clear() {
	if i.Dirty {
		if i.Inverted {
			i.Image.Fill(theme.Colors.BackgroundInverted)
		} else {
			i.Image.Fill(theme.Colors.Background)
		}
		i.Dirty = false
	}

	i.Node.Clear()
}

func (i *Container) IsInverted() bool {
	return i.Inverted
}

func (i *Container) SetInverted(inverted bool) {
	i.Inverted = inverted
	i.Dirty = true
}

func (i *Container) Targetable() bool {
	return false
}

func IsNodeInverted(n INode) bool {
	for n != nil {
		if c, ok := n.(*Container); ok {
			return c.IsInverted()
		}
		n = n.GetParent()
	}

	return false
}
