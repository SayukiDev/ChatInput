package api

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type Api struct {
	c       *resty.Client
	speaker int
}

func New(baseurl string) *Api {
	return &Api{
		c: resty.New().
			SetBaseURL(baseurl).
			SetTimeout(20 * time.Second),
	}
}
