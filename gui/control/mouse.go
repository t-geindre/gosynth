package control

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/event"
)

type Mouse struct {
	target         ITarget
	leftDownTarget ITarget
	enterTarget    ITarget
}

func NewMouse(target ITarget) *Mouse {
	return &Mouse{
		target: target,
	}
}

func (m *Mouse) Update() {
	mouseTarget, _ := m.target.GetTargetAt(ebiten.CursorPosition())

	if mouseTarget != m.enterTarget {
		if m.enterTarget != nil {
			m.enterTarget.Dispatch(event.NewEvent(MouseLeaveEvent, m.enterTarget))
		}
		if mouseTarget != nil {
			mouseTarget.Dispatch(event.NewEvent(MouseEnterEvent, mouseTarget))
		}
		m.enterTarget = mouseTarget
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if m.leftDownTarget != nil {
			m.leftDownTarget.Dispatch(event.NewEvent(LeftMouseUpEvent, m.leftDownTarget))
		}

		m.leftDownTarget = mouseTarget
		if m.leftDownTarget != nil {
			m.leftDownTarget.Dispatch(event.NewEvent(LeftMouseDownEvent, m.leftDownTarget))
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && m.leftDownTarget != nil {
		m.leftDownTarget.Dispatch(event.NewEvent(LeftMouseUpEvent, m.leftDownTarget))
		m.leftDownTarget = nil
	}
}
