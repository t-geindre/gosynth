package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Node struct {
	Parent        INode
	Children      []INode
	Options       *ebiten.DrawImageOptions
	Image         *ebiten.Image
	PosX, PosY    int
	Width, Height int
	INode         INode
}

func NewNode(width, height int, inode INode) *Node {
	// Todo fix this inode shit
	n := &Node{}
	n.Children = make([]INode, 0)
	n.Options = &ebiten.DrawImageOptions{}
	n.INode = inode
	n.Resize(width, height)

	return n
}

func (n *Node) Resize(width, height int) {
	if n.Width == width && n.Height == height {
		return
	}

	n.Width = width
	n.Height = height

	n.Image = ebiten.NewImage(width, height)
}

func (n *Node) Append(child INode) {
	child.SetParent(n)
	n.Children = append(n.Children, child)
}

func (n *Node) Remove(child INode) {
	for i, c := range n.Children {
		if c.GetINode() == child.GetINode() {
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
	n.Image.Clear() // Todo remove image from node, create a sprite type to handle image
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

	if x >= n.PosX && x <= n.PosX+n.Image.Bounds().Dx() &&
		y >= n.PosY && y <= n.PosY+n.Image.Bounds().Dy() {

		x -= n.PosX
		y -= n.PosY

		node = n

		// range backward on children, so the top most child will be returned
		for i := len(n.Children) - 1; i >= 0; i-- {
			node := n.Children[i].GetNodeAt(x, y)
			if node != nil {
				return node
			}
		}
	}

	if node == nil {
		return nil
	}

	return node.GetINode()
}

func (n *Node) SetPosition(x, y int) {
	n.Options.GeoM.Reset()
	n.Options.GeoM.Translate(float64(x), float64(y))

	n.PosX = x
	n.PosY = y
}

func (n *Node) Dispose() {
	n.Image.Dispose()
}

func (n *Node) SetParent(parent INode) {
	n.Parent = parent
}

func (n *Node) GetINode() INode {
	return n.INode
}

func (n *Node) MoveFront(child INode) {
	for i, c := range n.Children {
		if c.GetINode() == child.GetINode() {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			n.Children = append(n.Children, child.GetINode())
			return
		}
	}
}

func (n *Node) MoveBy(x, y int) {
	n.SetPosition(n.PosX+x, n.PosY+y)
}

func (n *Node) MoveChildrenBy(x, y int) {
	for _, child := range n.Children {
		child.MoveBy(x, y)
	}
}

func (n *Node) MouseLeftDown(target INode) {
	if parent := n.GetParent(); parent != nil {
		parent.GetINode().MouseLeftDown(target)
	}
}

func (n *Node) MouseLeftUp(target INode) {
	if parent := n.GetParent(); parent != nil {
		parent.GetINode().MouseLeftUp(target)
	}
}
