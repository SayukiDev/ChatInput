package ui

import (
	"ChatInput/options"
	"ChatInput/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"time"
)

type Ui struct {
	app               fyne.App
	opt               *options.Options
	srv               *service.Service
	lastSendInputting time.Time
	mw                fyne.Window
}

func New(opt *options.Options, srv *service.Service) *Ui {
	return &Ui{
		app: app.New(),
		opt: opt,
		srv: srv,
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
