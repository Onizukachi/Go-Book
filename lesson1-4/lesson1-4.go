package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue

			}
			result := countLines(f)
			f.Close()

			hasRepeated := false

			for _, count := range result {
				if count > 1 {
					hasRepeated = true
					break
				}
			}

			if hasRepeated {
				fmt.Println(arg)
			}
		}
	}
}

func countLines(f *os.File) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	return counts
}
