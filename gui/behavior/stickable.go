package behavior

import "gosynth/gui/component"

type StickDirection uint8

const (
	StickDirectionNone  StickDirection = 1 << iota
	StickDirectionLeft                 = 1 << iota
	StickDirectionRight                = 1 << iota
	StickDirectionUp                   = 1 << iota
	StickDirectionDown                 = 1 << iota
)

type Stickable struct {
}

func NewStickable(node component.Component, directions StickDirection) *Stickable {
	s := &Stickable{}

	return s
}
