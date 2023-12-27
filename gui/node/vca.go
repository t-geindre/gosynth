package node

import (
	"gosynth/event"
	"gosynth/gui/theme"
	"gosynth/module"
)

type VCA struct {
	*Module
	audioVca *module.VCA
	slider   *Slider
}

func NewVCA(audioVca *module.VCA) *VCA {
	v := &VCA{}
	width, height := 65, 500
	v.Module = NewModule(width, height, v)

	v.audioVca = audioVca

	v.slider = NewSlider()
	v.slider.SetRange(-1, 1)
	v.slider.SetValue(0.5)
	v.AppendWithOptions(v.slider, NewAppendOptions().HorizontallyFill(100).VerticallyFill(100))

	lineToCv := NewLine(10, 1, LineOrientationVertical)
	v.AppendWithOptions(lineToCv, NewAppendOptions().HorizontallyCentered())

	cvPlug := NewPlug()
	v.AppendWithOptions(cvPlug, NewAppendOptions().HorizontallyCentered())

	cvLabel := NewLabel("CV", theme.Fonts.Small)
	v.AppendWithOptions(cvLabel, NewAppendOptions().HorizontallyCentered().Margins(3, 0, 0, 0))

	separatorLine := NewLine(10, 1, LineOrientationHorizontal)
	v.AppendWithOptions(
		separatorLine,
		NewAppendOptions().
			HorizontallyCentered().
			HorizontallyFill(100).
			Margins(15, 15, 10, 10),
	)

	inLabel := NewLabel("IN", theme.Fonts.Small)
	v.AppendWithOptions(inLabel, NewAppendOptions().HorizontallyCentered().Margins(0, 3, 0, 0))

	inPlug := NewPlug()
	v.AppendWithOptions(inPlug, NewAppendOptions().HorizontallyCentered())

	inOutLine := NewLine(10, 1, LineOrientationVertical)
	v.AppendWithOptions(inOutLine, NewAppendOptions().HorizontallyCentered())

	outPlugLabel := NewLabel("OUT", theme.Fonts.Small)

	outPlug := NewPlug()
	outPlugContainer := NewContainer(
		// TODO some kind of fluid container would be nice
		outPlug.GetOuterWidth()+10, outPlug.GetOuterHeight()+outPlugLabel.GetOuterHeight()+17,
	)
	outPlugContainer.SetInverted(true)

	v.AppendWithOptions(outPlugContainer, NewAppendOptions().HorizontallyCentered().Padding(5))

	outPlugContainer.AppendWithOptions(outPlugLabel, NewAppendOptions().HorizontallyCentered().Margin(3))
	outPlugContainer.AppendWithOptions(outPlug, NewAppendOptions().HorizontallyCentered())

	v.Dispatcher.AddListener(&v, ValueChangedEvent, v.OnSliderValueChanged)

	return v
}

func (v *VCA) Update() error {
	v.slider.SetValue(v.audioVca.GetGain())
	return v.Module.Update()
}

func (v *VCA) OnSliderValueChanged(e event.IEvent) {
	e.StopPropagation()
	v.audioVca.SendCommand(module.PortInGain, e.GetSource().(*Slider).GetValue())
}
