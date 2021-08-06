package response

import "Toolkit-GO/constants"

type Error struct {
	Code  constants.ErrorCodeType   `json:"Code"`
	Msg   string                    `json:"Msg"`
}

func (e *Error) WithCode(code constants.ErrorCodeType) *Error {
	e.Code = code
	return e
}

func (e *Error) WithMsg(msg string) *Error {
	e.Msg = msg
	return e
}

func NewErrorDefault(code constants.ErrorCodeType,msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func NewErrorWithLevel(code constants.ErrorCodeType, desc, msg string, level constants.ErrorLevelType) *Error {
	return &Error{
		Code:  code,
		Msg:   msg,
	}
}

func NewErrorWithType(code constants.ErrorCodeType, desc, msg string, typ constants.ApplicationType) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
func NewError(code constants.ErrorCodeType, msg string) *Error {
	return &Error{
		Code:  code,
		Msg:   msg,
	}
}

func NewErrorWithErrorCode(code constants.ErrorCodeType) *Error {
	return GetErrorWithErrorCode(code)
}
