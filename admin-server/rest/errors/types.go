package errors

import "encoding/json"

type ErrCode int

const (
	ERR_VALIDATOR_NOT_FOUND ErrCode = 10001
)

type Error struct {
	Code int
	Msg  interface{}
}

func NewError(code int, msg interface{}) *Error {
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	bin, _ := json.Marshal(e)
	return string(bin)
}
