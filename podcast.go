package castos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type PodcastsService service

type Podcast struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RssUrl      string    `json:"rss_url"`
	Subdomain   string    `json:"subdomain"`
	Author      string    `json:"author"`
	Copyright   string    `json:"copyright"`
	Image       string    `json:"image"`
	Language    string    `json:"language"`
	Categories  []string  `json:"categories"`
	Website     string    `json:"website"`
	Itunes      string    `json:"itunes"`
	GooglePLay  string    `json:"google_play"`
	Stitcher    string    `json:"stitcher"`
	Spotify     string    `json:"spotify"`
	Explicit    bool      `json:"explicit"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Podcast) UnmarshalJSON(data []byte) error {
	type Alias Podcast

	aux := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Explicit  int    `json:"explicit"`
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	p.CreatedAt, _ = time.Parse(DateFormat, aux.CreatedAt)
	p.UpdatedAt, _ = time.Parse(DateFormat, aux.UpdatedAt)

	p.Explicit = aux.Explicit == 1

	return nil
}

func (service *PodcastsService) GetAll() ([]*Podcast, error) {
	path := fmt.Sprintf("/podcasts")

	req, err := service.c.newRequest(http.MethodGet, path, url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	podcastList := map[string]map[int64]string{}

	err = service.c.do(req, &podcastList)
	if err != nil {
		return nil, err
	}

	if _, exists := podcastList["podcast_list"]; !exists {
		return nil, errors.New("no podcast list found in response data")
	}

	podcasts := make([]*Podcast, 0)

	for id, title := range podcastList["podcast_list"] {
		podcasts = append(podcasts, &Podcast{
			Id:    id,
			Title: title,
		})
	}

	return podcasts, nil
}

func (service *PodcastsService) Get(id int64) (*Podcast, error) {
	path := fmt.Sprintf("/podcasts/%d", id)

	req, err := service.c.newRequest(http.MethodGet, path, url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	var podcast Podcast

	err = service.c.do(req, &podcast)
	if err != nil {
		return nil, err
	}

	return &podcast, nil
}
