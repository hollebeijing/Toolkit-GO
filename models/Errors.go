package models

import "errors"

var (
	// ErrInvalidParameter 参数错误.
	ErrInvalidParameter = errors.New("Invalid parameter")
)


type Error struct {
	code    int
	message string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func NewError(code int, message string) Error {
	return Error{code: code, message: message}
}

