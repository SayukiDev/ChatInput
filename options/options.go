package options

import (
	"encoding/json"
	"os"
	"sync"
)

type Options struct {
	path         string
	lock         sync.Mutex
	SendPort     int      `json:"send_port"`
	RecvPort     int      `json:"recv_port"`
	RealtimeSend bool     `json:"realtime"`
	TTS          bool     `json:"tts"`
	VoiceControl bool     `json:"voice_control"`
	VoiceVox     VoiceVox `json:"voicevox"`
}

type VoiceVox struct {
	Address string `json:"address"`
	Speaker int    `json:"speaker"`
}

func NewOptions(p string) *Options {
	return &Options{
		path:         p,
		SendPort:     9000,
		RecvPort:     9001,
		RealtimeSend: false,
		TTS:          false,
		VoiceControl: false,
		VoiceVox: VoiceVox{
			Address: "http://127.0.0.1:50021",
			Speaker: 4,
		},
	}
}

func (o *Options) Load() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	file, err := os.Open(o.path)
	if err != nil {
		if os.IsNotExist(err) {
			o.lock.Unlock()
			defer o.lock.Lock()
			return o.Save()
		}
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(o)
}

func (o *Options) Save() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	file, err := os.OpenFile(o.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(o)
}
