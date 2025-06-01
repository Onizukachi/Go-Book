package cipherer

import (
	"encoding/base64"
	"errors"
	"fmt"
)

func Cipher(rawStr, key string) (string, error) {
	processResult := process([]byte(rawStr), []byte(key))

	fmt.Println(processResult)
	return base64.StdEncoding.EncodeToString(processResult), nil
}

func Decipher(base64Str, key string) (string, error) {
	decodedStr, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("не получилось раздекодировать строку")
	}

	processResult := process([]byte(decodedStr), []byte(key))

	return string(processResult), nil

}

func process(str, key []byte) []byte {
	for i, byte := range str {
		str[i] = byte ^ key[i%len(key)]
	}

	return str
}
