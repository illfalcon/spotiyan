package translator

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/illfalcon/spotiyan/internal/spoti"
	"github.com/illfalcon/spotiyan/internal/yandex"
)

type Server struct {
	spotiClient *spoti.Client
	yaClient    *yandex.Client
}

func NewServer(spotiClient *spoti.Client, yaClient *yandex.Client) *Server {
	return &Server{spotiClient: spotiClient, yaClient: yaClient}
}

func (s *Server) HandleTranslate(w http.ResponseWriter, r *http.Request) {
	yandexTrackID := chi.URLParam(r, "yandexTrackID")
	yandexTrackIDInt, err := strconv.ParseInt(yandexTrackID, 10, 64)
	if err != nil {
		log.Printf("incorrect track id: %v", yandexTrackID)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect yandex track id"))

		return
	}

	yandexTrack, err := s.yaClient.GetTrackInfo(yandexTrackIDInt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("err"))

		return
	}

	spotifyTrack, err := s.spotiClient.SearchForTrack(yandexTrack.String())
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("err"))

		return
	}

	shareURL, err := spoti.GetShareURL(&spotifyTrack)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("err"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shareURL))
}
