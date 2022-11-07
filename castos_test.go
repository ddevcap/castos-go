package castos

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClientEmptyToken(t *testing.T) {
	client, err := NewClient("")

	if client != nil || err == nil {
		t.Error()
	}
}

func TestClientAuthTransport(t *testing.T) {
	token := "token"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if !q.Has("token") {
			t.Error()
		}

		if q.Get("token") != token {
			t.Error()
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Error()
		}

		w.Write([]byte(`{}`))
	}))

	defer svr.Close()

	client := Client{
		baseUrl: svr.URL,
		http: &http.Client{
			Transport: &authTransport{
				token:            token,
				defaultTransport: http.DefaultTransport,
			},
		},
	}

	req, err := client.newRequest(http.MethodGet, "", url.Values{}, nil)
	if err != nil {
		t.Error()
	}

	err = client.do(req, nil)
	if err != nil {
		t.Error()
	}
}

func ExampleNewClient() {
	token := "token"

	client, err := NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	podcasts, err := client.Podcasts.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, podcast := range podcasts {
		fmt.Println(podcast.Title)
	}
}
