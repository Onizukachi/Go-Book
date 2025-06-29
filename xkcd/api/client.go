package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	Title       string
	URL         string
	Description string `json:"transcript"`
}

const XKCD_URL = "https://xkcd.com/"

func GetComics() *[]Book {
	var books []Book

	for i := 1; i < 10; i++ {
		url := XKCD_URL + fmt.Sprintf("%d/info.0.json", i)
		response, err := http.Get(url)
		if err != nil || response.StatusCode != http.StatusOK {
			fmt.Printf("Failed to fetch book with id %d: %v\n", i, err)
			continue
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Failed to read body for id %d: %v\n", i, err)
			continue
		}

		var book Book

		err = json.Unmarshal(body, &book)
		if err != nil {
			fmt.Printf("Failed to unmarshal JSON for id %d: %v\n", i, err)
			continue
		}

		book.URL = url
		books = append(books, book)

	}

	return &books
}
