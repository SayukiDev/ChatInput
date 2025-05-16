package api

import (
	"fmt"
	"testing"
)

func TestLVox_ListSpeaker(t *testing.T) {
	ss, err := l.ListSpeaker()
	if err != nil {
		t.Fatal(err)
	}
	msg := "Speakers:\n"
	for i, sp := range ss {
		var types string
		for _, s := range sp.Styles {
			types += fmt.Sprintf("Type: %s ID: %d\n", s.Name, s.ID)
		}
		msg += fmt.Sprintf(
			"SpeakerName: %s\n%s",
			sp.Name, types,
		)
		if i != len(ss)-1 {
			msg += "--------------\n"
		}
	}
	t.Log(msg)
}
