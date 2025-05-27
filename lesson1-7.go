package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Запуск горутины
	}

	var result string

	for range os.Args[1:] {
		result += <-ch + "\n"
	}

	result += fmt.Sprintf("%.8fs elapsed\n", time.Since(start).Seconds())
	f, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	l, err := f.WriteString(result)
	fmt.Println(l, "bytes written successfully")
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка ошибки в канал `ch`
		return
	}
	nbytes, err := io.Copy(io.Discard, response.Body)
	_ = response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.8fs %7d %s", secs, nbytes, url)
}
