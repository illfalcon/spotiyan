package spoti

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	tokenURL          = "https://accounts.spotify.com/api/token"
	spotiClientID     = "SPOTI_CLIENT_ID"
	spotiClientSecret = "SPOTI_CLIENT_SECRET"
)

type Client struct {
	encodedCredentials string
	httpClient         *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{
		httpClient:         client,
		encodedCredentials: encodeCredentials(getCredentials()),
	}
}

func (c *Client) GetToken() (string, error) {
	request, err := c.createRequestToGetToken()
	if err != nil {
		return "", err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", NewApiError(err)
	}

	authToken, err := decodeResponse(response)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

func decodeResponse(response *http.Response) (string, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var tokenResp AccessTokenResponse

	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}

func (c *Client) createRequestToGetToken() (*http.Request, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.encodedCredentials))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request, nil
}

func getCredentials() (clientID, clientSecret string) {
	return os.Getenv(spotiClientID), os.Getenv(spotiClientSecret)
}

func encodeCredentials(clientID, clientSecret string) string {
	plain := fmt.Sprintf("%v:%v", clientID, clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(plain))

	return encoded
}
