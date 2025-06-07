package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	commaRun()
	isAnagramRun()
}

func isAnagramRun() int {
	fmt.Println(isAnagram("hello", "word"))  // false
	fmt.Println(isAnagram("Липа", "пила"))   // true
	fmt.Println(isAnagram("Актер", "терка")) // true
	fmt.Println(isAnagram("Сосна", "Насос")) // true

	return 0
}

func isAnagram(str1, str2 string) bool {
	counter := make(map[rune]int)
	for _, el := range strings.ToLower(str1) {
		counter[el]++
	}

	for _, el := range strings.ToLower(str2) {
		counter[el]--

		if counter[el] < 0 {
			return false
		}
	}

	return true
}

func commaRun() int {
	fmt.Println(comma("-12345321789.231231")) // -12,345,321,789.231231
	fmt.Println(comma("-12345321789"))        // -12,345,321,789
	fmt.Println(comma("12345321789.23"))      // 12,345,321,789.23

	return 0
}

func comma(s string) string {
	var buf bytes.Buffer

	signIndex := 0
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		signIndex = 1
	}

	dotIndex := strings.LastIndex(s, ".")
	if dotIndex == -1 {
		dotIndex = len(s) - 1
	}

	text := s[signIndex:dotIndex]
	length := len(text)

	for i := 0; i <= length-1; i++ {
		buf.WriteByte(text[i])

		if (length-i-1)%3 == 0 && i != length-1 {
			buf.WriteByte(',')
		}
	}

	buf.WriteString(s[dotIndex:])
	return buf.String()
}
