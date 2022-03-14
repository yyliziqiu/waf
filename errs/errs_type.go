package errs

import "fmt"

/**
Error
*/
type Error struct {
	StatusCode int
	Code       int
	Message    string
}

/**
Error()
*/
func (e *Error) Error() string {
	return fmt.Sprintf("status: %d, code: %d, message: %s", e.StatusCode, e.Code, e.Message)
}

/**
自定义错误信息
*/
func (e Error) WithMessage(message string) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    message,
	}
}

/**
自定义错误信息
*/
func (e Error) WithMessageF(message string, a ...interface{}) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    fmt.Sprintf(message, a...),
	}
}
