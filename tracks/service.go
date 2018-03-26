package tracks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Service interface {
		TopTracks(token string) (*[]Track, error)
	}

	Artist struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	Album struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	Track struct {
		Id      string   `json:"id"`
		Name    string   `json:"name"`
		Artists []Artist `json:"artists"`
		Album   Album    `json:"album"`
	}

	TopTracksResponse struct {
		Tracks []Track `json:"items"`
	}

	service struct {
		httpClient *http.Client
	}
)

const (
	timeRange = "medium_term"
)

func NewService(httpClient *http.Client) Service {
	return &service{
		httpClient: httpClient,
	}
}

func (s *service) TopTracks(token string) (*[]Track, error) {
	var topTracks TopTracksResponse

	topTracksURL := fmt.Sprintf("https://api.spotify.com/v1/me/top/tracks?time_range=%s", timeRange)
	req, err := http.NewRequest("GET", topTracksURL, nil)
	if err != nil {
		return &topTracks.Tracks, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := s.httpClient.Do(req)
	if err != nil {
		return &topTracks.Tracks, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&topTracks)
	return &topTracks.Tracks, err
}
