package castos

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	BaseUrl    = "https://app.castos.com/api/v2"
	DateFormat = "2006-01-02 15:04:05"
)

type Client struct {
	Token              string
	BaseUrl            string
	QueryParams        map[string]string
	Headers            map[string]string
	http               *http.Client
	commonService      service
	Podcasts           *PodcastsService
	Episodes           *EpisodesService
	PrivateSubscribers *PrivateSubscribersService
	Categories         *CategoriesService
}

type service struct {
	c *Client
}

type authTransport struct {
	token            string
	defaultTransport http.RoundTripper
}

type apiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	q := req.URL.Query()
	q.Set("token", t.token)

	req.Header.Set("Accept", "application/json")
	req.URL.RawQuery = q.Encode()

	return t.defaultTransport.RoundTrip(req)
}

func NewClient(token string) *Client {
	c := &Client{
		BaseUrl: BaseUrl,
		http: &http.Client{
			Transport: &authTransport{
				token:            token,
				defaultTransport: http.DefaultTransport,
			},
		},
	}

	c.commonService.c = c
	c.Podcasts = (*PodcastsService)(&c.commonService)
	c.Episodes = (*EpisodesService)(&c.commonService)
	c.PrivateSubscribers = (*PrivateSubscribersService)(&c.commonService)
	c.Categories = (*CategoriesService)(&c.commonService)

	return c
}

func (c *Client) newRequest(method, path string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(BaseUrl + path)
	if err != nil {
		return nil, err
	}

	for key, value := range c.QueryParams {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		return c.handleError(res)
	}

	if v != nil {
		defer res.Body.Close()
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var responseData = apiResponse{
		Data: v,
	}

	var payload map[string]interface{}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		return err
	}

	if _, hasData := payload["data"]; hasData {
		return json.Unmarshal(b, &responseData)
	}

	return json.Unmarshal(b, &v)
}

func (c *Client) handleError(res *http.Response) error {
	var e ErrorResponse

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	return e
}
