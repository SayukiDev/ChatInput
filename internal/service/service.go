package service

import (
	"ChatInput/options"
	"ChatInput/pkg/tasks"
	"ChatInput/pkg/voicevox"
	"github.com/SayukiDev/VRCOSC"
)

type Service struct {
	OSC    *VRCOSC.VRCOsc
	VV     *voicevox.VoiceVox
	Option *options.Options
	Tasks  *tasks.Tasks
}

func New(opt *options.Options) (*Service, error) {
	s := &Service{
		Option: opt,
		Tasks:  tasks.New(),
	}
	s.initOsc(opt)
	s.initVoiceVox(opt)
	return s, nil
}

func (s *Service) Close() error {
	return s.VV.Close()
}
