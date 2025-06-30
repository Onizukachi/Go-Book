package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	elementCounts := make(map[string]int)

	countElements(doc, elementCounts)

	fmt.Println("Количество элементов по тегам:")
	for tag, count := range elementCounts {
		fmt.Printf("%s: %d\n", tag, count)
	}
}

func countElements(node *html.Node, counts map[string]int) {
	if node == nil {
		return
	}

	if node.Type == html.ElementNode {
		counts[node.Data]++
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		countElements(child, counts)
	}
}
