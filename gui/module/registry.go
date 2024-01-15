package module

import (
	"gosynth/event"
	"gosynth/gui-lib/component"
	"gosynth/gui/widget"
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
	mod := modules[id].Build(r.rack)
	r.container.Append(mod)
	mod.AddListener(&r, widget.MenuOpenEvent, func(e event.IEvent) {
		e.GetSource().(*widget.Menu).AddContextualOption("Remove", func() {
			r.container.Remove(mod)
			r.container.GetGraphic().ScheduleUpdate()
		})
	})
}

type moduleId uint8
type moduleBuilder func(rack *audio.Rack) component.IComponent
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
	Register("Output", func(rack *audio.Rack) component.IComponent {
		return NewOutput(rack)
	})
	Register("VCO", func(rack *audio.Rack) component.IComponent {
		audioMod := audio.NewVCO(rack.GetSampleRate())
		rack.AddModule(audioMod)
		return NewVCO(audioMod)
	})
	Register("VCA", func(rack *audio.Rack) component.IComponent {
		audioMod := audio.NewVCA(rack.GetSampleRate())
		rack.AddModule(audioMod)
		return NewVCA(audioMod)
	})

	Register("delay", func(rack *audio.Rack) component.IComponent {
		audioMod := audio.NewDelay(rack.GetSampleRate())
		rack.AddModule(audioMod)
		return NewDelay(audioMod)
	})

	Register("Sequencer4", func(rack *audio.Rack) component.IComponent {
		audioSequencer4 := audio.NewSequencer4(rack.GetSampleRate())
		rack.AddModule(audioSequencer4)
		return NewSequencer4(audioSequencer4)
	})

	Register("Mixer", func(rack *audio.Rack) component.IComponent {
		audioMixer := audio.NewMixer(rack.GetSampleRate())
		rack.AddModule(audioMixer)
		return NewMixer(audioMixer)
	})

	Register("Multiplier", func(rack *audio.Rack) component.IComponent {
		audioMultiplier := audio.NewMultiplier(rack.GetSampleRate())
		rack.AddModule(audioMultiplier)
		return NewMultiplier(audioMultiplier)
	})

	Register("Quantizer (CMS)", func(rack *audio.Rack) component.IComponent {
		audioQuantizer := audio.NewQuantizer(rack.GetSampleRate())
		rack.AddModule(audioQuantizer)
		return NewQuantizer(audioQuantizer)
	})

	Register("Clock", func(rack *audio.Rack) component.IComponent {
		audioClock := audio.NewClock(rack.GetSampleRate())
		rack.AddModule(audioClock)
		return NewClock(audioClock)
	})

	Register("VCF", func(rack *audio.Rack) component.IComponent {
		audioVCF := audio.NewVCF(rack.GetSampleRate())
		rack.AddModule(audioVCF)
		return NewVCF(audioVCF)
	})
}
