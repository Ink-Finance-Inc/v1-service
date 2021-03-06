package errors

import (
	"encoding/json"
)

var (
	EndpointTypeError = NewError(1000, "endpoints error!")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e *Error) String() string {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (e *Error) Error() string {
	return e.String()
}

func NewError(code int, msg string) error {
	e := &Error{
		Code:    code,
		Message: msg,
	}
	return e
}

func ParseError(data string) *Error {
	var terr Error
	err := json.Unmarshal([]byte(data), &terr)
	if err != nil {
		return &Error{
			Code:    500,
			Message: data,
		}
	}
	return &terr
}
