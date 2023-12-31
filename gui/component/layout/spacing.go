package layout

type Spacing struct {
	Top, Bottom, Left, Right int
}

func (s *Spacing) SetAll(all int) {
	s.Top = all
	s.Bottom = all
	s.Left = all
	s.Right = all
}

func (s *Spacing) Set(top, bottom, left, right int) {
	s.Top = top
	s.Bottom = bottom
	s.Left = left
	s.Right = right
}

func (s *Spacing) SetTop(top int) {
	s.Top = top
}

func (s *Spacing) SetBottom(bottom int) {
	s.Bottom = bottom
}

func (s *Spacing) SetLeft(left int) {
	s.Left = left
}

func (s *Spacing) SetRight(right int) {
	s.Right = right
}
