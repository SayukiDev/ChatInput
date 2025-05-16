package voicevox

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestVoiceVox_Start(t *testing.T) {
	v := New("C:\\Users\\Sayuki\\Downloads\\windows-directml\\run.exe", 56)
	err := v.Start()
	if err != nil {
		t.Fatal(err)
	}
	v.SetLogUpdateHook(func(l []string) {
		println(strings.Join(l, "\n"))
	})
	v.SetStartedHook(func(_ bool) {
		ss, err := v.ListSpeaker()
		if err != nil {
			t.Fatal(err)
		}
		log.Print(ss)
	})
	select {}
	/*<-time.Tick(10 * time.Second)
	v.Close()*/
}
