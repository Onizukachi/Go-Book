package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearhResult struct {
	TotalCount int
	Items      []Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	Body      string
	CreatedAt time.Time `json:"created_at"`
	User      *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	url := IssuesURL + "?q=" + "decoder"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("Error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var result IssueSearhResult

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %v  %9.9s  %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
	}

	var lastWeek []Issue
	var lastLastWeek []Issue

	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Round(time.Hour) < 7*24*time.Hour {
			lastWeek = append(lastWeek, item)
		} else {
			lastLastWeek = append(lastLastWeek, item)
		}
	}

	fmt.Println("LAST WEKKK")
	for _, item := range lastWeek {
		fmt.Printf("#%-5d %v  %9.9s  %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
	}

	fmt.Println("LAST PREV WEKKK")
	for _, item := range lastLastWeek {
		fmt.Printf("#%-5d %v  %9.9s  %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
