package voicevox

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type VoiceVox struct {
	c *resty.Client
}

func New(baseurl string) *VoiceVox {
	return &VoiceVox{
		c: resty.New().
			SetBaseURL(baseurl).
			SetTimeout(20 * time.Second),
	}
}
