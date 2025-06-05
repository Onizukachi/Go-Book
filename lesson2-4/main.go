package main

import (
	"fmt"
)

func main() {
	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("a\\b\t")

	for _, n := range nums {
		fmt.Printf("%8d", n)
	}

	fmt.Printf("\n\n")

	for _, outer := range nums {
		fmt.Printf("%d\t", outer)
		for _, inner := range nums {
			fmt.Printf("%8d", inner*outer)
		}

		fmt.Println()
	}
}
