package service

import (
	"ChatInput/api/voicevox"
	"ChatInput/options"
	"github.com/SayukiDev/VRCOSC"
)

type Service struct {
	osc *VRCOSC.VRCOsc
	vv  *voicevox.VoiceVox
	opt *options.Options
}

func New(opt *options.Options) (*Service, error) {
	s := &Service{
		opt: opt,
	}
	s.initOsc(opt)
	s.initVoiceVox(opt)
	return s, nil
}
