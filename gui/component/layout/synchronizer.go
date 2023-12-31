package layout

import (
	"slices"
)

var Sync *Synchronizer

// Synchronizer Ensure that all layouts are updated once in the correct order
type Synchronizer struct {
	layouts []ILayout
}

func NewSynchronizer() *Synchronizer {
	return &Synchronizer{
		layouts: make([]ILayout, 0),
	}
}

func (s *Synchronizer) ScheduleUpdate(layout ILayout) {
	if !s.Has(layout) {
		s.layouts = append(s.layouts, layout)
	}
}

func (s *Synchronizer) Update() {
	if len(s.layouts) == 0 {
		return
	}

	// Copy the current updates list as we might trigger new updates during update
	updates := make([]ILayout, len(s.layouts))
	copy(updates, s.layouts)

	s.layouts = make([]ILayout, 0)

	// Lowest depth (higher layouts) first
	slices.SortFunc[[]ILayout, ILayout](updates, func(a, b ILayout) int {
		if a.GetDepth() > b.GetDepth() {
			return 1
		}
		if a.GetDepth() < b.GetDepth() {
			return -1
		}
		return 0
	})

	for _, l := range updates {
		s.Remove(l)
		l.Update()
	}

	// We might have triggered new updates during update, handle them
	s.Update()
}

func (s *Synchronizer) Has(layout ILayout) bool {
	for _, l := range s.layouts {
		if l == layout {
			return true
		}
	}
	return false
}

func (s *Synchronizer) Remove(layout ILayout) {
	for i, l := range s.layouts {
		if l == layout {
			s.layouts = append(s.layouts[:i], s.layouts[i+1:]...)
			return
		}
	}
}

func init() {
	Sync = NewSynchronizer()
}
