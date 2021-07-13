package main

import (
	"fmt"
	"net/http"

	"github.com/illfalcon/spotiyan/internal/yandex"
)

func main() {
	client := &http.Client{}
	yaClient := yandex.NewClient(client)
	track, _ := yaClient.GetTrackInfo(85534379)
	fmt.Println(track)
}
