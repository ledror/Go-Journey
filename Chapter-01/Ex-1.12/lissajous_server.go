package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
	blueIndex  = 3
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		cycles  int
		res     float64
		size    int
		nframes int
		delay   int
		err     error
	)
	params := r.URL.Query()
	if params.Has("cycles") {
		cycles, err = strconv.Atoi(params["cycles"][0])
		if err != nil {
			fmt.Fprintf(w, "Invalid Parameter: cycles=%s\n", params["cycles"][0])
			return
		}
	} else {
		cycles = 5
	}
	if params.Has("res") {
		res, err = strconv.ParseFloat(params["res"][0], 64)
		if err != nil {
			fmt.Fprintf(w, "Invalid Parameter: res=%s\n", params["res"][0])
			return
		}
	} else {
		res = 0.001
	}
	if params.Has("size") {
		size, err = strconv.Atoi(params["size"][0])
		if err != nil {
			fmt.Fprintf(w, "Invalid Parameter: size=%s\n", params["size"][0])
			return
		}
	} else {
		size = 100
	}
	if params.Has("nframes") {
		nframes, err = strconv.Atoi(params["nframes"][0])
		if err != nil {
			fmt.Fprintf(w, "Invalid Parameter: nframes=%s\n", params["nframes"][0])
			return
		}
	} else {
		nframes = 64
	}
	if params.Has("delay") {
		delay, err = strconv.Atoi(params["delay"][0])
		if err != nil {
			fmt.Fprintf(w, "Invalid Parameter: delay=%s\n", params["delay"][0])
			return
		}
	} else {
		delay = 8
	}
	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(1 + rand.Intn(len(palette)-1))
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				colorIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // ignoring encoding errors
}
