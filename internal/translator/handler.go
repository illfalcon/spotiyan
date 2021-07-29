package translator

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/illfalcon/spotiyan/pkg/httperrors"
)

type Handler struct {
	translator *Service
}

func NewHandler(translator *Service) *Handler {
	return &Handler{translator: translator}
}

func (h *Handler) HandleTranslate(w http.ResponseWriter, r *http.Request) {
	yandexTrackID := chi.URLParam(r, "yandexTrackID")
	yandexTrackIDInt, err := strconv.ParseInt(yandexTrackID, 10, 64)
	if err != nil {
		log.Printf("incorrect track id: %v", yandexTrackID)
		httperrors.WriteError(err, w)

		return
	}

	shareURL, err := h.translator.TranslateYandexToSpotify(yandexTrackIDInt)
	if err != nil {
		log.Print(err)
		httperrors.WriteError(err, w)

		return
	}

	_, _ = w.Write([]byte(shareURL))
}
