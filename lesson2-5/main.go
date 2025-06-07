package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-12345321789.231231")) // -12,345,321,789.231231
	fmt.Println(comma("-12345321789"))        // -12,345,321,789
	fmt.Println(comma("12345321789.23"))      // 12,345,321,789.23
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
