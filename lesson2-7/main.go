package main

import (
	"fmt"
	"unicode"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverseArr(&arr)
	fmt.Println(arr)

	slice := []string{"s", "s", "s", "a", "a", "a", "a", "b", "s", "b", "p", "o", "o", "o", "o", "a"} // s a b s b p o
	fmt.Println(removeDup(slice))
}

func reverseSlice(slice []byte) []byte {
	for first, last := 0, len(slice)-1; first < last; first, last = first+1, last-1 {
		slice[first], slice[last] = slice[last], slice[first]
	}

	return slice
}

func mergeSpaces(slice []byte) []byte {
	correctIndex := 0

	for i, el := range slice {
		if i == 0 {
			continue
		}

		if unicode.IsSpace(rune(el)) && unicode.IsSpace(rune(slice[correctIndex])) {
			continue
		} else {
			correctIndex++
			slice[correctIndex] = el
		}
	}

	return slice[:correctIndex+1]
}

func reverseArr(arr *[5]int) *[5]int {
	for first, last := 0, len(*arr)-1; first < last; first, last = first+1, last-1 {
		arr[first], arr[last] = arr[last], arr[first]
	}

	return arr
}

func removeDup(s []string) []string {
	correctIndex := 0

	for i := 1; i < len(s); i++ {
		if s[correctIndex] != s[i] {
			correctIndex++
			s[correctIndex] = s[i]
		}
	}

	return s[:correctIndex+1]
}
