package radio

import "fyne.io/fyne/v2/widget"

func NewOnOffRadio(on, off string, v *bool) *widget.RadioGroup {
	var rP *widget.RadioGroup
	r := widget.NewRadioGroup([]string{on, off}, func(s string) {
		if s == on {
			*v = true
		} else {
			*v = false
		}
		if s == "" {
			rP.SetSelected(off)
		}
	})
	rP = r
	if *v {
		r.SetSelected(on)
	} else {
		r.SetSelected(off)
	}
	return r
}
