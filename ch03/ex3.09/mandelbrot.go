// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var stdout io.Writer = os.Stdout

const (
	mandelbrotImg = iota
	acosImg
	sqrtImg
	newtonImg
	typeSize
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			t, x, y, s := parseQuery(r.URL)
			draw(w, int(t)%typeSize, x, y, s)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	draw(stdout, 0, 2.0, 2.0, 1.0)
}

func parseQuery(u *url.URL) (t int64, x, y, s float64) {
	t, err := strconv.ParseInt(u.Query().Get("type"), 10, 32)
	if err != nil {
		t = 0
	}

	x, err = strconv.ParseFloat(u.Query().Get("x"), 64)
	if err != nil {
		x = 0
	}
	y, err = strconv.ParseFloat(u.Query().Get("y"), 64)
	if err != nil {
		y = 0
	}
	s, err = strconv.ParseFloat(u.Query().Get("scale"), 64)
	if err != nil {
		s = 1
	}
	return
}

func draw(out io.Writer, typ int, ix, iy, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	s := math.Abs(zoom)
	v := float64(2) / s
	if math.IsNaN(v) {
		v = 2
	}
	dx, dy := ix/(s*width), iy/(s*height)
	if math.IsNaN(dx) {
		dx = 0
	}
	if math.IsNaN(dy) {
		dy = 0
	}

	var xmin, ymin, xmax, ymax float64 = -v, -v, +v, +v

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x+dx, y+dy)
			// Image point (px, py) represents complex value z.
			switch typ {
			case mandelbrotImg:
				img.Set(px, py, mandelbrot(z))
			case acosImg:
				img.Set(px, py, acos(z))
			case sqrtImg:
				img.Set(px, py, sqrt(z))
			case newtonImg:
				img.Set(px, py, newton(z))
			}
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorize(v, z, n)
		}
	}
	return color.Black
}

func colorize(v, z complex128, n uint8) color.Color {
	const contrast = 15
	blue := 255 - contrast*n
	red := 255 - blue
	green := lerp(red, blue, n%1)

	return color.RGBA{red, green, blue, 255}
}

func lerp(v0, v1, t uint8) uint8 {
	return v0 + t*(v1-v0)
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
