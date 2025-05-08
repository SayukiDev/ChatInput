package service

import (
	"ChatInput/api/voicevox"
	"ChatInput/options"
	"ChatInput/tasks"
	"github.com/SayukiDev/VRCOSC"
)

type Service struct {
	osc   *VRCOSC.VRCOsc
	vv    *voicevox.VoiceVox
	opt   *options.Options
	Tasks *tasks.Tasks
}

func New(opt *options.Options) (*Service, error) {
	s := &Service{
		opt:   opt,
		Tasks: tasks.New(),
	}
	s.initOsc(opt)
	s.initVoiceVox(opt)
	return s, nil
}
