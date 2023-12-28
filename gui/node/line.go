package node

import (
	"gosynth/gui/theme"
)

type LineOrientation uint8

const (
	LineOrientationHorizontal LineOrientation = iota
	LineOrientationVertical
)

type Line struct {
	*Node
	Orientation LineOrientation
	Length      int
	Width       int
	Inverted    bool
}

func NewLine(length, width int, orientation LineOrientation) *Line {
	l := &Line{}
	w, h := l.ComputeImageSize(length, width, orientation)
	l.Node = NewNode(w, h, l)
	l.Orientation = orientation
	l.Length = length
	l.Width = width
	l.Dirty = true

	return l
}

func (l *Line) Clear() {
	if l.Dirty {
		l.Inverted = IsNodeInverted(l)

		color := theme.Colors.BackgroundInverted
		if l.Inverted {
			color = theme.Colors.Background
		}

		l.Image.Fill(color)
		l.Dirty = false
	}

	l.Node.Clear()
}

func (l *Line) ComputeImageSize(length, width int, orientation LineOrientation) (int, int) {
	if orientation == LineOrientationHorizontal {
		return length, width
	} else {
		return width, length
	}
}

func (l *Line) Targetable() bool {
	return false
}
