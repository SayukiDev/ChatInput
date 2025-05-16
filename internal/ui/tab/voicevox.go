package tab

import (
	"ChatInput/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
	"strings"
)

func NewVoiceVoxTab(s *service.Service, w fyne.Window) *container.TabItem {
	e := widget.NewMultiLineEntry()
	refresh := widget.NewButton("Refresh", func() {
		e.SetText(strings.Join(s.VV.Log(), "\n"))
	})
	{
		first := false
		s.VV.SetLogUpdateHook(func(i []string) {
			if !first {
				e.SetText(strings.Join(i, "\n"))
				first = true
			}
		})
	}
	start := widget.NewButton("Start", func() {
		err := s.VV.Start()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowConfirm("Started", "VoiceVox Started!", func(b bool) {}, w)
	})
	spsSelectMap := make(map[string]int)
	spsS := widget.NewSelect(nil, func(key string) {
		id := spsSelectMap[key]
		s.VV.SetSpeaker(id)
		s.Option.VoiceVox.Selected = id
		s.Option.Save()
	})
	spsS.PlaceHolder = "Select speaker"
	s.VV.SetStartedHook(func(_ bool) {
		var so []string
		ss, err := s.VV.ListSpeaker()
		if err != nil {
			log.WithError(err).Error("List Speaker failed")
			return
		}
		selected := ""
		for _, ssv := range ss {
			for _, sty := range ssv.Styles {
				name := ssv.Name + "-" + sty.Name
				spsSelectMap[name] = sty.ID
				so = append(so, name)
				if sty.ID == s.Option.VoiceVox.Selected {
					selected = name
				}
			}
		}
		spsS.SetOptions(so)
		if selected != "" {
			spsS.SetSelected(selected)
		}
	})

	stop := widget.NewButton("Stop", func() {
		if err := s.VV.Close(); err != nil {
			dialog.ShowError(err, w)
			return
		}
		spsS.SetOptions(nil)
		spsS.ClearSelected()
		e.SetText("")
		dialog.ShowConfirm("Stopped", "VoiceVox Stopped!", func(b bool) {}, w)
	})
	c := container.NewGridWithRows(2,
		e,
		container.NewGridWithRows(3,
			container.NewGridWithColumns(3,
				start,
				stop,
				refresh,
			),
			spsS,
		))
	return container.NewTabItem("VoiceVox", c)
}
