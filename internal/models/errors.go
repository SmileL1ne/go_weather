package models

import "errors"

var (
	ErrNotFound     = errors.New("no weather info found by requested city")
	ErrNotAvailable = errors.New("open weather api is not available")
)

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}
