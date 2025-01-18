package service

import (
	"ChatInput/beep"
	"strings"
)

func (s *Service) SendChatboxMsg(text string, tts bool) error {
	err := s.osc.ChatBoxInput(text, true, true)
	if err != nil {
		return err
	}
	text = strings.ReplaceAll(text, "w", "笑い")
	if s.opt.TTS && text != "" && tts {
		b, err := s.vv.TTS(text, s.opt.VoiceVox.Speaker)
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
