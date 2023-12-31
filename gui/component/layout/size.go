package layout

type Size struct {
	w, h     int
	onChange func(w, h int)
}

func NewSize() *Size {
	s := &Size{}

	return s
}

func (s *Size) Set(w, h int) {
	if w == s.w && h == s.h {
		return
	}

	s.w = w
	s.h = h

	if s.onChange != nil {
		s.onChange(w, h)
	}
}

func (s *Size) SetWidth(w int) {
	s.Set(w, s.h)
}

func (s *Size) SetHeight(h int) {
	s.Set(s.w, h)
}

func (s *Size) Get() (int, int) {
	return s.w, s.h
}

func (s *Size) GetWidth() int {
	return s.w
}

func (s *Size) GetHeight() int {
	return s.h
}

func (s *Size) Add(w, h int) {
	s.Set(s.w+w, s.h+h)
}

func (s *Size) AddWidth(w int) {
	s.Set(s.w+w, s.h)
}

func (s *Size) AddHeight(h int) {
	s.Set(s.w, s.h+h)
}

func (s *Size) setOnChangeFunc(f func(w, h int)) {
	s.onChange = f
}
