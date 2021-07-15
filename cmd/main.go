package main

import (
	"log"
	"net/http"
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

	_ = tgBot.Listen()
}
