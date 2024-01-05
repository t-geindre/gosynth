package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/module"
	"gosynth/output"
)

const WindowWidth = 1600
const WindowHeight = 1200

type App struct {
	Streamer *output.Streamer
	Root     component.IComponent
}

func NewApp(str *output.Streamer) *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(WindowWidth, WindowHeight)

	a := &App{}
	a.Streamer = str

	a.Root = component.NewRoot()

	rack := connection.NewRack()
	menu := module.NewMenu(rack)

	a.Root.Append(menu)
	a.Root.Append(rack)

	a.Root.Append(component.NewFPS())

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Root.Draw(screen)
}

func (a *App) Update() error {
	a.Root.Update()

	return nil
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	w, h := float64(outsideWidth)*s, float64(outsideHeight)*s
	a.Root.GetLayout().SetSize(w, h)
	return int(w), int(h)
}
