package spoti

import (
	"fmt"

	"github.com/illfalcon/spotiyan/pkg/httperrors"
)

type ApiError struct {
	errorType error
	apiError  error
}

func NewApiError(apiError error) *ApiError {
	return &ApiError{apiError: apiError, errorType: httperrors.ErrDependency}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("error from spotify api: %v", e.apiError)
}

func (e *ApiError) Unwrap() error {
	return e.errorType
}
