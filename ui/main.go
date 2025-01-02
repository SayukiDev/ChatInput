package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/SayukiDev/VRCOSC"
	"strconv"
	"strings"
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
	e := widget.NewMultiLineEntry()
	e.OnChanged = func(text string) {
		if text == "" {
			return
		}
		if !strings.HasSuffix(text, "\n") {
			// vrchat have limit, so disabled for now, maybe next commit will add cooldown to resole the issue
			/*err := u.osc.ChatBoxInput("入力中...", true, true)
			if err != nil {
				dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
				return
			}*/
			return
		}
		e.SetText("")
		err := u.sendChatboxMsg(strings.TrimSuffix(text, "\n"))
		if err != nil {
			dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
			return
		}
	}
	clear := widget.NewButton("Clear", func() {
		err := u.sendChatboxMsg("")
		if err != nil {
			dialog.ShowError(fmt.Errorf("send msg error: %s", err), w)
			return
		}
		e.SetText("")
	})
	send := widget.NewButton("Send", func() {
		text := e.Text
		err := u.sendChatboxMsg(text)
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
	ttsV := u.opt.TTS
	tts := newOnOffRadio("On", "Off", &ttsV)
	rsV := u.opt.RealtimeSend
	rs := newOnOffRadio("On", "Off", &rsV)
	voiceV := u.opt.TTS
	voice := newOnOffRadio("On", "Off", &voiceV)
	f := widget.NewForm(
		widget.NewFormItem("Send Port", sendPort),
		widget.NewFormItem("Revc Port", revcPort),
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
		oscNeedReset := u.opt.SendPort != sp || u.opt.RecvPort != rp
		u.opt.SendPort = sp
		u.opt.RecvPort = rp
		if oscNeedReset {
			u.osc = VRCOSC.New(&VRCOSC.Options{
				SendPort: sp,
				RecvPort: rp,
			})
			return
		}
		u.opt.TTS = ttsV
		u.opt.RealtimeSend = rsV
		u.opt.VoiceControl = voiceV
		err = u.opt.Save()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Saved", "Options should be saved", w)
	}
	return container.NewTabItem("Options", f), nil
}
