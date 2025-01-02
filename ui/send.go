package ui

import (
	"ChatInput/beep"
	"strings"
)

func (u *Ui) sendChatboxMsg(text string) error {
	text = strings.ReplaceAll(text, "w", "笑い")
	err := u.osc.ChatBoxInput(text, true, true)
	if err != nil {
		return err
	}
	if u.opt.TTS && text != "" {
		b, err := u.vox.TTS(text, u.opt.VoiceVox.Speaker)
		if err != nil {
			return err
		}

		// voice toggle, but working is so bad, so disabled for now.
		/*err = u.osc.SendRaw(osc.NewMessage("/input/Voice", int32(0)))
		if err != nil {
			return err
		}*/
		err = beep.Play(b, "wav", func() {
			/*err = u.osc.SendRaw(osc.NewMessage("/input/Voice", int32(1)))
			if err != nil {
				log.Println("Warn: [ Send voice toggle error:", err, "]")
			}*/
		})
		if err != nil {
			return err
		}
	}
	return nil
}