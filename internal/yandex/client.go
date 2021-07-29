package yandex

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	host   = "api.music.yandex.net"
	port   = 443
	scheme = "https"

	yandexShareURL = "https://music.yandex.ru/track/%d"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

func (c *Client) GetTrackInfo(trackID int64) (Track, error) {
	req, err := createRequestToGetTrackInfo(trackID)
	if err != nil {
		return Track{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Track{}, NewApiError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Track{}, fmt.Errorf("error reading body: %w", err)
	}
	defer resp.Body.Close()
	var response GetTrackResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Track{}, fmt.Errorf("error when unmarshalling rep body: %w", err)
	}

	return response.ToTrack()
}

func createRequestToGetTrackInfo(trackID int64) (*http.Request, error) {
	return http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s://%s:%d/tracks/%d", scheme, host, port, trackID),
		nil,
	)
}

func GetTrackIDFromURL(urlString string) (int64, error) {
	pieces := strings.Split(urlString, "/")
	if len(pieces) == 0 {
		return 0, NewBadRequest("invalid urlString for yandex music")
	}

	trackIDStr := pieces[len(pieces)-1]
	trackID, err := strconv.ParseInt(trackIDStr, 10, 64)
	if err != nil {
		return 0, NewBadRequest("invalid urlString for yandex music")
	}

	return trackID, nil
}

func (c *Client) SearchTrackURL(searchQuery string) (string, error) {
	req, err := createRequestToSearchTrack(searchQuery)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", NewApiError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body: %w", err)
	}
	defer resp.Body.Close()
	var response SearchTrackResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("error when unmarshalling rep body: %w", err)
	}

	if err := checkResponse(response); err != nil {
		return "", err
	}

	return fmt.Sprintf(yandexShareURL, response.Result.Tracks.Results[0].ID), nil
}

func checkResponse(response SearchTrackResponse) error {
	if len(response.Result.Tracks.Results) == 0 {
		return NewNoResults()
	}

	return nil
}

func createRequestToSearchTrack(searchQuery string) (*http.Request, error) {
	vals := url.Values{}
	vals.Set("type", "track")
	vals.Set("page", "0")
	vals.Set("text", searchQuery)

	return http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s://%s:%d/search?%s", scheme, host, port, vals.Encode()),
		nil,
	)
}
