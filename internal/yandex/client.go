package yandex

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	host   = "api.music.yandex.net"
	port   = 443
	scheme = "https"
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
	var response Response
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

func GetTrackIDFromURL(url string) (int64, error) {
	pieces := strings.Split(url, "/")
	if len(pieces) == 0 {
		return 0, NewBadRequest("invalid url for yandex music")
	}

	trackIDStr := pieces[len(pieces)-1]
	trackID, err := strconv.ParseInt(trackIDStr, 10, 64)
	if err != nil {
		return 0, NewBadRequest("invalid url for yandex music")
	}

	return trackID, nil
}
