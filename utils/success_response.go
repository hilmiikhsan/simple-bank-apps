package utils

type SuccessResponse interface {
	StatusCode() int
	Message() string
}

type SuccessResponseWithData interface {
	StatusCode() int
	Message() string
	Data() interface{}
}

type MessageSuccess struct {
	SuccessStatusCode int    `json:"code"`
	SuccessMessage    string `json:"message"`
}

type MessageSuccessWithData struct {
	SuccessStatusCode int         `json:"code"`
	ResultData        interface{} `json:"data"`
	SuccessMessage    string      `json:"message"`
}

func (e *MessageSuccess) Message() string {
	return e.SuccessMessage
}

func (e *MessageSuccess) StatusCode() int {
	return e.SuccessStatusCode
}

func (e *MessageSuccessWithData) Message() string {
	return e.SuccessMessage
}

func (e *MessageSuccessWithData) StatusCode() int {
	return e.SuccessStatusCode
}

func (e *MessageSuccessWithData) Data() interface{} {
	return e.Data
}

func NewSuccessResponse(code int, message string) SuccessResponse {
	return &MessageSuccess{
		SuccessStatusCode: code,
		SuccessMessage:    message,
	}
}

func NewSuccessResponseWithData(code int, data interface{}, message string) SuccessResponseWithData {
	return &MessageSuccessWithData{
		SuccessStatusCode: code,
		ResultData:        data,
		SuccessMessage:    message,
	}
}
