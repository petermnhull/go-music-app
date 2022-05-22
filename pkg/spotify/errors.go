package spotify

import (
	"fmt"
)

type SpotifyRequestErr struct {
	Message string
}

func (e *SpotifyRequestErr) Error() string {
	return e.Message
}

func NewSpotifyRequestErr(msg string) *SpotifyRequestErr {
	return &SpotifyRequestErr{Message: msg}
}

type SpotifyResponseErr struct {
	Message string
	Code    int
}

func (e *SpotifyResponseErr) Error() string {
	return fmt.Sprintf("external request to spotify failed (%d): %s", e.Code, e.Message)
}

func NewSpotifyResponseErr(msg string, code int) *SpotifyResponseErr {
	return &SpotifyResponseErr{Message: msg, Code: code}
}

type SpotifyValidationErr struct {
	Message string
}

func (e *SpotifyValidationErr) Error() string {
	return e.Message
}

func NewSpotifyValidationErr(msg string) *SpotifyValidationErr {
	return &SpotifyValidationErr{Message: msg}
}
