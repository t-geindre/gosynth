package layout

type Position struct {
	x, y     float64
	onChange func(x, y float64)
}

func (p *Position) Set(x, y float64) {
	if x == p.x && y == p.y {
		return
	}

	p.x = x
	p.y = y

	if p.onChange != nil {
		p.onChange(x, y)
	}
}

func (p *Position) SetX(x float64) {
	p.Set(x, p.y)
}

func (p *Position) SetY(y float64) {
	p.Set(p.x, y)
}

func (p *Position) MoveBy(x, y float64) {
	p.Set(p.x+x, p.y+y)
}

func (p *Position) MoveByX(x float64) {
	p.Set(p.x+x, p.y)
}

func (p *Position) MoveByY(y float64) {
	p.Set(p.x, p.y+y)
}

func (p *Position) Get() (float64, float64) {
	return p.x, p.y
}

func (p *Position) GetX() float64 {
	return p.x
}

func (p *Position) GetY() float64 {
	return p.y
}

func (p *Position) setOnChangeFunc(f func(x, y float64)) {
	p.onChange = f
}
