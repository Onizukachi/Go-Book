package main

import (
	"fmt"
)

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// reverseArr(&arr)
	// fmt.Println(arr)

	slice := []string{"s", "s", "s", "a", "a", "a", "a", "b", "s", "b", "p", "o", "o"} // s a b s b p o
	fmt.Println(removeDup(slice))

}

func reverseArr(arr *[5]int) *[5]int {
	for first, last := 0, len(*arr)-1; first < last; first, last = first+1, last-1 {
		arr[first], arr[last] = arr[last], arr[first]
	}

	return arr
}

func reverseSlice(slice []int) []int {
	for first, last := 0, len(slice)-1; first < last; first, last = first+1, last-1 {
		slice[first], slice[last] = slice[last], slice[first]
	}

	return slice
}

func removeDup(s []string) []string {
	if len(s) == 0 {
		return s
	}

	writeIdx := 1

	for i := 1; i < len(s); i++ {
		if s[i] != s[writeIdx-1] {
			s[writeIdx] = s[i]
			writeIdx++
		}
	}

	return s[:writeIdx]
}
