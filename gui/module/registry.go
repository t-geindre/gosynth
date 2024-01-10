package module

import (
	"gosynth/gui-lib/component"
	"gosynth/gui/connection"
	audio "gosynth/module"
)

type registry struct {
	modules map[string]func(r *registry, rack *connection.Rack)
	lastMod component.IComponent
}

var Registry = &registry{
	modules: make(map[string]func(r *registry, rack *connection.Rack)),
}

func (r *registry) Register(name string, constructor func(r *registry, rack *connection.Rack)) {
	r.modules[name] = constructor
}

func (r *registry) Get(name string) func(r *registry, rack *connection.Rack) {
	return r.modules[name]
}

func (r *registry) GetNames() []string {
	names := make([]string, 0)
	for name := range r.modules {
		names = append(names, name)
	}
	return names
}

func (r *registry) Offset(m component.IComponent) {
	if r.lastMod != nil {
		x, y := r.lastMod.GetLayout().GetPosition()
		w, _ := r.lastMod.GetLayout().GetSize()
		m.GetLayout().SetPosition(x+w, y)
	}
	r.lastMod = m
}

func init() {
	Registry.Register("Output", func(r *registry, rack *connection.Rack) {
		guiMod := NewOutput(rack.GetAudioRack())
		rack.Append(guiMod)
		r.Offset(guiMod)
	})

	Registry.Register("VCO", func(r *registry, rack *connection.Rack) {
		audioMod := audio.NewVCO(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioMod)
		guiMod := NewVCO(audioMod)
		rack.Append(guiMod)
		r.Offset(guiMod)
	})

	Registry.Register("VCA", func(r *registry, rack *connection.Rack) {
		audioMod := audio.NewVCA(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioMod)
		guiMod := NewVCA(audioMod)
		rack.Append(guiMod)
		r.Offset(guiMod)
	})

	Registry.Register("delay", func(r *registry, rack *connection.Rack) {
		audioMod := audio.NewDelay(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioMod)
		guiMod := NewDelay(audioMod)
		rack.Append(guiMod)
		r.Offset(guiMod)
	})

	Registry.Register("Sequencer4", func(r *registry, rack *connection.Rack) {
		audioSequencer4 := audio.NewSequencer4(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioSequencer4)
		guiMod := NewSequencer4(audioSequencer4)
		rack.Append(guiMod)
		r.Offset(guiMod)
	})

	Registry.Register("Mixer", func(r *registry, rack *connection.Rack) {
		audioMixer := audio.NewMixer(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioMixer)
		guiMod := NewMixer(audioMixer)
		rack.Append(guiMod)
		r.Offset(guiMod)
	})

	Registry.Register("Multiplier", func(r *registry, rack *connection.Rack) {
		audioMultiplier := audio.NewMultiplier(rack.GetAudioRack().GetSampleRate())
		rack.GetAudioRack().AddModule(audioMultiplier)
		guiMod := NewMultiplier(audioMultiplier)
		rack.Append(guiMod)
		r.Offset(guiMod)
	})
}
