package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Node struct {
	Parent     INode
	Children   []INode
	Options    *ebiten.DrawImageOptions
	Image      *ebiten.Image
	PosX, PosY int
}

func NewNode(parent INode, width, height int) *Node {
	n := &Node{}
	n.Parent = parent
	n.Children = make([]INode, 0)
	n.Options = &ebiten.DrawImageOptions{}
	n.Image = ebiten.NewImage(width, height)

	return n
}

func (n *Node) Resize(width, height int) {
	n.Image = ebiten.NewImage(width, height)
}

func (n *Node) Append(child INode) {
	n.Children = append(n.Children, child)
}

func (n *Node) Remove(child INode) {
	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			return
		}
	}
}

func (n *Node) Clear() {
	n.Children = make([]INode, 0)
}

func (n *Node) GetParent() INode {
	return n.Parent
}

func (n *Node) Draw(dest *ebiten.Image) {
	for _, child := range n.Children {
		child.Draw(n.Image)
	}

	dest.DrawImage(n.Image, n.Options)
}

func (n *Node) Update() error {
	for _, child := range n.Children {
		err := child.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) GetNodeAt(x, y int) INode {
	var node INode = nil
	for _, child := range n.Children {
		node := child.GetNodeAt(x, y)
		if node != nil {
			return node
		}
	}

	return node
}

func (n *Node) SetPosition(x, y int) {
	n.Options.GeoM.Reset()
	n.Options.GeoM.Translate(float64(x), float64(y))

	n.PosX = x
	n.PosY = y
}
