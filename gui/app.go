package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	"gosynth/gui/module"
	"gosynth/gui/widget"
	audio "gosynth/module"
)

const WindowWidth = 1600
const WindowHeight = 1200

type App struct {
	Root component.IComponent
}

func NewApp(audioRack *audio.Rack) *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(WindowWidth, WindowHeight)

	a := &App{}

	a.Root = component.NewRoot()

	grid := widget.NewGrid(module.ModuleUWidth, module.ModuleHeight)

	menu := module.NewMenu(module.NewRegistry(audioRack, grid))

	rack := connection.NewRack(audioRack)
	rack.Append(grid)

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
