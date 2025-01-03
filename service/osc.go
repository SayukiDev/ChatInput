package service

import (
	"ChatInput/options"
	"github.com/SayukiDev/VRCOSC"
)

func (s *Service) initOsc(opt *options.Options) {
	s.osc = VRCOSC.New(&VRCOSC.Options{
		SendPort: opt.SendPort,
		RecvPort: opt.RecvPort,
	})
	opt.AddHook(func(o *options.Options) error {
		s.osc = VRCOSC.New(&VRCOSC.Options{
			SendPort: o.SendPort,
			RecvPort: o.RecvPort,
		})
		return nil
	})
}