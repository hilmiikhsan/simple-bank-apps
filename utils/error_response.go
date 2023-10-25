package utils

import "net/http"

const (
	STATUS_INTERNAL_ERR = "STATUS_INTERNAL_ERROR"
	STATUS_BAD_REQUEST  = "STATUS_BAD_REQUEST"
	STATUS_UNAUTHORIZED = "STATUS_UNAUTHORIZED"
	STATUS_NOT_FOUND    = "STATUS_NOT_FOUND"
)

type ErrorResponse interface {
	StatusCode() int
	Message() string
	Error() string
}

type MessageErrData struct {
	ErrStatusCode int    `json:"code"`
	ErrMessage    string `json:"message"`
	ErrError      string `json:"error"`
}

func (e *MessageErrData) Message() string {
	return e.ErrMessage
}

func (e *MessageErrData) StatusCode() int {
	return e.ErrStatusCode
}

func (e *MessageErrData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) ErrorResponse {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      STATUS_INTERNAL_ERR,
	}
}

func NewBadRequest(message string) ErrorResponse {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusBadRequest,
		ErrError:      STATUS_BAD_REQUEST,
	}
}

func NewNotFound(message string) ErrorResponse {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusNotFound,
		ErrError:      STATUS_NOT_FOUND,
	}
}

func NewUnauthorized(message string) ErrorResponse {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnauthorized,
		ErrError:      STATUS_UNAUTHORIZED,
	}
}
