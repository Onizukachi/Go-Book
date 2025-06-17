package main

import (
	"fmt"
)

func main() {
	data := []string{"sex", "", "aaa", "", "qqq"}
	fmt.Printf("%q\n", nonemptyv2(data))
	fmt.Printf("%q\n", data)

	stack := []int{5, 7, 8, 9, 10}
	fmt.Println(stack[len(stack)-1])       // last element
	fmt.Println(stack[:len(stack)-1])      // removel last element
	fmt.Println(removeFromStack(stack, 0)) // remove from middle
}

func removeFromStack(stack []int, index int) []int {
	copy(stack[index:], stack[index+1:])

	return stack[:len(stack)-1]
}

func nonempty(strings []string) []string {
	i := 0

	for _, el := range strings {
		if el != "" {
			strings[i] = el
			i++
		}
	}

	return strings
}

func nonemptyv2(strings []string) []string {
	baseSlice := strings[:0]

	for _, el := range strings {
		if el != "" {
			baseSlice = append(baseSlice, el)
		}
	}

	return baseSlice
}
