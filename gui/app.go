package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	Root        *Node
	MouseTarget INode
}

func NewApp() *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 600)

	a := &App{}
	a.Root = NewNode(800, 600, a.Root)

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Root.Draw(screen)
}

func (a *App) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if a.MouseTarget != nil {
			a.MouseTarget.MouseLeftUp()
		}
		a.MouseTarget = a.Root.GetNodeAt(ebiten.CursorPosition()).GetINode()
		if a.MouseTarget != nil {
			a.MouseTarget.MouseLeftDown()
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && a.MouseTarget != nil {
		a.MouseTarget.MouseLeftUp()
		a.MouseTarget = nil
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mod := NewModule()
		mod.SetPosition(ebiten.CursorPosition())
		a.Root.Append(mod)
	}

	return a.Root.Update()
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	a.Root.Resize(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}
