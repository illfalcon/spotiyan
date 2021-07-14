package translator

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("incorrect yandex track id"))

		return
	}

	shareURL, err := h.translator.Translate(yandexTrackIDInt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("err"))

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(shareURL))
}
