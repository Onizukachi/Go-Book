package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func concatArgs() string {
	var res, sep string
	for _, arg := range os.Args[1:] {
		res += arg + sep
		sep = " "
	}
	return res
}

func main() {

	// Тест strings.Join
	start := time.Now()
	result1 := strings.Join(os.Args[1:], " ")
	elapsed1 := time.Since(start)

	// Тест конкатенации
	start = time.Now()
	result2 := concatArgs()
	elapsed2 := time.Since(start)

	// Вывод результатов (после всех замеров)
	fmt.Println("Результат Join:", result1)
	fmt.Printf("Join занял: %.8fs\n", elapsed1.Seconds())

	fmt.Println("Результат конкатенации:", result2)
	fmt.Printf("Конкатенация заняла: %.8fs\n", elapsed2.Seconds())

	// Сравнение производительности
	fmt.Printf("\nКонкатенация медленнее в %.2f раз\n",
		float64(elapsed2.Nanoseconds())/float64(elapsed1.Nanoseconds()))
}
