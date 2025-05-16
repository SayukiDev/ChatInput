package tab

import (
	"ChatInput/internal/service"
	"ChatInput/internal/ui/widget/radio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func NewOptionsTab(s *service.Service, w fyne.Window) *container.TabItem {
	sendPort := widget.NewEntry()
	sendPort.SetText(strconv.Itoa(s.Option.SendPort))
	revcPort := widget.NewEntry()
	revcPort.SetText(strconv.Itoa(s.Option.RecvPort))
	etmV := s.Option.EnableTypingMsg
	etm := radio.NewOnOffRadio("On", "Off", &etmV)
	ttsV := s.Option.TTS
	tts := radio.NewOnOffRadio("On", "Off", &ttsV)
	rsV := s.Option.RealtimeSend
	rs := radio.NewOnOffRadio("On", "Off", &rsV)
	voiceV := s.Option.VoiceControl
	voice := radio.NewOnOffRadio("On", "Off", &voiceV)
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
		s.Option.SendPort = sp
		s.Option.RecvPort = rp
		s.Option.EnableTypingMsg = etmV
		s.Option.TTS = ttsV
		s.Option.RealtimeSend = rsV
		s.Option.VoiceControl = voiceV
		err = s.Option.Save()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		err = s.Option.Updated()
		if err != nil {
			dialog.ShowError(fmt.Errorf("triggering options update hooks error: %s", err), w)
		}
		dialog.ShowInformation("Saved", "Options should be saved", w)
	}
	return container.NewTabItem("Options", f)
}
