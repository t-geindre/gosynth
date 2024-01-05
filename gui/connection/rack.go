package connection

import (
	"gosynth/event"
	"gosynth/gui-lib/behavior"
	"gosynth/gui-lib/component"
	"gosynth/gui-lib/graphic"
	audio "gosynth/module"
	"image/color"
)

type Rack struct {
	*component.Component
	cables       []*Cable
	currentCable *Cable
	audioRack    *audio.Rack
}

func NewRack(audioRack *audio.Rack) *Rack {
	r := &Rack{
		Component: component.NewComponent(),
		audioRack: audioRack,
	}

	behavior.NewDraggable(r)

	r.AddListener(&r, behavior.DragEvent, r.onDrag)
	r.AddListener(&r, ConnectionStartEvent, r.onConnectionStart)
	r.AddListener(&r, ConnectionStopEvent, r.onConnectionStop)
	r.AddListener(&r, ConnectionEnterEvent, r.onConnectionEnter)
	r.AddListener(&r, ConnectionLeaveEvent, r.onConnectionLeave)

	r.GetGraphic().AddListener(&r, graphic.DrawStartEvent, r.onDrawStart)
	r.GetGraphic().AddListener(&r, graphic.DrawEndEvent, r.onDrawEnd)

	r.GetLayout().SetFill(100)

	return r
}

func (r *Rack) onDrag(e event.IEvent) {
	ev := e.(*behavior.DragEventDetails)
	for _, c := range r.GetChildren() {
		x, y := c.GetLayout().GetPosition()
		c.GetLayout().SetPosition(float64(ev.DeltaX)+x, float64(ev.DeltaY)+y)
	}
	e.StopPropagation()
}

func (r *Rack) onDrawStart(e event.IEvent) {
	r.GetGraphic().GetImage().Fill(color.RGBA{R: 26, G: 26, B: 26, A: 255})
}

func (r *Rack) onDrawEnd(e event.IEvent) {
	x, y := r.GetLayout().GetAbsolutePosition()
	for _, c := range r.cables {
		c.Draw(r.GetGraphic().GetImage(), x, y)
	}
}

func (r *Rack) addCable(c *Cable) {
	r.cables = append(r.cables, c)
}

func (r *Rack) removeCable(c *Cable) {
	for i, cable := range r.cables {
		if cable == c {
			r.cables = append(r.cables[:i], r.cables[i+1:]...)
		}
	}
}

func (r *Rack) onConnectionStart(e event.IEvent) {
	if r.currentCable != nil && r.currentCable.GetDst() == nil {
		r.removeCable(r.currentCable)
	}

	p := e.GetSource().(*Plug)
	c := r.GetPlugCable(p)

	if c == nil {
		r.currentCable = NewCable(p)
		r.addCable(r.currentCable)
		return
	}

	r.deleteConnection(c)

	if c.GetSrc() == p {
		c.SetSrc(c.GetDst())
	}

	c.SetDst(nil)
	r.currentCable = c
}

func (r *Rack) onConnectionStop(e event.IEvent) {
	if r.currentCable != nil && r.currentCable.GetDst() == nil {
		r.removeCable(r.currentCable)
		r.currentCable = nil
		return
	}

	r.createConnection(r.currentCable)

	r.currentCable = nil
}

func (r *Rack) onConnectionEnter(e event.IEvent) {
	if r.currentCable != nil {
		p := e.GetSource().(*Plug)
		if r.IsPlugFree(p) {
			r.currentCable.SetDst(p)
		}
	}
}

func (r *Rack) onConnectionLeave(e event.IEvent) {
	if r.currentCable != nil {
		r.currentCable.SetDst(nil)
	}
}

func (r *Rack) createConnection(cable *Cable) {
	srcMod, srcPort := cable.GetSrc().GetBinding()
	dstMod, dstPort := cable.GetDst().GetBinding()

	if cable.GetSrc().GetDirection() == PlugDirectionIn {
		dstMod, dstPort, srcMod, srcPort = srcMod, srcPort, dstMod, dstPort
	}

	r.audioRack.CreateModuleConnection(srcMod, srcPort, dstMod, dstPort)
}

func (r *Rack) deleteConnection(cable *Cable) {
	srcMod, srcPort := cable.GetSrc().GetBinding()
	dstMod, dstPort := cable.GetDst().GetBinding()

	if cable.GetSrc().GetDirection() == PlugDirectionIn {
		dstMod, dstPort, srcMod, srcPort = srcMod, srcPort, dstMod, dstPort
	}

	r.audioRack.DeleteModuleConnection(srcMod, srcPort, dstMod, dstPort)
}

func (r *Rack) GetPlugCable(p *Plug) *Cable {
	for _, c := range r.cables {
		if c.GetSrc() == p || c.GetDst() == p {
			return c
		}
	}

	return nil
}

func (r *Rack) IsPlugFree(p *Plug) bool {
	return r.GetPlugCable(p) == nil
}

func (r *Rack) GetAudioRack() *audio.Rack {
	return r.audioRack
}
