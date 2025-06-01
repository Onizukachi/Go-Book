package main

import (
	"errors"
	"fmt"
	"lesson1-9/distanceconv"
	"lesson1-9/tempconv"
	"os"
	"strconv"
)

func main() {
	input, err := getInput()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("вы ввели не число", err)
		os.Exit(1)
	}

	fmt.Printf("Число %g в метрах будет равно %s\n", number, distanceconv.FtToMetr(distanceconv.Ft(number)))
	fmt.Printf("Число %g в фунтах будет равно %s\n", number, distanceconv.MetrToFt(distanceconv.Metr(number)))

	fmt.Printf("Число %g в цельсиях будет равно %s\n", number, tempconv.FToC(tempconv.Fahrenheit(number)))
	fmt.Printf("Число %g в фарейнгейтах будет равно %s\n", number, tempconv.CToF(tempconv.Celsius(number)))
}

func getInput() (string, error) {
	input, err := getArgsInput()

	if err != nil {
		input, err = getPromptInput()

		if err != nil {
			return "", errors.New("ввод не получен")
		}
	}

	return input, nil
}

func getArgsInput() (string, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return "", errors.New("aргумент не предоставлен")
	}

	return args[0], nil
}

func getPromptInput() (string, error) {
	fmt.Println("Введите число для конвертаций!")

	input := new(string)
	_, err := fmt.Scan(input)

	if err != nil {
		return "", errors.New("вы не ничего ввели")
	}

	return *input, nil
}
