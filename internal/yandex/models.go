package yandex

import (
	"fmt"
	"strings"
)

type Response struct {
	Result []Result `json:"result"`
}

type Result struct {
	Title   string   `json:"title"`
	Artists []Artist `json:"artists"`
	Albums  []Album  `json:"albums"`
}

type Artist struct {
	Name string `json:"name"`
}

type Album struct {
	Title string `json:"title"`
}

type Track struct {
	Title   string
	Artists string
	Albums  string
}

func (r Response) ToTrack() (Track, error) {
	if len(r.Result) == 0 {
		return Track{}, NewNoResults()
	}

	result := r.Result[0]

	return Track{
		Title:   result.Title,
		Artists: concatArtists(result.Artists),
		Albums:  concatAlbums(result.Albums),
	}, nil
}

func concatAlbums(albums []Album) string {
	var titles []string
	for _, a := range albums {
		titles = append(titles, a.Title)
	}

	return strings.Join(titles, " ")
}

func concatArtists(artists []Artist) string {
	var titles []string
	for _, a := range artists {
		titles = append(titles, a.Name)
	}

	return strings.Join(titles, " ")
}

func (t Track) String() string {
	return fmt.Sprintf("%v %v %v", t.Title, t.Albums, t.Artists)
}
