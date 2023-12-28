package gui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/event"
	"gosynth/gui/node"
	"gosynth/module"
	"gosynth/output"
	clock "gosynth/time"
	"time"
)

type App struct {
	Rack        *node.Rack
	MouseTarget node.INode
	Streamer    *output.Streamer
	Clock       *clock.Clock
	AudioVca    *module.VCA // TODO TEST PURPOSE REMOVE ME
}

func NewApp(str *output.Streamer) *App {
	ebiten.SetWindowTitle("Gosynth")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 600)

	a := &App{}
	a.Rack = node.NewRack(800, 600)
	a.Streamer = str

	a.Clock = clock.NewClock(time.Second / time.Duration(ebiten.TPS()))

	for _, audioModule := range str.GetRack().GetModules() {
		if aVca, ok := audioModule.(*module.VCA); ok {
			vca := node.NewVCA(aVca)
			a.Rack.Append(vca)
			a.AudioVca = aVca
		}
	}
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
	a.Clock.Tick()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if a.Streamer.IsSilenced() {
			a.Streamer.Silence() <- false
		} else {
			a.Streamer.Silence() <- true
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
		mod := node.NewVCA(a.AudioVca) // TODO TEST PURPOSE REMOVE ME
		mod.SetPosition(ebiten.CursorPosition())
		a.Rack.Append(mod)
	}

	return a.Rack.Update(a.Clock.GetTime())
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	w, h := int(float64(outsideWidth)*s), int(float64(outsideHeight)*s)
	a.Rack.SetSize(w, h)
	return w, h
}
