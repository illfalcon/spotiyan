package translator

import (
	"fmt"

	"github.com/illfalcon/spotiyan/internal/spoti"
	"github.com/illfalcon/spotiyan/internal/yandex"
)

type Service struct {
	spotiClient *spoti.Client
	yaClient    *yandex.Client
}

func NewService(spotiClient *spoti.Client, yaClient *yandex.Client) *Service {
	return &Service{spotiClient: spotiClient, yaClient: yaClient}
}

func (s *Service) TranslateYandexToSpotify(yandexTrackID int64) (string, error) {
	yandexTrack, err := s.yaClient.GetTrackInfo(yandexTrackID)
	if err != nil {
		return "", fmt.Errorf("error translating track with id %v: %w", yandexTrackID, err)
	}

	spotifyTrack, err := s.spotiClient.SearchForTrack(yandexTrack.String())
	if err != nil {
		return "", fmt.Errorf("error translating track with id %v: %w", yandexTrackID, err)
	}

	shareURL, err := spoti.GetShareURL(&spotifyTrack)
	if err != nil {
		return "", fmt.Errorf("error translating track with id %v: %w", yandexTrackID, err)
	}

	return shareURL, nil
}

func (s *Service) TranslateSpotifyToYandex(spotifyTrackID string) (string, error) {
	spotifyTrack, err := s.spotiClient.GetTrack(spotifyTrackID)
	if err != nil {
		return "", fmt.Errorf("error translating track with id %v: %w", spotifyTrackID, err)
	}

	yandexTrackURL, err := s.yaClient.SearchTrackURL(spoti.TrackToString(spotifyTrack))
	if err != nil {
		return "", fmt.Errorf("error translating track with id %v: %w", spotifyTrackID, err)
	}

	return yandexTrackURL, nil
}
