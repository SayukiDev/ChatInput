package service

import (
	"ChatInput/api/voicevox"
	"ChatInput/options"
)

func (s *Service) initVoiceVox(opt *options.Options) {
	s.vv = voicevox.New(opt.VoiceVox.Address)
	opt.AddHook(func(o *options.Options) error {
		s.vv = voicevox.New(opt.VoiceVox.Address)
		return nil
	})
	return
}
