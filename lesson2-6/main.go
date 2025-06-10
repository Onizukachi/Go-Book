package main

import (
	"fmt"
)

func main() {
	orig := make([]int, 5, 6) // len=5, cap=6, все элементы = 0
	copy(orig, []int{1, 2, 3, 4, 5})
	newArr := appendInt(orig, 10)

	fmt.Println(newArr)
}

func appendInt(slice []int, target int) []int {
	if len(slice)+1 <= cap(slice) {
		newSlice := slice[:len(slice)+1]

		newSlice[len(slice)] = target
		return newSlice
	} else {
		newSlice := make([]int, len(slice)+1, (len(slice)+1)*2)

		copy(newSlice, slice)
		newSlice[len(slice)] = target
		return newSlice
	}
}
