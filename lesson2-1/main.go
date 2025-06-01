package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	for {
		fmt.Println("Введите число hex:")

		var input string
		fmt.Scanln(&input)

		if input == "stop" {
			break
		}

		hexNum := formatHexPrefix(input)
		result := new(big.Int)
		_, ok := result.SetString(hexNum, 16)

		if !ok {
			fmt.Println("Ввведено не верное hex число")
			continue
		}

		fmt.Println("Результат ", result)

	}
}

func formatHexPrefix(s string) string {
	return strings.TrimPrefix(strings.ToLower(s), "0x")
}
