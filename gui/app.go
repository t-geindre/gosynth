package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	Root *Node
}

func NewApp() *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 600)

	a := &App{}
	a.Root = NewNode(nil, 800, 600)

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Root.Draw(screen)
}

func (a *App) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mod := NewModule()
		x, y := ebiten.CursorPosition()
		mod.SetPosition(x, y)
		a.Root.Append(mod)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

	}

	return nil
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	a.Root.Resize(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}
