package api

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"strconv"
)

type SynthParams struct {
	AccentPhrases      []AccentPhrases `json:"accent_phrases"`
	SpeedScale         float64         `json:"speedScale"`
	PitchScale         float64         `json:"pitchScale"`
	IntonationScale    float64         `json:"intonationScale"`
	VolumeScale        float64         `json:"volumeScale"`
	PrePhonemeLength   float64         `json:"prePhonemeLength"`
	PostPhonemeLength  float64         `json:"postPhonemeLength"`
	OutputSamplingRate int             `json:"outputSamplingRate"`
	OutputStereo       bool            `json:"outputStereo"`
	Kana               string          `json:"kana"`
}

type Mora struct {
	Text            string   `json:"text"`
	Consonant       *string  `json:"consonant"`
	ConsonantLength *float64 `json:"consonant_length"`
	Vowel           string   `json:"vowel"`
	VowelLength     float64  `json:"vowel_length"`
	Pitch           float64  `json:"pitch"`
}

type AccentPhrases struct {
	Moras           []Mora `json:"moras"`
	Accent          int    `json:"accent"`
	PauseMora       *Mora  `json:"pause_mora"`
	IsInterrogative bool   `json:"is_interrogative"`
}

type Styles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (l *Api) getQuery(id int, text string) (*SynthParams, error) {
	r, err := l.c.R().SetQueryParams(map[string]string{
		"speaker": strconv.Itoa(id),
		"text":    text,
	}).Post("/audio_query")
	if err != nil {
		return nil, err
	}
	err = errorCheck(r, err, true)
	if err != nil {
		log.Error(r.String())
		return nil, err
	}
	var params *SynthParams
	if err = json.Unmarshal(r.Body(), &params); err != nil {
		return nil, err
	}
	return params, nil
}

func (l *Api) synth(id int, params *SynthParams) ([]byte, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	r, err := l.c.R().
		SetDoNotParseResponse(true).
		SetHeaders(map[string]string{
			"Accept":       "audio/wav",
			"Content-Type": "application/json",
		}).
		SetQueryParam("speaker", strconv.Itoa(id)).
		SetBody(b).
		Post("/synthesis")
	if err != nil {
		return nil, err
	}
	defer r.RawBody().Close()
	err = errorCheck(r, err, true)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, r.RawBody()); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (l *Api) TTS(text string) ([]byte, error) {
	params, err := l.getQuery(l.speaker, text)
	if err != nil {
		return nil, err
	}
	params.IntonationScale = 1
	params.SpeedScale = 1
	params.PitchScale = 0
	params.VolumeScale = 2
	b, err := l.synth(l.speaker, params)
	if err != nil {
		return nil, err
	}
	return b, err
}
