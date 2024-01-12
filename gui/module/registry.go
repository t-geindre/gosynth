package module

import (
	"gosynth/gui-lib/component"
	audio "gosynth/module"
	"slices"
	"strings"
)

type Registry struct {
	rack      *audio.Rack
	container component.IComponent
}

func NewRegistry(rack *audio.Rack, container component.IComponent) *Registry {
	return &Registry{
		rack:      rack,
		container: container,
	}
}

func (r *Registry) GetModules() []*ModuleEntry {
	mods := make([]*ModuleEntry, 0, len(modules))
	for _, m := range modules {
		mods = append(mods, m)
	}

	// Sort by Name
	slices.SortFunc[[]*ModuleEntry, *ModuleEntry](mods, func(a, b *ModuleEntry) int {
		return strings.Compare(a.Name, b.Name)
	})

	return mods
}

func (r *Registry) Build(id moduleId) {
	modules[id].Build(r.rack, r.container)
}

type moduleId uint8
type moduleBuilder func(rack *audio.Rack, container component.IComponent)
type ModuleEntry struct {
	Build moduleBuilder
	Id    moduleId
	Name  string
}

var lastId moduleId
var modules = make(map[moduleId]*ModuleEntry)

func Register(name string, build moduleBuilder) {
	modules[lastId] = &ModuleEntry{
		Build: build,
		Id:    lastId,
		Name:  name,
	}
	lastId++
}

func init() {
	Register("Output", func(rack *audio.Rack, container component.IComponent) {
		guiMod := NewOutput(rack)
		container.Append(guiMod)
	})
	Register("VCO", func(rack *audio.Rack, container component.IComponent) {
		audioMod := audio.NewVCO(rack.GetSampleRate())
		rack.AddModule(audioMod)
		guiMod := NewVCO(audioMod)
		container.Append(guiMod)
	})
	Register("VCA", func(rack *audio.Rack, container component.IComponent) {
		audioMod := audio.NewVCA(rack.GetSampleRate())
		rack.AddModule(audioMod)
		guiMod := NewVCA(audioMod)
		container.Append(guiMod)
	})

	Register("delay", func(rack *audio.Rack, container component.IComponent) {
		audioMod := audio.NewDelay(rack.GetSampleRate())
		rack.AddModule(audioMod)
		guiMod := NewDelay(audioMod)
		container.Append(guiMod)
	})

	Register("Sequencer4", func(rack *audio.Rack, container component.IComponent) {
		audioSequencer4 := audio.NewSequencer4(rack.GetSampleRate())
		rack.AddModule(audioSequencer4)
		guiMod := NewSequencer4(audioSequencer4)
		container.Append(guiMod)
	})

	Register("Mixer", func(rack *audio.Rack, container component.IComponent) {
		audioMixer := audio.NewMixer(rack.GetSampleRate())
		rack.AddModule(audioMixer)
		guiMod := NewMixer(audioMixer)
		container.Append(guiMod)
	})

	Register("Multiplier", func(rack *audio.Rack, container component.IComponent) {
		audioMultiplier := audio.NewMultiplier(rack.GetSampleRate())
		rack.AddModule(audioMultiplier)
		guiMod := NewMultiplier(audioMultiplier)
		container.Append(guiMod)
	})

	Register("Quantizer (CMS)", func(rack *audio.Rack, container component.IComponent) {
		audioQuantizer := audio.NewQuantizer(rack.GetSampleRate())
		rack.AddModule(audioQuantizer)
		guiMod := NewQuantizer(audioQuantizer)
		container.Append(guiMod)
	})
}
