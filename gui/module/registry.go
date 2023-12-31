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
		audioVCO := audio.NewVCO(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioVCO)
		rack.Append(NewVCO(audioVCO))
	})

	Registry.Register("LFO", func(rack *connection.Rack) {
		audioLFO := audio.NewLFO(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioLFO)
		rack.Append(NewLFO(audioLFO))
	})

	Registry.Register("VCA", func(rack *connection.Rack) {
		audioVca := audio.NewVCA(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioVca)
		rack.Append(NewVCA(audioVca))
	})

	Registry.Register("delay", func(rack *connection.Rack) {
		audioDelay := audio.NewDelay(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioDelay)
		rack.Append(NewDelay(audioDelay))
	})
}
