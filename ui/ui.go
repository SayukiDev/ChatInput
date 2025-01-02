package ui

import (
	"ChatInput/options"
	"ChatInput/voicevox"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/SayukiDev/VRCOSC"
)

type Ui struct {
	app fyne.App
	opt *options.Options
	osc *VRCOSC.VRCOsc
	vox *voicevox.VoiceVox
	mw  fyne.Window
}

func New(opt *options.Options, vox *voicevox.VoiceVox) *Ui {
	return &Ui{
		app: app.New(),
		opt: opt,
		osc: VRCOSC.New(&VRCOSC.Options{
			SendPort: opt.SendPort,
			RecvPort: opt.RecvPort,
		}),
		vox: vox,
	}
}

func (u *Ui) Run() error {
	var err error
	u.mw, err = u.buildMainWindow()
	if err != nil {
		return err
	}
	u.mw.ShowAndRun()
	return nil
}
