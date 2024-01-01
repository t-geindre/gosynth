package layout

import (
	"gosynth/event"
)

type Layout struct {
	parent             ILayout
	children           []ILayout
	dispatcher         *event.Dispatcher
	position           *Position
	size               *Size
	wantedSize         *Size
	padding            *Spacing
	margin             *Spacing
	fill               float64
	absPos             bool
	contentOrientation Orientation
}

func NewLayout() *Layout {
	l := &Layout{
		children:           make([]ILayout, 0),
		dispatcher:         event.NewDispatcher(),
		position:           &Position{},
		size:               NewSize(),
		wantedSize:         NewSize(),
		padding:            &Spacing{},
		margin:             &Spacing{},
		contentOrientation: Vertical,
		absPos:             false,
		fill:               0,
	}

	l.ScheduleUpdate()

	l.size.setOnChangeFunc(func(w, h float64) {
		l.GetDispatcher().Dispatch(event.NewEvent(ResizeEvent, l))
		l.ScheduleUpdate()
	})

	l.wantedSize.setOnChangeFunc(func(w, h float64) {
		p := l.GetParent()
		if p != nil {
			p.ScheduleUpdate()
		}
	})

	l.position.setOnChangeFunc(func(x, y float64) {
		l.GetDispatcher().Dispatch(event.NewEvent(MoveEvent, l))
	})

	return l
}

func (l *Layout) GetChildren() []ILayout {
	return l.children
}

func (l *Layout) GetParent() ILayout {
	return l.parent
}

func (l *Layout) SetParent(parent ILayout) {
	l.parent = parent
}

func (l *Layout) Append(child ILayout) {
	l.children = append(l.children, child)
	child.SetParent(l)
}

func (l *Layout) Remove(child ILayout) {
	for i, c := range l.children {
		if c == child {
			l.children = append(l.children[:i], l.children[i+1:]...)
		}
	}
	child.SetParent(nil)
}

func (l *Layout) GetDispatcher() event.IDispatcher {
	return l.dispatcher
}

func (l *Layout) GetMargin() *Spacing {
	return l.margin
}

func (l *Layout) GetPadding() *Spacing {
	return l.padding
}

func (l *Layout) GetPosition() *Position {
	return l.position
}

func (l *Layout) GetSize() *Size {
	return l.size
}

func (l *Layout) GetWantedSize() *Size {
	return l.wantedSize
}

func (l *Layout) SetContentOrientation(orientation Orientation) {
	l.contentOrientation = orientation
}

func (l *Layout) GetContentOrientation() Orientation {
	return l.contentOrientation
}

func (l *Layout) SetAbsolutePositioning(absolute bool) {
	l.absPos = absolute
}

func (l *Layout) GetAbsolutePositioning() bool {
	return l.absPos
}

func (l *Layout) SetFill(fill float64) {
	l.fill = fill
}

func (l *Layout) GetFill() float64 {
	return l.fill
}

func (l *Layout) PointCollides(x, y float64) bool {
	return l.position.x <= x && x <= l.position.x+l.size.w &&
		l.position.y <= y && y <= l.position.y+l.size.h
}

func (l *Layout) GetDepth() int {
	deep := 0

	parent := l.GetParent()
	for parent != nil {
		deep++
		parent = parent.GetParent()
	}

	return deep
}

func (l *Layout) ScheduleUpdate() {
	Sync.ScheduleUpdate(l)
}

func (l *Layout) Update() {
	l.GetDispatcher().Dispatch(event.NewEvent(UpdateStartsEvent, l))

	if len(l.children) > 0 {
		children := make([]ILayout, 0)
		for _, c := range l.children {
			if c.GetAbsolutePositioning() {
				continue
			}
			children = append(children, c)
		}

		if len(children) > 0 {
			freeX, freeY := l.placingPass(children)
			l.fillingPass(children, freeX, freeY)
		}
	}

	l.GetDispatcher().Dispatch(event.NewEvent(UpdatedEvent, l))
}

func (l *Layout) placingPass(children []ILayout) (float64, float64) {
	xOffset := l.GetPadding().GetLeft()
	yOffset := l.GetPadding().GetTop()

	width, height := l.GetSize().Get()

	innerWidth := width - l.GetPadding().GetLeft() - l.GetPadding().GetRight()
	innerHeight := height - l.GetPadding().GetTop() - l.GetPadding().GetBottom()

	for _, c := range children {
		c.GetPosition().Set(
			xOffset+c.GetMargin().GetLeft(),
			yOffset+c.GetMargin().GetTop(),
		)

		if l.GetContentOrientation() == Vertical {
			c.GetSize().Set(
				innerWidth-c.GetMargin().GetLeft()-c.GetMargin().GetRight(),
				c.GetWantedSize().GetHeight(),
			)
			yOffset += c.GetSize().h + c.GetMargin().GetTop() + c.GetMargin().GetBottom()
		} else {
			c.GetSize().Set(
				c.GetWantedSize().GetWidth(),
				innerHeight-c.GetMargin().GetTop()-c.GetMargin().GetBottom(),
			)
			xOffset += c.GetSize().w + c.GetMargin().GetLeft() + c.GetMargin().GetRight()
		}
	}

	return innerWidth - xOffset + l.GetPadding().GetLeft(), innerHeight - yOffset + l.GetPadding().GetTop()
}

func (l *Layout) fillingPass(children []ILayout, freeX, freeY float64) {
	deltaCount := float64(len(children))
	allFreeX, allFreeY := freeX, freeY

	for _, fill := range [...]bool{true, false} {
		shiftX, shiftY := float64(0), float64(0)

		if deltaCount == 0 {
			// Abort second pass, all component already handled
			return
		}

		for _, c := range children {
			deltaX, deltaY := float64(0), float64(0)

			if l.contentOrientation == Vertical {
				if c.GetFill() == 0 && !fill {
					deltaY = freeY / deltaCount
					deltaCount--
				}
				if c.GetFill() > 0 && fill {
					if allFreeY > 0 {
						deltaY = allFreeY * c.GetFill() / 100
					}
					deltaCount--
				}

			} else {
				if c.GetFill() == 0 && !fill {
					deltaX = freeX / deltaCount
					deltaCount--
				}
				if c.GetFill() > 0 && fill {
					if allFreeX > 0 {
						deltaX = allFreeX * c.GetFill() / 100
					}
					deltaCount--
				}
			}

			c.GetPosition().MoveBy(shiftX, shiftY)
			c.GetSize().Add(deltaX, deltaY)

			freeX += -deltaX
			freeY += -deltaY

			shiftX += deltaX
			shiftY += deltaY
		}
	}
}
