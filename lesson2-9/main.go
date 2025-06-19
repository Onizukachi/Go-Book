package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during open file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := map[string]int{}

	for scanner.Scan() {
		wordCount[scanner.Text()]++
	}

	for k, v := range wordCount {
		fmt.Printf("Word: %s, Count: %d\n", k, v)
	}
}
