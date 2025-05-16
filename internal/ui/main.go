package ui

import (
	tab2 "ChatInput/internal/ui/tab"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func (u *Ui) buildMainWindow() (fyne.Window, error) {
	a := app.New()
	w := a.NewWindow("ChatInput")
	w.Resize(fyne.NewSize(640, 460))
	c := container.NewAppTabs(
		tab2.NewInputTab(u.srv, w),
		tab2.NewOptionsTab(u.srv, w),
		tab2.NewVoiceVoxTab(u.srv, w),
	)
	w.SetContent(c)
	return w, nil
}
