package component

type Filler struct {
	*Component
}

func NewFiller(fill float64) *Filler {
	f := &Filler{
		Component: NewComponent(),
	}

	f.GetLayout().SetFill(fill)

	return f
}
