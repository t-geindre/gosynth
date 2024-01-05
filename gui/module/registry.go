package module

type registry struct {
	modules map[string]func() *Module
}

var Registry = &registry{
	modules: make(map[string]func() *Module),
}

func (r *registry) Register(name string, constructor func() *Module) {
	r.modules[name] = constructor
}

func (r *registry) Get(name string) *Module {
	return r.modules[name]()
}

func (r *registry) GetNames() []string {
	names := make([]string, 0)
	for name := range r.modules {
		names = append(names, name)
	}
	return names
}

func init() {
	Registry.Register("Output", func() *Module {
		return NewOutput()
	})

	Registry.Register("VCO", func() *Module {
		return NewVCO()
	})

	Registry.Register("VCA", func() *Module {
		return NewVCA()
	})

	Registry.Register("Delay", func() *Module {
		return NewDelay()
	})
}
