package yandex

import (
	"fmt"

	"github.com/illfalcon/spotiyan/pkg/httperrors"
)

type NoResultsFromYandex struct {
	trackID int64
	error   error
}

func NewNoResultsFromYandex(trackID int64) *NoResultsFromYandex {
	return &NoResultsFromYandex{trackID: trackID, error: httperrors.ErrNotFound}
}

func (e *NoResultsFromYandex) Error() string {
	return fmt.Sprintf("no results returned for track with id %v", e.trackID)
}

func (e *NoResultsFromYandex) Unwrap() error {
	return e.error
}
