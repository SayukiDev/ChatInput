package api

import (
	"testing"
)

func TestLVox_TTS(t *testing.T) {
	_, err := l.TTS("あいうえお", 4)
	if err != nil {
		t.Fatal(err)
	}
}
