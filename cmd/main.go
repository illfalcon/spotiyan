package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	service := translator.NewService(spotiClient, yaClient)
	handler := translator.NewHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Minute))

	r.Get("/translate/{yandexTrackID}", handler.HandleTranslate)

	srv := &http.Server{Addr: ":3000", Handler: r}
	stopAppCh := make(chan struct{})
	sigquit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigquit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigquit
		log.Printf("captured signal: %v\n", s)
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("could not shutdown service: %s", err)
		}
		log.Printf("shut down gracefully")
		stopAppCh <- struct{}{}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-stopAppCh
}
