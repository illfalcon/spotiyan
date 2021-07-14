package httperrors

import (
	"fmt"
)

const (
	NotFoundCode = iota
	DependencyErrorCode
	BadRequest
)

type HTTPError struct {
	code int
	msg  string
}

func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{code: code, msg: msg}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("code: %v, msg: %v", e.code, e.msg)
}

var (
	ErrNotFound   = NewHTTPError(NotFoundCode, "not found")
	ErrDependency = NewHTTPError(DependencyErrorCode, "dependency error")
	ErrBadRequest = NewHTTPError(BadRequest, "bad request")
)
