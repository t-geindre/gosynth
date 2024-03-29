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

	a := &App{}
	a.windowSetup()
	a.uiSetup(audioRack)

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

func (a *App) windowSetup() {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
}

func (a *App) uiSetup(audioRack *audio.Rack) {
	a.Root = component.NewRoot()

	grid := widget.NewGrid(module.ModuleUWidth, module.ModuleHeight)
	header := module.NewHeader()
	rack := connection.NewRack(audioRack)
	rack.Append(grid)

	a.Root.Append(header)
	a.Root.Append(rack)
	a.Root.Append(component.NewFPS())

	menu := widget.NewMenu()
	registry := module.NewRegistry(audioRack, grid)

	for _, mod := range registry.GetModules() {
		menu.AddOption(mod.Name, func(mod *module.ModuleEntry) func() {
			return func() {
				registry.Build(mod.Id)
			}
		}(mod))
	}
	rack.Append(menu)
}
