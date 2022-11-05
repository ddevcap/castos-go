package castos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type EpisodesService service

type Episode struct {
	Id               int64     `json:"id"`
	PostId           string    `json:"post_id"`
	PostTitle        string    `json:"post_title"`
	PostContent      string    `json:"post_content"`
	PostDate         string    `json:"post_date"`
	PodcastId        int64     `json:"podcast_id"`
	UserId           int64     `json:"user_id"`
	Keywords         string    `json:"keywords"`
	SeriesNumber     int64     `json:"series_number"`
	EpisodeNumber    int64     `json:"episode_number"`
	EpisodeImage     string    `json:"episode_image"`
	EpisodeType      string    `json:"episode_type"`
	Explicit         bool      `json:"explicit"`
	PostSlug         string    `json:"post_slug"`
	YouTubeId        string    `json:"youtube_id"`
	RepublishTrigger int64     `json:"republish_trigger"`
	WebsiteSync      int64     `json:"website_sync"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (e *Episode) UnmarshalJSON(data []byte) error {
	type Alias Episode

	aux := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Explicit  int    `json:"explicit"`
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	e.CreatedAt, _ = time.Parse(DateFormat, aux.CreatedAt)
	e.UpdatedAt, _ = time.Parse(DateFormat, aux.UpdatedAt)

	e.Explicit = aux.Explicit == 1

	return nil
}

func (service *EpisodesService) GetAll(podcastId int64) ([]*Episode, error) {
	path := fmt.Sprintf("/podcasts/%d/episodes", podcastId)

	req, err := service.c.newRequest(http.MethodGet, path, url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	episodes := make([]*Episode, 0)

	err = service.c.do(req, &episodes)
	if err != nil {
		return nil, err
	}

	return episodes, nil
}

func (service *EpisodesService) Get(podcastId, id int64) (*Episode, error) {
	path := fmt.Sprintf("/podcasts/%d/episodes/%d", podcastId, id)

	req, err := service.c.newRequest(http.MethodGet, path, url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	var episode Episode

	err = service.c.do(req, &episode)
	if err != nil {
		return nil, err
	}

	return &episode, nil
}
