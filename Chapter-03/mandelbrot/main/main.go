package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"strconv"

	// "io"
	"math"
	"math/cmplx"
	"os"
)

const (
	width  = 1024
	height = 1024
	xmin   = -2
	ymin   = -2
	xmax   = +2
	ymax   = +2
	zoom   = 1
)

func main() {
	// if len(os.Args) == 2 && os.Args[1] == "super" {
	// 	png.Encode(os.Stdout, superSampling())
	// } else {
	// 	png.Encode(os.Stdout, normalSampling())
	// }
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	png.Encode(os.Stdout, normalSampling(xmin, ymin, xmax, ymax))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := map[string]float64{
		"xmin": xmin,
		"ymin": ymin,
		"xmax": xmax,
		"ymax": ymax,
		"zoom": zoom,
	}
	sample := normalSampling
	reqParams := r.URL.Query()
	for param := range params {
		if reqParams.Has(param) {
			val, err := strconv.ParseFloat(reqParams[param][0], 64)
			if err != nil {
				http.Error(
					w,
					fmt.Sprintf("Invalid Format of param %s: %s", param, err),
					http.StatusBadRequest,
				)
				return
			}
			params[param] = val
		}
	}
	if params["xmax"] <= params["xmin"] || params["ymax"] <= params["ymin"] || params["zoom"] <= 0 {
		http.Error(w, "Invalid Arguments", http.StatusBadRequest)
		return
	}
	if reqParams.Has("super") {
		sample = superSampling
	}
	xlen := params["xmax"] - params["xmin"]
	ylen := params["ymax"] - params["ymin"]
	xmid := params["xmin"] + xlen/2
	ymid := params["ymin"] + ylen/2
	params["xmin"] = xmid - xlen/2/params["zoom"]
	params["ymin"] = ymid - ylen/2/params["zoom"]
	params["xmax"] = xmid + xlen/2/params["zoom"]
	params["ymax"] = ymid + ylen/2/params["zoom"]
	png.Encode(w, sample(params["xmin"], params["ymin"], params["xmax"], params["ymax"]))
}

func normalSampling(xmin, ymin, xmax, ymax float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func superSampling(xmin, ymin, xmax, ymax float64) image.Image {
	img := normalSampling(xmin, ymin, xmax, ymax)
	superImg := image.NewRGBA(image.Rect(0, 0, width, height))
	// we can now super sample
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			pixels := make([]color.Color, 4)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					pixels[i+2*j] = img.At(px+i, py+j)
				}
			}
			superImg.Set(px, py, avg(pixels))
		}
	}
	return superImg
}

func avg(colors []color.Color) color.Color {
	var avg_r, avg_g, avg_b, avg_a uint16
	n := uint32(len(colors))
	for _, color := range colors {
		r, g, b, a := color.RGBA()
		avg_r += uint16(r / n)
		avg_g += uint16(g / n)
		avg_b += uint16(b / n)
		avg_a += uint16(a / n)
	}
	return color.RGBA64{avg_r, avg_g, avg_b, avg_a}
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 255
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n < 50:
				return color.RGBA{2 * n, 0, 0, 255}
			case n < 100:
				return color.RGBA{0, 2 * n, 0, 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
