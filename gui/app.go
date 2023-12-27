package gui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/event"
	"gosynth/gui/node"
	"gosynth/output"
)

type App struct {
	Rack        *node.Rack
	MouseTarget node.INode
	Streamer    *output.Streamer
}

func NewApp(str *output.Streamer) *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 600)

	a := &App{}
	a.Rack = node.NewRack(800, 600)
	a.Streamer = str

	mod := node.NewVCA()
	a.Rack.Append(mod)

	return a
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Rack.Clear()
	a.Rack.Draw(screen)
	// Draw FPS
	w, _ := a.Rack.GetSize()
	y, _ := a.Rack.GetPosition()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%0.2f", ebiten.ActualFPS()), w-40, y+10)
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
			a.MouseTarget.Dispatch(event.NewEvent(node.LeftMouseUpEvent, a.MouseTarget))
		}

		if ebiten.IsKeyPressed(ebiten.KeyAlt) {
			a.MouseTarget = a.Rack
		} else {
			a.MouseTarget = a.Rack.GetTargetNodeAt(ebiten.CursorPosition())
		}

		if a.MouseTarget != nil {
			a.MouseTarget.Dispatch(event.NewEvent(node.LeftMouseDownEvent, a.MouseTarget))
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && a.MouseTarget != nil {
		a.MouseTarget.Dispatch(event.NewEvent(node.LeftMouseUpEvent, a.MouseTarget))
		a.MouseTarget = nil
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mod := node.NewVCA()
		mod.SetPosition(ebiten.CursorPosition())
		a.Rack.Append(mod)
	}

	return a.Rack.Update()
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	w, h := int(float64(outsideWidth)*s), int(float64(outsideHeight)*s)
	a.Rack.Resize(w, h)
	return w, h
}
