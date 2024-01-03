package layout

type Size struct {
	w, h     float64
	onChange func(w, h float64)
}

func NewSize() *Size {
	s := &Size{}

	return s
}

func (s *Size) Set(w, h float64) {
	if w == s.w && h == s.h {
		return
	}

	s.w = w
	s.h = h

	if s.onChange != nil {
		s.onChange(w, h)
	}
}

func (s *Size) SetWidth(w float64) {
	s.Set(w, s.h)
}

func (s *Size) SetHeight(h float64) {
	s.Set(s.w, h)
}

func (s *Size) Get() (float64, float64) {
	return s.w, s.h
}

func (s *Size) GetWidth() float64 {
	return s.w
}

func (s *Size) GetHeight() float64 {
	return s.h
}

func (s *Size) Add(w, h float64) {
	s.Set(s.w+w, s.h+h)
}

func (s *Size) AddWidth(w float64) {
	s.Set(s.w+w, s.h)
}

func (s *Size) AddHeight(h float64) {
	s.Set(s.w, s.h+h)
}

func (s *Size) setOnChangeFunc(f func(w, h float64)) {
	s.onChange = f
}
