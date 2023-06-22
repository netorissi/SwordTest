package shared

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidBind    = errors.New("invalid request")
	ErrInvalidPayload = errors.New("invalid payload")
)

type Error struct {
	HTTPCode int    `json:"code"`
	Message  string `json:"message"`
}

func (e *Error) Error() (msgErr string) {
	if e.HTTPCode > 0 {
		msgErr += fmt.Sprintf("HTTP Code: %v. ", e.HTTPCode)
	}

	if len(e.Message) > 0 {
		msgErr += fmt.Sprintf("Message: %v. ", e.Message)
	}

	return
}

func NewError(httpCode int, message string) error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
	}
}

func GetHTTPCode(err error) int {
	e := GetError(err)
	if e == nil {
		return 0
	}

	return e.HTTPCode
}

func GetError(err error) *Error {
	if err == nil {
		return nil
	}

	e, ok := err.(*Error)
	if !ok {
		return &Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  err.Error(),
		}
	}

	return e
}
