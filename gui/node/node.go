package node

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/event"
)

type Node struct {
	Parent          INode
	Children        []INode
	Options         *ebiten.DrawImageOptions
	AppendOptions   *AppendOptions
	Image           *ebiten.Image
	PosX, PosY      int
	Width, Height   int
	INode           INode
	LayoutDirty     bool
	LayoutComputing bool
	Dirty           bool
	event.Dispatcher
}

func NewNode(width, height int, inode INode) *Node {
	n := &Node{}
	n.Children = make([]INode, 0)
	n.Options = &ebiten.DrawImageOptions{}
	n.INode = inode
	n.LayoutComputing = true
	n.SetSize(width, height)
	n.Dispatcher.Init()

	return n
}

func (n *Node) DisableLayoutComputing() {
	n.LayoutComputing = false
}

func (n *Node) SetSize(width, height int) {
	if n.Width == width && n.Height == height {
		return
	}

	n.Width = width
	n.Height = height

	n.Image = ebiten.NewImage(width, height)

	n.Dirty = true
	n.LayoutDirty = n.LayoutComputing
}

func (n *Node) Append(child INode) {
	n.AppendWithOptions(child, nil)
}

func (n *Node) AppendWithOptions(child INode, options *AppendOptions) {
	child.SetAppendOptions(options)
	child.SetParent(n)

	n.Children = append(n.Children, child)
}

func (n *Node) GetChildren() []INode {
	return n.Children
}

func (n *Node) Remove(child INode) {
	for i, c := range n.Children {
		if c.GetINode() == child.GetINode() {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			child.SetParent(nil)
			return
		}
	}
}

func (n *Node) RemoveAll() {
	n.Children = make([]INode, 0)
}

func (n *Node) GetParent() INode {
	if n.Parent != nil {
		return n.Parent.GetINode()
	}

	return nil
}

func (n *Node) Clear() {
	for _, child := range n.Children {
		child.Clear()
	}

	if n.LayoutDirty {
		ComputeLayout(n)
		n.LayoutDirty = false
	}
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

func (n *Node) GetTargetNodeAt(x, y int) INode {
	var node INode = nil

	if x >= n.PosX && x <= n.PosX+n.Image.Bounds().Dx() &&
		y >= n.PosY && y <= n.PosY+n.Image.Bounds().Dy() {

		x -= n.PosX
		y -= n.PosY

		if n.GetINode().Targetable() {
			node = n
		}

		// range backward on children, so the top most child will be returned
		for i := len(n.Children) - 1; i >= 0; i-- {
			node := n.Children[i].GetTargetNodeAt(x, y)
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

	n.Dirty = true
}

func (n *Node) SetPositionX(x int) {
	n.SetPosition(x, n.PosY)
}

func (n *Node) SetPositionY(y int) {
	n.SetPosition(n.PosX, y)
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

func (n *Node) GetSize() (int, int) {
	return n.Width, n.Height
}

func (n *Node) GetWidth() int {
	return n.Width
}

func (n *Node) GetHeight() int {
	return n.Height
}

func (n *Node) SetWidth(width int) {
	n.SetSize(width, n.Height)
}

func (n *Node) SetHeight(height int) {
	n.SetSize(n.Width, height)
}

func (n *Node) GetPosition() (int, int) {
	return n.PosX, n.PosY
}

func (n *Node) GetPositionX() int {
	return n.PosX
}

func (n *Node) GetPositionY() int {
	return n.PosY
}

func (n *Node) GetAbsolutePosition() (int, int) {
	x, y := n.PosX, n.PosY
	if parent := n.GetParent(); parent != nil {
		px, py := parent.GetINode().GetAbsolutePosition()
		x += px
		y += py
	}
	return x, y
}

func (n *Node) Targetable() bool {
	return true
}

func (n *Node) HCenter() {
	if parent := n.GetParent(); parent != nil {
		pw, _ := parent.GetINode().GetSize()
		w, _ := n.GetSize()
		n.SetPosition(pw/2-w/2, n.PosY)
	}
}

func (n *Node) Dispatch(e event.IEvent) {
	n.Dispatcher.Dispatch(e)

	if e.IsPropagationStopped() {
		return
	}

	if parent := n.GetParent(); parent != nil {
		parent.Dispatch(e)
	}
}

func (n *Node) SetAppendOptions(options *AppendOptions) {
	n.AppendOptions = options
}

func (n *Node) GetAppendOptions() *AppendOptions {
	return n.AppendOptions
}

func (n *Node) GetOuterSize() (int, int) {
	w, h := n.GetSize()
	if options := n.GetAppendOptions(); options != nil {
		w += options.MarginLeft + options.MarginRight
		h += options.MarginTop + options.MarginBottom
	}
	return w, h
}

func (n *Node) GetOuterWidth() int {
	w, _ := n.GetOuterSize()
	return w
}

func (n *Node) GetOuterHeight() int {
	_, h := n.GetOuterSize()
	return h
}
