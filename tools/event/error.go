package event

import (
	"net/http"
	"strings"

	"github.com/dlbarduzzi/sentinel/tools/inflector"
)

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Error makes it compatible with the `error` interface.
func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(status int, message string) *ApiError {
	message = strings.TrimSpace(message)
	if message == "" {
		message = http.StatusText(status)
	}

	return &ApiError{
		Status:  status,
		Message: inflector.FormatSentence(message),
	}
}

func NewInternalServerError(message string) *ApiError {
	message = strings.TrimSpace(message)
	if message == "" {
		message = "Something went wrong while processing this request."
	}

	return NewApiError(http.StatusInternalServerError, message)
}
