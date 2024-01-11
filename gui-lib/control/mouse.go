package control

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gosynth/event"
)

type Mouse struct {
	rootNode       ITarget
	targetHandlers map[event.Id]ITarget
}

func NewMouse(rootNode ITarget) *Mouse {
	return &Mouse{
		rootNode: rootNode,
		targetHandlers: map[event.Id]ITarget{
			RightMouseUpEvent: nil,
			LeftMouseUpEvent:  nil,
			MouseEnterEvent:   nil,
		},
	}
}

func (m *Mouse) Update() {
	currentTarget, _ := m.rootNode.GetTargetAt(ebiten.CursorPosition())
	m.mouseEnterLeave(currentTarget)
	m.mouseDownUp(currentTarget, ebiten.MouseButtonLeft, LeftMouseUpEvent, LeftMouseDownEvent)
	m.mouseDownUp(currentTarget, ebiten.MouseButtonRight, RightMouseUpEvent, RightMouseDownEvent)
}

func (m *Mouse) mouseEnterLeave(currentTarget ITarget) {
	if currentTarget != m.targetHandlers[MouseEnterEvent] {
		if m.targetHandlers[MouseEnterEvent] != nil {
			m.targetHandlers[MouseEnterEvent].Dispatch(event.NewEvent(MouseLeaveEvent, m.targetHandlers[MouseEnterEvent]))
		}
		if currentTarget != nil {
			currentTarget.Dispatch(event.NewEvent(MouseEnterEvent, currentTarget))
		}
		m.targetHandlers[MouseEnterEvent] = currentTarget
	}
}

func (m *Mouse) mouseDownUp(currentTarget ITarget, button ebiten.MouseButton, upEvent, downEvent event.Id) {
	if inpututil.IsMouseButtonJustPressed(button) {
		if m.targetHandlers[upEvent] != nil {
			m.targetHandlers[upEvent].Dispatch(event.NewEvent(upEvent, m.targetHandlers[upEvent]))
		}

		m.targetHandlers[upEvent] = currentTarget
		if m.targetHandlers[upEvent] != nil {
			m.targetHandlers[upEvent].Dispatch(event.NewEvent(downEvent, m.targetHandlers[upEvent]))
			if upEvent == LeftMouseUpEvent {
				m.targetHandlers[upEvent].Dispatch(event.NewEvent(FocusEvent, m.targetHandlers[upEvent]))
			}
		}
	}

	if inpututil.IsMouseButtonJustReleased(button) && m.targetHandlers[upEvent] != nil {
		m.targetHandlers[upEvent].Dispatch(event.NewEvent(upEvent, m.targetHandlers[upEvent]))
		m.targetHandlers[upEvent] = nil
	}
}
