package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/illfalcon/spotiyan/internal/spoti"
	"github.com/illfalcon/spotiyan/internal/telegram"
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

	tgBot := telegram.NewBot(service)
	err = tgBot.Init()
	if err != nil {
		log.Fatal(err)
	}

	go herokuHack()

	_ = tgBot.Listen()
}

func herokuHack() {
	log.Print("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		name := os.Getenv("NAME")
		if name == "" {
			name = "World"
		}
		_, _ = fmt.Fprintf(w, "Hello %s!\n", name)
	})

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
