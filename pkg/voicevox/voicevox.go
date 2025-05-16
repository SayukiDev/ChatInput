package voicevox

import (
	"ChatInput/pkg/voicevox/api"
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os/exec"
	"strings"
	"sync/atomic"
	"unicode"
)

const host = "127.0.0.1"
const port = "5001"

type VoiceVox struct {
	*api.Api
	running       atomic.Bool
	closed        chan struct{}
	log           []string
	reader        *io.PipeReader
	writer        *io.PipeWriter
	LogUpdateHook func([]string)
	StartedHook   func(runed bool)
	runed         bool
	cmd           []string
	process       *exec.Cmd
}

func New(path string, lineLimit int, options ...string) *VoiceVox {
	args := []string{
		path,
		"--host", host,
		"--port", port,
	}
	args = append(args, options...)
	l := make([]string, 0, lineLimit)
	l = append(l, strings.Join(args, ""))
	return &VoiceVox{
		Api: api.New("http://" + host + ":" + port + "/"),
		log: l,
		cmd: args,
	}
}

func (v *VoiceVox) SetLogUpdateHook(hook func([]string)) {
	v.LogUpdateHook = hook
}

func (v *VoiceVox) Log() []string {
	return v.log
}

func (v *VoiceVox) readToLogLoop() {
	br := bufio.NewReader(v.reader)
	runned := false
	for {
		line, err := br.ReadString('\n')
		if errors.Is(err, io.ErrClosedPipe) {
			close(v.closed)
			return
		}
		if err != nil {
			log.WithError(err).Error("Read log failed")
		}
		if len(line) == 0 {
			continue
		}
		c := true
		for _, r := range []rune(line) {
			if !unicode.IsSpace(r) {
				c = false
			}
		}
		if c {
			continue
		}
		if len(v.log) == cap(v.log) {
			newLog := make([]string, 0, len(v.log))
			newLog = append(newLog, v.log[len(v.log)-1+(len(v.log)/5):]...)
			v.log = newLog
		}
		if strings.Contains(line, "running") {
			v.StartedHook(v.runed)
			runned = true
			v.runed = true
		}
		v.log = append(v.log, line)
		if runned {
			v.LogUpdateHook(v.log)
		}
	}
}

func (v *VoiceVox) SetStartedHook(hook func(runed bool)) {
	v.StartedHook = hook
}

func (v *VoiceVox) Running() bool {
	return v.running.Load()
}

func (v *VoiceVox) Start() error {
	if v.running.Load() {
		return nil
	}
	v.closed = make(chan struct{})
	v.running.Store(true)
	v.reader, v.writer = io.Pipe()
	v.process = exec.Command(v.cmd[0], v.cmd[1:]...)
	v.process.Stdout = v.writer
	v.process.Stderr = v.writer
	if err := v.process.Start(); err != nil {
		return err
	}
	go v.readToLogLoop()
	return nil
}

func (v *VoiceVox) Close() error {
	if !v.running.Load() {
		return nil
	}
	v.running.Store(false)
	err := v.reader.Close()
	if err != nil {
		return err
	}
	err = v.writer.Close()
	if err != nil {
		return err
	}
	v.process.Process.Kill()
	_ = v.process.Wait()
	<-v.closed
	return nil
}
