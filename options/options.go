package options

import (
	"encoding/json"
	"os"
	"sync"
)

type Options struct {
	// internal
	path  string
	lock  sync.Mutex
	hooks []HookFunc

	SendPort        int      `json:"send_port"`
	RecvPort        int      `json:"recv_port"`
	EnableTypingMsg bool     `json:"enable_typing_msg"`
	RealtimeSend    bool     `json:"realtime"`
	TTS             bool     `json:"tts"`
	VoiceControl    bool     `json:"voice_control"`
	VoiceVox        VoiceVox `json:"voicevox"`
}

type HookFunc func(o *Options) error

type VoiceVox struct {
	Path      string   `json:"path"`
	LineLimit int      `json:"line_limit"`
	Selected  int      `json:"selected"`
	Args      []string `json:"args"`
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
			Path:      "./windows-nvidia/run.exe",
			LineLimit: 50,
			Selected:  -1,
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

func (o *Options) AddHook(f HookFunc) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.hooks = append(o.hooks, f)
	return
}

func (o *Options) Updated() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	for _, hook := range o.hooks {
		if err := hook(o); err != nil {
			return err
		}
	}
	return nil
}
