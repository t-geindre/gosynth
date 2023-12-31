package layout

type Spacing struct {
	top, bottom, left, right int
}

func (s *Spacing) SetAll(all int) {
	s.top = all
	s.bottom = all
	s.left = all
	s.right = all
}

func (s *Spacing) Set(top, bottom, left, right int) {
	s.top = top
	s.bottom = bottom
	s.left = left
	s.right = right
}

func (s *Spacing) SetTop(top int) {
	s.top = top
}

func (s *Spacing) SetBottom(bottom int) {
	s.bottom = bottom
}

func (s *Spacing) SetLeft(left int) {
	s.left = left
}

func (s *Spacing) SetRight(right int) {
	s.right = right
}

func (s *Spacing) Get() (int, int, int, int) {
	return s.top, s.bottom, s.left, s.right
}

func (s *Spacing) GetTop() int {
	return s.top
}

func (s *Spacing) GetBottom() int {
	return s.bottom
}

func (s *Spacing) GetLeft() int {
	return s.left
}

func (s *Spacing) GetRight() int {
	return s.right
}
