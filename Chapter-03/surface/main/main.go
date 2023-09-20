package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 1366, 768
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.2
	angle         = math.Pi / 6
)

type zFunc func(x, y float64) float64

var (
	sin30, cos30 = math.Sin(angle), math.Cos(angle)
	f            zFunc
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	svg(os.Stdout)
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if params.Has("sinc") {
		f = sinc
	} else if params.Has("eggbox") {
		f = eggBox
	} else if params.Has("saddle") {
		f = saddle
	} else {
		f = sinc
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	svg(w)
}

func svg(w io.Writer) {
	zmin, zmax := minmax()
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) ||
				math.IsInf(cx, 0) ||
				math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) ||
				math.IsInf(dy, 0) {
				continue
			}
			fmt.Fprintf(w,
				"<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zmin, zmax),
				ax,
				ay,
				bx,
				by,
				cx,
				cy,
				dx,
				dy,
			)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func minmaxPoly(i, j int) (min float64, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for a := 0; a <= 1; a++ {
		for b := 0; b <= 1; b++ {
			x := xyrange * (float64(i+a)/cells - 0.5)
			y := xyrange * (float64(j+b)/cells - 0.5)
			z := f(x, y)
			min = math.Min(min, z)
			max = math.Max(max, z)
		}
	}
	return
}

func minmax() (min float64, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			polyMin, polyMax := minmaxPoly(i, j)
			min = math.Min(min, polyMin)
			max = math.Max(max, polyMax)
		}
	}
	return
}

func color(i, j int, min, max float64) string {
	minPoly, maxPoly := minmaxPoly(i, j)
	var color string
	var redIntensity float64
	var blueIntensity float64
	blueIntensity = math.Exp(min) / math.Exp(minPoly) * 255
	redIntensity = math.Exp(maxPoly) / math.Exp(max) * 255
	color = fmt.Sprintf("#%02x00%02x", int(redIntensity), int(blueIntensity))
	return color
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func sinc(x, y float64) float64 {
	r := math.Hypot(x, y)
	res := math.Sin(r) / r
	if math.IsNaN(res) {
		return 1.0
	}
	return res
}

func eggBox(x, y float64) float64 {
	const a = 0.5
	return a * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	const (
		a = 30.0
		b = 20.0
	)
	return (y*y)/(a*a) - (x*x)/(b*b)
}
