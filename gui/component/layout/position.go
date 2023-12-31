package layout

type Position struct {
	x, y     int
	onChange func(x, y int)
}

func (p *Position) Set(x, y int) {
	if x == p.x && y == p.y {
		return
	}

	p.x = x
	p.y = y

	if p.onChange != nil {
		p.onChange(x, y)
	}
}

func (p *Position) SetX(x int) {
	p.Set(x, p.y)
}

func (p *Position) SetY(y int) {
	p.Set(p.x, y)
}

func (p *Position) MoveBy(x, y int) {
	p.Set(p.x+x, p.y+y)
}

func (p *Position) MoveByX(x int) {
	p.Set(p.x+x, p.y)
}

func (p *Position) MoveByY(y int) {
	p.Set(p.x, p.y+y)
}

func (p *Position) Get() (int, int) {
	return p.x, p.y
}

func (p *Position) GetX() int {
	return p.x
}

func (p *Position) GetY() int {
	return p.y
}

func (p *Position) setOnChangeFunc(f func(x, y int)) {
	p.onChange = f
}
