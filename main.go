package main

import (
	"ChatInput/options"
	"ChatInput/service"
	"ChatInput/ui"
)

func main() {
	opt := options.NewOptions("./data.json")
	err := opt.Load()
	if err != nil {
		panic(err)
	}
	srv, err := service.New(opt)
	if err != nil {
		panic(err)
	}
	u := ui.New(opt, srv)
	err = u.Run()
	if err != nil {
		panic(err)
	}
}
