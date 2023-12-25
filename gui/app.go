package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/output"
)

type App struct {
	Rack        *Rack
	MouseTarget INode
	Streamer    *output.Streamer
}

func NewApp(str *output.Streamer) *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 600)

	a := &App{}
	a.Rack = NewRack(800, 600)
	a.Streamer = str

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Rack.Draw(screen)
}

func (a *App) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if a.Streamer.IsPlaying() {
			a.Streamer.Play() <- false
		} else {
			a.Streamer.Play() <- true
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if a.MouseTarget != nil {
			a.MouseTarget.MouseLeftUp()
		}

		if ebiten.IsKeyPressed(ebiten.KeyAlt) {
			a.MouseTarget = a.Rack
		} else {
			a.MouseTarget = a.Rack.GetNodeAt(ebiten.CursorPosition())
		}

		if a.MouseTarget != nil {
			a.MouseTarget.MouseLeftDown()
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && a.MouseTarget != nil {
		a.MouseTarget.MouseLeftUp()
		a.MouseTarget = nil
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mod := NewModule(100, 100)
		mod.SetPosition(ebiten.CursorPosition())
		a.Rack.Append(mod)
	}

	return a.Rack.Update()
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	a.Rack.Resize(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}
