package beep

import (
	"os"
	"testing"
)

func TestPlay(t *testing.T) {
	f, err := os.Open("./test_data/1.mp3")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	done := make(chan bool)
	err = PlayFromReader(f, "mp3", func() {
		done <- true
	})
	if err != nil {
		t.Error(err)
		return
	}
	<-done
}
