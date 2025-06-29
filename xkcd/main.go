package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"xkcd/api"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Incorrect Input")
	}

	books := api.GetComics()

	var result []api.Book
	for _, book := range *books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(strings.TrimSpace(input))) {
			result = append(result, book)
		}
	}

	if len(result) == 0 {
		log.Fatalf("Not Found books for %s", input)
	}

	for i, book := range result {
		fmt.Printf("%d. Title: %s; URL: %s\n", i+1, book.Title, book.URL)
	}
}
