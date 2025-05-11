package ui

import (
	"ChatInput/common/widget/entry"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
	"time"
)

func (u *Ui) buildMainWindow() (fyne.Window, error) {
	a := app.New()
	w := a.NewWindow("ChatInput")
	w.Resize(fyne.NewSize(640, 460))
	it, err := u.buildInputTab(w)
	if err != nil {
		return nil, err
	}
	ot, err := u.buildOptionsTab(w)
	if err != nil {
		return nil, err
	}
	c := container.NewAppTabs(it, ot)
	w.SetContent(c)
	return w, nil
}

func (u *Ui) buildInputTab(w fyne.Window) (*container.TabItem, error) {
	e := entry.NewExtEntry()
	e.OnFocusGained = func() bool {
		if !u.opt.EnableTypingMsg {
			return false
		}
		err := u.srv.Tasks.Add("Typing", func(c chan struct{}) {
			for ; ; <-time.Tick(10 * time.Second) {
				select {
				case <-c:
					return
				default:
				}
				if !u.opt.EnableTypingMsg {
					return
				}
				if e.Text == "" {
					continue
				}
				err := u.srv.SendChatboxMsg("入力中...", false, true)
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
		if !u.opt.EnableTypingMsg {
			return false
		}
		u.srv.Tasks.Remove("Typing")
		return false
	}
	e.OnTypedKey = func(event *fyne.KeyEvent) bool {
		if event.Name == fyne.KeyReturn {
			text := e.Text
			e.SetText("")
			err := u.srv.SendChatboxMsg(strings.ReplaceAll(text, "\n", ""), true, false)
			if err != nil {
				dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
				return true
			}
		}
		return true
	}
	clear := widget.NewButton("Clear", func() {
		err := u.srv.SendChatboxMsg("", false, false)
		if err != nil {
			dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
			return
		}
		e.SetText("")
	})
	send := widget.NewButton("Send", func() {
		text := e.Text
		err := u.srv.SendChatboxMsg(text, true, false)
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
	return container.NewTabItem("Input", c), nil
}

func (u *Ui) buildOptionsTab(w fyne.Window) (*container.TabItem, error) {
	sendPort := widget.NewEntry()
	sendPort.SetText(strconv.Itoa(u.opt.SendPort))
	revcPort := widget.NewEntry()
	revcPort.SetText(strconv.Itoa(u.opt.RecvPort))
	etmV := u.opt.EnableTypingMsg
	etm := newOnOffRadio("On", "Off", &etmV)
	ttsV := u.opt.TTS
	tts := newOnOffRadio("On", "Off", &ttsV)
	rsV := u.opt.RealtimeSend
	rs := newOnOffRadio("On", "Off", &rsV)
	voiceV := u.opt.VoiceControl
	voice := newOnOffRadio("On", "Off", &voiceV)
	f := widget.NewForm(
		widget.NewFormItem("Send Port", sendPort),
		widget.NewFormItem("Revc Port", revcPort),
		widget.NewFormItem("Enable Typing Message", etm),
		widget.NewFormItem("TTS", tts),
		widget.NewFormItem("Realtime Send", rs),
		widget.NewFormItem("Voice Control", voice),
	)
	f.SubmitText = "Save"
	f.OnSubmit = func() {
		sp, err := strconv.Atoi(sendPort.Text)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		rp, err := strconv.Atoi(revcPort.Text)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		u.opt.SendPort = sp
		u.opt.RecvPort = rp
		u.opt.EnableTypingMsg = etmV
		u.opt.TTS = ttsV
		u.opt.RealtimeSend = rsV
		u.opt.VoiceControl = voiceV
		err = u.opt.Save()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		err = u.opt.Updated()
		if err != nil {
			dialog.ShowError(fmt.Errorf("triggering options update hooks error: %s", err), w)
		}
		dialog.ShowInformation("Saved", "Options should be saved", w)
	}
	return container.NewTabItem("Options", f), nil
}
