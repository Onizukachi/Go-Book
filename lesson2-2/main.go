package main

import (
	"bufio"
	"flag"
	"fmt"
	"lesson2-2/cipherer"
	"os"
	"strings"
)

var mode = flag.String("mode", "cipher", "Режим работы может быть cipher или decipher")
var key = flag.String("key", "", "Ключ для шифрования")

func main() {
	flag.Parse()

	if len(*key) == 0 {
		fmt.Println("Ключ шифрования не был указан")
		os.Exit(1)
	}

	input := getUserInput()

	switch *mode {
	case "cipher":
		cipheredText, err := cipherer.Cipher(input, *key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Произошло ошибка при шифровании: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipheredText)
	case "decipher":
		decipheredText, err := cipherer.Decipher(input, *key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Произошло ошибка при дешифровании: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(decipheredText)
	default:
		fmt.Println("Не поддерживыемый режим")
	}
}

func getUserInput() string {
	for {
		fmt.Println("Введите строку для шифрования:")
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')

		if err != nil || len(strings.TrimSpace(str)) == 0 {
			continue
		}

		return strings.TrimSpace(str)
	}
}
