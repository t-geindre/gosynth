package layout

import "gosynth/event"

type Orientation uint8

const (
	Horizontal Orientation = iota
	Vertical
)

type ILayout interface {
	GetChildren() []ILayout
	GetParent() ILayout
	SetParent(parent ILayout)
	Append(child ILayout)
	Remove(child ILayout)

	GetDispatcher() event.IDispatcher

	GetMargin() *Spacing
	GetPadding() *Spacing

	// GetPosition
	// Component position will be overridden by the layouting system if the component is not absolute
	GetPosition() *Position

	// GetSize
	// Component size will be overridden by the layouting system if the component is not absolute
	GetSize() *Size

	// GetWantedSize
	// the layouting system will try to set the component size to the wanted size
	// but if there is no filler to fill the remaining space, the wanted size will be overridden
	GetWantedSize() *Size

	// SetContentOrientation
	// the content orientation is the orientation of the children
	// the layouting system will make the components fill all the space of the opposite orientation
	// and make them fill the space of the content orientation equally if there is no filler
	// to fill the remaining free space
	SetContentOrientation(orientation Orientation)
	GetContentOrientation() Orientation

	// SetAbsolutePositioning
	// Component ignored by the layouting system
	// Wanted size has no effect, use size instead
	SetAbsolutePositioning(absolute bool)
	GetAbsolutePositioning() bool

	// SetFill
	// The component will fill its parent according to its contentOrientation
	// Filling is defined in percentage of the remaining free space
	SetFill(fill float64)
	GetFill() float64

	PointCollides(x, y float64) bool

	GetDepth() int
	ScheduleUpdate()
	Update()
}
