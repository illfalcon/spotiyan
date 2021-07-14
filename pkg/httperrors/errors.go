package httperrors

import (
	"errors"
	"fmt"
	"io"
)

const (
	NotFoundCode = iota
	DependencyErrorCode
	BadRequest
	InternalError
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
	ErrInternal   = NewHTTPError(InternalError, "internal error")
)

func WriteError(err error, writer io.Writer) {
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		_, _ = writer.Write([]byte(httpErr.Error()))

		return
	}

	_, _ = writer.Write([]byte(ErrInternal.Error()))
}

func WriteErrorAsString(err error) string {
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return httpErr.Error()
	}

	return ErrInternal.Error()
}
