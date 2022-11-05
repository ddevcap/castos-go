package castos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type PrivateSubscribersService service

type PrivateSubscriber struct {
	Id        int64     `json:"id"`
	PodcastId int64     `json:"podcast_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Uuid      string    `json:"uuid"`
	FeedUrl   string    `json:"feed_url"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ps *PrivateSubscriber) UnmarshalJSON(data []byte) error {
	type Alias PrivateSubscriber

	aux := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias: (*Alias)(ps),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	ps.CreatedAt, _ = time.Parse(DateFormat, aux.CreatedAt)
	ps.UpdatedAt, _ = time.Parse(DateFormat, aux.UpdatedAt)

	return nil
}

func (service *PrivateSubscribersService) GetAll(podcastId int64) ([]*PrivateSubscriber, error) {
	path := "/private-subscribers"

	q := url.Values{}
	q.Set("podcast_id", fmt.Sprintf("%d", podcastId))

	req, err := service.c.newRequest(http.MethodGet, path, q, nil)
	if err != nil {
		return nil, err
	}

	privateSubscriber := make([]*PrivateSubscriber, 0)

	err = service.c.do(req, &privateSubscriber)
	if err != nil {
		return nil, err
	}

	return privateSubscriber, nil
}
