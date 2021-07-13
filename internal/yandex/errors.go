package yandex

import (
	"fmt"

	"github.com/illfalcon/spotiyan/pkg/httperrors"
)

type NoResults struct {
	error error
}

func NewNoResults() *NoResults {
	return &NoResults{error: httperrors.ErrNotFound}
}

func (e *NoResults) Error() string {
	return fmt.Sprintf("no results returned from yandex: %s", e.error.Error())
}

func (e *NoResults) Unwrap() error {
	return e.error
}

type ApiError struct {
	errorType error
	apiError  error
}

func NewApiError(apiError error) *ApiError {
	return &ApiError{apiError: apiError, errorType: httperrors.ErrDependency}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("error from yandex api: %v", e.apiError)
}

func (e *ApiError) Unwrap() error {
	return e.errorType
}
