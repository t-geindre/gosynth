package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/gui/component"
	"gosynth/gui/control"
	"gosynth/gui/demo"
	"gosynth/gui/layout"
	"gosynth/gui/module"
	widget2 "gosynth/gui/widget"
	"gosynth/output"
)

const WindowWidth = 1600
const WindowHeight = 1200

type App struct {
	Streamer *output.Streamer
	Root     component.IComponent
	Mouse    *control.Mouse
}

func NewApp(str *output.Streamer) *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(WindowWidth, WindowHeight)

	a := &App{}
	a.Streamer = str

	a.Root = widget2.NewContainer()
	a.Root.Append(widget2.NewMenu())

	mod := module.NewVCA()

	rack := widget2.NewRack()
	rack.Append(mod)

	a.Root.Append(rack)
	a.Root.Append(widget2.NewFPS())

	a.Root = demo.NewDemo()

	a.Mouse = control.NewMouse(a.Root)

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Root.Draw(screen)
}

func (a *App) Update() error {
	a.Mouse.Update()
	a.Root.Update()
	layout.Sync.Update()

	return nil
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	w, h := float64(outsideWidth)*s, float64(outsideHeight)*s
	a.Root.GetLayout().GetSize().Set(w, h)
	return int(w), int(h)
}
