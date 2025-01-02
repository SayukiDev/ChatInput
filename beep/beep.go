package beep

import (
	"bytes"
	"fmt"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
	"io"
	"sync"
	"time"
)

var lock sync.Mutex
var playing bool

type bytesReadCloser struct {
	*bytes.Reader
}

func (rc bytesReadCloser) Close() error {
	return nil
}

func Play(b []byte, mediaType string, callback func()) error {
	return PlayFromReader(&bytesReadCloser{Reader: bytes.NewReader(b)}, mediaType, callback)
}

func PlayFromReader(r io.ReadCloser, mediaType string, callback func()) error {
	var (
		s   beep.StreamSeekCloser
		f   beep.Format
		err error
	)
	switch mediaType {
	case "wav":
		s, f, err = wav.Decode(r)
		if err != nil {
			return err
		}
	case "mp3":
		s, f, err = mp3.Decode(r)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported media type: %s", mediaType)
	}
	lock.Lock()
	if !playing {
		err = speaker.Init(f.SampleRate, f.SampleRate.N(time.Second/10))
		if err != nil {
			return err
		}
	}
	playing = true
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		defer s.Close()
		callback()
		lock.Unlock()
	})))
	return nil
}
