package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/event"
)

type App struct {
	Root  *Node
	Mouse *Mouse
}

func NewApp() *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 600)

	a := &App{}
	a.Root = NewNode(800, 600, a.Root)
	a.Mouse = NewMouse()

	// Move module front
	a.Mouse.AddListener(&a, a.Mouse.Events.Click, func(e event.ListenerArgs) {
		m := e.(MouseEvent)
		targetNode := a.Root.GetNodeAt(m.x, m.y)
		if targetNode != nil {
			a.MoveContainingModuleFront(targetNode)
		}
	})

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Root.Draw(screen)
}

func (a *App) Update() error {
	a.Mouse.Update()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mod := NewModule()
		mod.SetPosition(ebiten.CursorPosition())
		a.Root.Append(mod)
	}

	return nil
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	a.Root.Resize(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

func (a *App) MoveContainingModuleFront(node INode) {
	for {
		switch node.GetINode().(type) {
		case *Module:
			if node.GetParent() != nil {
				node.GetParent().MoveFront(node)
				return
			}
		}
		node = node.GetParent()
		if node == nil {
			return
		}
	}
}
