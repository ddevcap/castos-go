package castos

import (
	"fmt"
	"log"
)

func ExampleNewClient() {
	const token = "token"

	client := NewClient(token)

	podcasts, err := client.Podcasts.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, podcast := range podcasts {
		fmt.Println(podcast.Title)
	}
}
