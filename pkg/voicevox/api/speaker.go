package api

import "encoding/json"

type Speakers []struct {
	Name        string   `json:"name"`
	SpeakerUUID string   `json:"speaker_uuid"`
	Styles      []Styles `json:"styles"`
	Version     string   `json:"version"`
}

func (l *Api) ListSpeaker() (Speakers, error) {
	r, err := l.c.R().Get("/speakers")
	err = errorCheck(r, err, true)
	if err != nil {
		return nil, err
	}
	var speakers Speakers
	if err = json.Unmarshal(r.Body(), &speakers); err != nil {
		return nil, err
	}
	return speakers, nil
}

func (l *Api) SetSpeaker(id int) {
	l.speaker = id
}
