package layout

import (
	"gosynth/event"
	"math"
)

type Layout struct {
	*event.Dispatcher

	parent   ILayout
	children []ILayout

	x, y           float64
	w, h           float64
	ww, wh         float64
	pt, pb, pl, pr float64
	mt, mb, ml, mr float64
	fill           float64
	absPos         bool
	orientation    Orientation
}

func NewLayout() *Layout {
	l := &Layout{
		children:    make([]ILayout, 0),
		Dispatcher:  event.NewDispatcher(),
		orientation: Vertical,
	}

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

func (l *Layout) GetMargin() (float64, float64, float64, float64) {
	return l.mt, l.mb, l.ml, l.mr
}

func (l *Layout) SetMargin(top, bottom, left, right float64) {
	l.mt, l.mb, l.ml, l.mr = top, bottom, left, right
}

func (l *Layout) GetPadding() (float64, float64, float64, float64) {
	return l.pt, l.pb, l.pl, l.pr
}

func (l *Layout) SetPadding(top, bottom, left, right float64) {
	l.pt, l.pb, l.pl, l.pr = top, bottom, left, right
}

func (l *Layout) GetPosition() (float64, float64) {
	return l.x, l.y
}

func (l *Layout) SetPosition(x, y float64) {
	if l.x != x || l.y != y {
		l.x, l.y = x, y
		l.Dispatch(event.NewEvent(MoveEvent, l))
	}
}

func (l *Layout) GetAbsolutePosition() (float64, float64) {
	x, y := l.GetPosition()
	parent := l.GetParent()
	for parent != nil {
		px, py := parent.GetPosition()
		x += px
		y += py
		parent = parent.GetParent()
	}
	return x, y
}

func (l *Layout) GetSize() (float64, float64) {
	return l.w, l.h
}

func (l *Layout) SetSize(w, h float64) {
	if l.w != w || l.h != h {
		l.w, l.h = w, h
		l.Dispatch(event.NewEvent(ResizeEvent, l))
		l.ScheduleUpdate()
	}
}

func (l *Layout) GetWantedSize() (float64, float64) {
	return l.ww, l.wh
}

func (l *Layout) SetWantedSize(w, h float64) {
	if l.ww != w || l.wh != h {
		if p := l.GetParent(); p != nil {
			p.ScheduleUpdate()
		}
		l.ww, l.wh = w, h
	}
}

func (l *Layout) SetContentOrientation(orientation Orientation) {
	l.orientation = orientation
}

func (l *Layout) GetContentOrientation() Orientation {
	return l.orientation
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
	return l.x <= x && x <= l.x+l.w &&
		l.y <= y && y <= l.y+l.h
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
	l.Dispatch(event.NewEvent(UpdateStartsEvent, l))
	l.doUpdate()
	l.Dispatch(event.NewEvent(UpdatedEvent, l))
}

func (l *Layout) doUpdate() {
	// No children to handle
	if len(l.GetChildren()) == 0 {
		return
	}

	// Filter out absolute children
	children := make([]ILayout, 0, len(l.GetChildren()))

	for _, c := range l.GetChildren() {
		if !c.GetAbsolutePositioning() {
			children = append(children, c)
		}
	}

	// No children to handle left
	if len(children) == 0 {
		return
	}

	// Store orientation for further usage
	orientVertical := l.orientation == Vertical

	// Compute required size
	contentWidth := float64(0)
	for _, c := range children {
		cww, cwh := c.GetWantedSize()
		cmt, cmb, cml, cmr := c.GetMargin()

		if orientVertical {
			contentWidth += cwh + cmt + cmb
		} else {
			contentWidth += cww + cml + cmr
		}
	}

	// Calculate scaling factor and free space
	innerWidth := l.w - l.pl - l.pr
	innerHeight := l.h - l.pt - l.pb

	if orientVertical {
		innerWidth, innerHeight = innerHeight, innerWidth
	}

	scaleFact := innerWidth / contentWidth
	freeSpace := math.Max(0, innerWidth-contentWidth)

	// If there is free space, we need to distribute it across the filler children
	if freeSpace > 0 {
		totalFill := float64(0)
		for _, c := range children {
			if c.GetFill() > 0 {
				totalFill += c.GetFill()
			}
		}

		if totalFill > 100 {
			panic("Total fill cannot be greater than 100%")
		}

		contentWidth += freeSpace * totalFill / 100
		scaleFact = innerWidth / contentWidth
	}

	// Place children and apply scale
	xOffset := l.pl
	yOffset := l.pt
	for _, c := range children {
		cww, cwh := c.GetWantedSize()
		cmt, cmb, cml, cmr := c.GetMargin()

		fill := float64(0)
		if c.GetFill() > 0 && freeSpace > 0 {
			fill = c.GetFill() / 100 * freeSpace
		}

		if orientVertical {
			c.SetPosition(xOffset+cml, yOffset+cmt*scaleFact)
			ch := (cwh + fill) * scaleFact
			c.SetSize(innerHeight-cml-cmr, ch)
			yOffset += ch + (cmt+cmb)*scaleFact
		} else {
			c.SetPosition(xOffset+cml*scaleFact, yOffset+cmt)
			cw := (cww + fill) * scaleFact
			c.SetSize(cw, innerHeight-cmt-cmb)
			xOffset += cw + (cml+cmr)*scaleFact
		}
	}
}
