package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/illfalcon/spotiyan/internal/spoti"
	"github.com/illfalcon/spotiyan/internal/translator"
	"github.com/illfalcon/spotiyan/internal/yandex"
)

func main() {
	_ = godotenv.Load()

	httpClient := &http.Client{Timeout: time.Minute}
	spotiClient := spoti.NewClient(httpClient)
	err := spotiClient.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	yaClient := yandex.NewClient(httpClient)

	server := translator.NewServer(spotiClient, yaClient)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Minute))

	r.Get("/translate/{yandexTrackID}", server.HandleTranslate)

	http.ListenAndServe(":3000", r)
}
