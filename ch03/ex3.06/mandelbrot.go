// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
)

var stdout io.Writer = os.Stdout

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			draw(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	draw(stdout)
}

func draw(out io.Writer) {
	const (
		width, height = 1024, 1024
	)

	img := drawSuperSampledImg(width, height)
	png.Encode(out, img) // NOTE: ignoring errors
}

func drawSuperSampledImg(w, h int) image.Image {
	return resampleQuad(drawQuadrupleImg(w, h), w, h)
}

func resampleQuad(srcImg image.Image, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, average(px*2, py*2, srcImg))
		}
	}
	return img
}

func average(x, y int, s image.Image) color.Color {
	return average9(x, y, s)
}

func average4(x, y int, s image.Image) color.Color {
	var r, g, b uint32
	filters := []struct{ dx, dy int }{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	for _, f := range filters {
		sr, sg, sb, _ := s.At(x+f.dx, y+f.dy).RGBA()
		r += sr
		g += sg
		b += sb
	}
	return color.RGBA{uint8(r / 4), uint8(g / 4), uint8(b / 4), 255}
}

func average9(x, y int, s image.Image) color.Color {
	var r, g, b uint32

	// filter weight on pos:
	// 1 2 1
	// 2 4 2
	// 1 2 1
	filters := []struct {
		dx, dy int
		w      uint32
	}{
		{-1, -1, 1}, {0, -1, 2}, {1, -1, 1},
		{-1, 0, 2}, {0, 0, 4}, {1, 0, 2},
		{-1, 1, 1}, {0, 1, 2}, {1, 1, 1},
	}
	for _, f := range filters {
		sr, sg, sb, _ := s.At(x+f.dx+1, y+f.dy+1).RGBA()
		r += sr * f.w
		g += sg * f.w
		b += sb * f.w
	}
	return color.RGBA{uint8(r / 9 >> 8), uint8(g / 9 >> 8), uint8(b / 9 >> 8), 255}
}

func drawQuadrupleImg(w, h int) image.Image {
	return drawMandelbrot(w*2, h*2)
}

func drawMandelbrot(width, height int) image.Image {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
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
