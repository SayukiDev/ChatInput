package service

import (
	"ChatInput/options"
	"ChatInput/pkg/voicevox"
)

func (s *Service) initVoiceVox(opt *options.Options) {
	s.VV = voicevox.New(opt.VoiceVox.Path, opt.VoiceVox.LineLimit, opt.VoiceVox.Args...)
	s.VV.SetSpeaker(opt.VoiceVox.Selected)
	opt.AddHook(func(o *options.Options) error {
		s.VV.SetSpeaker(opt.VoiceVox.Selected)
		return nil
	})
	return
}
