package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/illfalcon/spotiyan/internal/spoti"
)

func main() {
	_ = godotenv.Load()
	client := spoti.NewClient(&http.Client{})
	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.SearchForTrack("Running Away VANO 3000 BADBADNOTGOOD Samuel T. Herring Running Away")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res.SimpleTrack.ExternalURLs)
}
