package main

import (
	"ChatInput/options"
	"ChatInput/ui"
	"ChatInput/voicevox"
)

func main() {
	opt := options.NewOptions("./data.json")
	err := opt.Load()
	if err != nil {
		panic(err)
	}
	vv := voicevox.New(opt.VoiceVox.Address)
	u := ui.New(opt, vv)
	err = u.Run()
	if err != nil {
		panic(err)
	}
}
