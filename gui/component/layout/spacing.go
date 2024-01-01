package layout

type Spacing struct {
	top, bottom, left, right float64
}

func (s *Spacing) SetAll(all float64) {
	s.top = all
	s.bottom = all
	s.left = all
	s.right = all
}

func (s *Spacing) Set(top, bottom, left, right float64) {
	s.top = top
	s.bottom = bottom
	s.left = left
	s.right = right
}

func (s *Spacing) SetTop(top float64) {
	s.top = top
}

func (s *Spacing) SetBottom(bottom float64) {
	s.bottom = bottom
}

func (s *Spacing) SetLeft(left float64) {
	s.left = left
}

func (s *Spacing) SetRight(right float64) {
	s.right = right
}

func (s *Spacing) Get() (float64, float64, float64, float64) {
	return s.top, s.bottom, s.left, s.right
}

func (s *Spacing) GetTop() float64 {
	return s.top
}

func (s *Spacing) GetBottom() float64 {
	return s.bottom
}

func (s *Spacing) GetLeft() float64 {
	return s.left
}

func (s *Spacing) GetRight() float64 {
	return s.right
}
