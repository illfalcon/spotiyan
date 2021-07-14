package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/illfalcon/spotiyan/internal/spoti"
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

	res, err := yaClient.GetTrackInfo(50290303)
	if err != nil {
		log.Fatal(err)
	}

	track, err := spotiClient.SearchForTrack(fmt.Sprintf("%v %v %v", res.Title, res.Artists, res.Albums))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(track.SimpleTrack.ExternalURLs["spotify"])
}
