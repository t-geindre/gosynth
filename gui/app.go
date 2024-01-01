package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gosynth/gui/component"
	"gosynth/gui/component/control"
	"gosynth/gui/component/layout"
	"gosynth/gui/component/module"
	"gosynth/gui/component/widget"
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

	a.Root = widget.NewContainer()
	a.Root.Append(widget.NewMenu())

	mod := module.NewVCA()

	rack := widget.NewRack()
	rack.Append(mod)

	a.Root.Append(rack)
	a.Root.Append(widget.NewDebug())
	//a.Root = demo.NewDemo()
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
	w, h := int(float64(outsideWidth)*s), int(float64(outsideHeight)*s)
	a.Root.GetLayout().GetSize().Set(w, h)
	return w, h
}
