package ui

import (
	"ChatInput/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Ui struct {
	app fyne.App
	srv *service.Service
	mw  fyne.Window
}

func New(srv *service.Service) *Ui {
	return &Ui{
		app: app.New(),
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
