package node

import (
	"gosynth/event"
	"gosynth/gui/theme"
	"gosynth/module"
	"time"
)

type VCA struct {
	*Module
	audioVca *module.VCA
	slider   *Slider
	CvInPlug *Plug
	OutPlug  *Plug
	InPlug   *Plug
}

func NewVCA(audioVca *module.VCA) *VCA {
	v := &VCA{}
	width, height := 65, 500
	v.Module = NewModule(width, height, v)

	v.audioVca = audioVca

	v.slider = NewSlider()
	v.slider.SetRange(-1, 1)
	v.AppendWithOptions(v.slider, NewAppendOptions().HorizontallyFill(100).VerticallyFill(100))

	lineToCv := NewLine(10, 1, LineOrientationVertical)
	v.AppendWithOptions(lineToCv, NewAppendOptions().HorizontallyCentered())

	v.CvInPlug = NewPlug()
	v.AppendWithOptions(v.CvInPlug, NewAppendOptions().HorizontallyCentered())

	cvLabel := NewLabel("CV", theme.Fonts.Small)
	v.AppendWithOptions(cvLabel, NewAppendOptions().HorizontallyCentered().Margins(2, 0, 0, 0))

	separatorLine := NewLine(10, 1, LineOrientationHorizontal)
	v.AppendWithOptions(
		separatorLine,
		NewAppendOptions().
			HorizontallyCentered().
			HorizontallyFill(100).
			Margins(15, 15, 10, 10),
	)

	inLabel := NewLabel("IN", theme.Fonts.Small)
	v.AppendWithOptions(inLabel, NewAppendOptions().HorizontallyCentered().Margins(0, 2, 0, 0))

	v.InPlug = NewPlug()
	v.AppendWithOptions(v.InPlug, NewAppendOptions().HorizontallyCentered())

	inOutLine := NewLine(10, 1, LineOrientationVertical)
	v.AppendWithOptions(inOutLine, NewAppendOptions().HorizontallyCentered())

	outPlugLabel := NewLabel("OUT", theme.Fonts.Small)

	v.OutPlug = NewPlug()
	outPlugContainer := NewContainer(
		// TODO some kind of fluid container would be nice
		v.OutPlug.GetOuterWidth()+10, v.OutPlug.GetOuterHeight()+outPlugLabel.GetOuterHeight()+17,
	)
	outPlugContainer.SetInverted(true)

	v.AppendWithOptions(outPlugContainer, NewAppendOptions().HorizontallyCentered().Padding(5))

	outPlugContainer.AppendWithOptions(outPlugLabel, NewAppendOptions().HorizontallyCentered().Margins(3, 3, 0, 0))
	outPlugContainer.AppendWithOptions(v.OutPlug, NewAppendOptions().HorizontallyCentered())

	v.Dispatcher.AddListener(&v, ValueChangedEvent, v.OnSliderValueChanged)

	return v
}

func (v *VCA) Update(time time.Duration) error {
	cvVal := v.audioVca.ReceiveInput(module.PortCvIn)

	if cvVal != nil {
		v.CvInPlug.On()
	} else {
		v.CvInPlug.Off()
	}

	if cvVal != nil {
		v.slider.SetValue(*cvVal)
	}

	outVal := v.audioVca.ReceiveOutput(module.PortOut)
	if outVal != nil && *outVal != 0 {
		v.OutPlug.On()
	} else {
		v.OutPlug.Off()
	}

	inVal := v.audioVca.ReceiveInput(module.PortIn)
	if inVal != nil && *inVal != 0 {
		v.InPlug.On()
	} else {
		v.InPlug.Off()
	}

	return v.Module.Update(time)
}

func (v *VCA) OnSliderValueChanged(e event.IEvent) {
	// Todo if cv in has a connection, do nothing
	e.StopPropagation()
	v.audioVca.SendInput(module.PortCvIn, e.GetSource().(*Slider).GetValue())
}
