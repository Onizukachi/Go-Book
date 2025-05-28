package main

import (
	"fmt"
	"image"
	color "image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var colorGreen = color.RGBA{G: 222, A: 255}
var palette = []color.Color{color.Black, colorGreen} // Инстанцирование среза
var mu sync.Mutex
var count int

const (
	blackIndex = 1
	whiteIndex = 0
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cyclesStr := r.URL.Query().Get("cycles")

		fmt.Printf("%T -> %v\n", cyclesStr, cyclesStr)

		cycles, err := strconv.Atoi(cyclesStr)
		if err != nil {
			cycles = 5
		}
		lissajous(w, cycles)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // Угловое разрешение
		size    = 100   // Размер канвы изображения [-size..+size]
		nframes = 64    // Количество фреймов
		delay   = 8     // Задержка между кадрами (в *10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0        // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes} // Инстанцирование структуры
	phase := 0.0                        // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}

}
