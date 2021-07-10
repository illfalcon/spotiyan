package main

import (
	"encoding/json"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	host   = "api.music.yandex.net"
	port   = 443
	scheme = "https"
)

type YandexResponse struct {
	Result []YandexResult `json:"result"`
}

type YandexResult struct {
	Title   string         `json:"title"`
	Artists []YandexArtist `json:"artists"`
	Albums  []YandexAlbum  `json:"albums"`
}

type YandexArtist struct {
	Name string `json:"name"`
}

type YandexAlbum struct {
	Title string `json:"title"`
}

func main() {
	resp, err := http.Get(fmt.Sprintf("%s://%s:%d/tracks/86349653", scheme, host, port))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var yandexTrack YandexResponse
	err = json.Unmarshal(body, &yandexTrack)
	if err != nil {
		log.Fatal(err)
	}

	//data := url.Values{}
	//data.Set("grant_type", "client_credentials")
	//request, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//request.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("secrets"))))
	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//client := http.Client{}
	//accessToken, err := client.Do(request)
	spotiAuth := spotify.NewAuthenticator("")
	spotiClient := spotiAuth.NewClient(&oauth2.Token{AccessToken: "token"})
	res, err := spotiClient.Search("Aurora Cure For Me", spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	//body, err := ioutil.ReadAll(res)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(res.Tracks.Tracks[0].ExternalURLs["spotify"])
}
