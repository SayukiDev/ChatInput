package service

import (
	"github.com/SayukiDev/Beep"
)

func (s *Service) SendChatboxMsg(text string, tts bool, disableSfx bool) error {
	err := s.osc.ChatBoxInput(text, true, !disableSfx)
	if err != nil {
		return err
	}
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
		err = Beep.Play(b, "wav", func() {
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
