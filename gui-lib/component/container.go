package component

type Container struct {
	*Component
	inverted bool
}

func NewContainer() *Container {
	c := &Container{
		Component: NewComponent(),
	}

	return c
}
