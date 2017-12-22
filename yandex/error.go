package yandex

import "fmt"

type Code int

const (
	OK                       Code  = iota
	KEY_INVALID
	KEY_BLOCKED
	DAILY_REQ_LIMIT_EXCEEDED
	TEXT_TOO_LONG
	LANG_NOT_SUPPORTED
	UNEXPECTED_ERROR
)

type ExternalError struct {
	Err  error
	Code Code
}

func (e *ExternalError) String() string {
	return fmt.Sprintf("%v, with code=%v", e.Err.Error(), e.Code)
}
