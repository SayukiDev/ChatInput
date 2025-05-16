package tab

import (
	"ChatInput/internal/service"
	"ChatInput/internal/ui/widget/entry"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strings"
	"time"
)

func NewInputTab(s *service.Service, w fyne.Window) *container.TabItem {
	e := entry.NewExtEntry()
	e.OnFocusGained = func() bool {
		if !s.Option.EnableTypingMsg {
			return false
		}
		err := s.Tasks.Add("Typing", func(c chan struct{}) {
			for ; ; <-time.Tick(10 * time.Second) {
				select {
				case <-c:
					return
				default:
				}
				if !s.Option.EnableTypingMsg {
					return
				}
				if e.Text == "" {
					continue
				}
				err := s.SendChatboxMsg("入力中...", false, true)
				if err != nil {
					dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
					return
				}
			}
		})
		if err != nil {
			dialog.ShowError(fmt.Errorf("add task error: %s", err), w)
		}
		return false
	}
	e.OnFocusLost = func() bool {
		if !s.Option.EnableTypingMsg {
			return false
		}
		s.Tasks.Remove("Typing")
		return false
	}
	e.OnTypedKey = func(event *fyne.KeyEvent) bool {
		if event.Name == fyne.KeyReturn {
			text := e.Text
			e.SetText("")
			err := s.SendChatboxMsg(strings.ReplaceAll(text, "\n", ""), true, false)
			if err != nil {
				dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
				return true
			}
		}
		return true
	}
	clear := widget.NewButton("Close", func() {
		err := s.SendChatboxMsg("", false, false)
		if err != nil {
			dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
			return
		}
		e.SetText("")
	})
	send := widget.NewButton("Send", func() {
		text := e.Text
		err := s.SendChatboxMsg(text, true, false)
		if err != nil {
			dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
			return
		}
		e.SetText("")
	})
	c := container.NewGridWithRows(2,
		e,
		container.NewGridWithRows(3,
			container.NewGridWithColumns(3,
				clear,
				send,
			),
		))
	return container.NewTabItem("Input", c)
}
