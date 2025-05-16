package api

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

type ErrorRsp struct {
	ErrorMessage string `json:"errorMessage"`
}

func errorCheck(r *resty.Response, err error, checkBody bool) error {
	if err != nil {
		return err
	}
	if r.StatusCode() < 400 {
		return nil
	}
	e := ErrorRsp{}
	err = json.Unmarshal(r.Body(), &e)
	if err != nil {
		return err
	}
	return errors.New(e.ErrorMessage)
}
