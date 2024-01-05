package module

import (
	"gosynth/gui/connection"
	audio "gosynth/module"
)

type registry struct {
	modules map[string]func(rack *connection.Rack)
}

var Registry = &registry{
	modules: make(map[string]func(rack *connection.Rack)),
}

func (r *registry) Register(name string, constructor func(rack *connection.Rack)) {
	r.modules[name] = constructor
}

func (r *registry) Get(name string) func(rack *connection.Rack) {
	return r.modules[name]
}

func (r *registry) GetNames() []string {
	names := make([]string, 0)
	for name := range r.modules {
		names = append(names, name)
	}
	return names
}

func init() {
	Registry.Register("Output", func(rack *connection.Rack) {
		rack.Append(NewOutput(rack.GetAudioRack()))
	})

	Registry.Register("VCO", func(rack *connection.Rack) {
		audioVCO := &audio.Oscillator{}
		rack.GetAudioRack().AddModule(audioVCO)
		rack.Append(NewVCO(audioVCO))
		audioVCO.SetFreq(440)
		audioVCO.SetShape(audio.OscillatorShapeSine)
	})

	Registry.Register("VCA", func(rack *connection.Rack) {
		audioVca := &audio.VCA{}
		rack.Append(NewVCA(audioVca))
		rack.GetAudioRack().AddModule(audioVca)
	})

	Registry.Register("Delay", func(rack *connection.Rack) {
		rack.Append(NewDelay())
	})
}
