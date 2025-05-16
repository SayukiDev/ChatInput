package main

import (
	"ChatInput/internal/service"
	"ChatInput/internal/ui"
	"ChatInput/options"
	"github.com/sirupsen/logrus"
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
	u := ui.New(srv)
	err = u.Run()
	if err != nil {
		panic(err)
	}

	logrus.Println("Shutdown...")
	err = srv.Close()
	if err != nil {
		panic(err)
	}
	logrus.Println("Done.")
}
