package service

import (
	"ChatInput/options"
	"ChatInput/pkg/voicevox"
	log "github.com/sirupsen/logrus"
)

func (s *Service) initVoiceVox(opt *options.Options) {
	s.VV = voicevox.New(opt.VoiceVox.Path, opt.VoiceVox.LineLimit, opt.VoiceVox.Args...)
	s.VV.SetSpeaker(opt.VoiceVox.Selected)
	opt.AddHook(func(o *options.Options) error {
		run := false
		if s.VV.Running() {
			run = true
		}
		s.VV = voicevox.New(o.VoiceVox.Path, o.VoiceVox.LineLimit, o.VoiceVox.Args...)
		if run {
			err := s.VV.Start()
			if err != nil {
				log.WithError(err).Error("Restart VoiceVox failed")
			}
		}
		s.VV.SetSpeaker(opt.VoiceVox.Selected)
		return nil
	})
	return
}
