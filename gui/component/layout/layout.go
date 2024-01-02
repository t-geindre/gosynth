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

func (l *Layout) GetAbsolutePosition() *Position {
	x, y := l.GetPosition().Get()
	parent := l.GetParent()
	for parent != nil {
		px, py := parent.GetPosition().Get()
		x += px
		y += py
		parent = parent.GetParent()
	}
	return &Position{x, y, nil}
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
	computeHorizontal(l)
	l.GetDispatcher().Dispatch(event.NewEvent(UpdatedEvent, l))
}
