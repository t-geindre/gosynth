package node

import "gosynth/gui/theme"

type Inverted struct {
	Node
	Inverted bool
	Dirty    bool
}

func NewInverted(width, height int) *Inverted {
	i := &Inverted{}
	i.Node = *NewNode(width, height, i)
	i.Inverted = true
	i.Dirty = true

	return i
}

func (i *Inverted) Clear() {
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

func (i *Inverted) SetParent(parent INode) {
	i.Node.SetParent(parent)
	i.Inverted = IsNodeInverted(i)
}

func (i *Inverted) Targetable() bool {
	return false
}

func IsNodeInverted(n INode) bool {
	inverted := false
	for n != nil {
		if _, ok := n.(*Inverted); ok {
			inverted = !inverted
		}
		n = n.GetParent()
	}

	return inverted
}
