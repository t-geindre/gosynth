package layout

import (
	"gosynth/event"
	"math"
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
	fill               int
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

	l.size.OnChange(func(w, h int) {
		l.GetDispatcher().Dispatch(event.NewEvent(ResizeEvent, l))
		l.ScheduleUpdate()
	})

	l.wantedSize.OnChange(func(w, h int) {
		p := l.GetParent()
		if p != nil {
			p.ScheduleUpdate()
		}
	})

	l.position.OnChange(func(x, y int) {
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

func (l *Layout) SetFill(fill int) {
	l.fill = fill
}

func (l *Layout) GetFill() int {
	return l.fill
}

func (l *Layout) PointCollides(x, y int) bool {
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
	if len(l.children) == 0 {
		return
	}

	// Filter out the absolute positioned children
	children := make([]ILayout, 0)
	for _, c := range l.children {
		if c.GetAbsolutePositioning() {
			continue
		}
		children = append(children, c)
	}

	// All children are absolute positioned, abort
	if len(children) == 0 {
		return
	}

	freeX, freeY := l.placingPass(children)
	l.fillingPass(children, freeX, freeY)

	l.GetDispatcher().Dispatch(event.NewEvent(UpdatedEvent, l))
}

func (l *Layout) placingPass(children []ILayout) (int, int) {
	xOffset := l.GetPadding().Left
	yOffset := l.GetPadding().Top

	width, height := l.GetSize().Get()

	innerWidth := width - l.GetPadding().Left - l.GetPadding().Right
	innerHeight := height - l.GetPadding().Top - l.GetPadding().Bottom

	// Range backwards to get the last child on top
	for _, c := range children {
		// Offset the child by its margins
		c.GetPosition().Set(
			xOffset+c.GetMargin().Left,
			yOffset+c.GetMargin().Top,
		)

		// Component always fits all the opposite orientation space
		if l.GetContentOrientation() == Vertical {
			c.GetSize().Set(
				innerWidth-c.GetMargin().Left-c.GetMargin().Right,
				c.GetWantedSize().GetHeight(),
			)
			yOffset += c.GetSize().h + c.GetMargin().Top + c.GetMargin().Bottom
		} else {
			c.GetSize().Set(
				c.GetWantedSize().GetWidth(),
				innerHeight-c.GetMargin().Top-c.GetMargin().Bottom,
			)
			xOffset += c.GetSize().w + c.GetMargin().Left + c.GetMargin().Right
		}
	}

	return innerWidth - xOffset + l.GetPadding().Left, innerHeight - yOffset + l.GetPadding().Top
}

func (l *Layout) fillingPass(children []ILayout, freeX, freeY int) {
	deltaCount := len(children)

	allFreeX, allFreeY := freeX, freeY

	for _, fill := range [...]bool{true, false} {
		shiftX, shiftY := 0, 0

		if deltaCount == 0 {
			return
		}

		for _, c := range children {
			deltaX, deltaY := 0, 0

			if l.contentOrientation == Vertical {
				if c.GetFill() == 0 && !fill {
					deltaY = freeY / deltaCount
					deltaCount--
				}
				if c.GetFill() > 0 && fill {
					if allFreeY > 0 {
						deltaY = int(math.Round(float64(allFreeY) * float64(c.GetFill()) / float64(100)))
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
						deltaX = int(math.Round(float64(allFreeX) * float64(c.GetFill()) / float64(100)))
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
