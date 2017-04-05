// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var out io.Writer = os.Stdout

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, //BLACK
	color.RGBA{0xff, 0x00, 0x00, 0xff}, //RED
	color.RGBA{0x00, 0xff, 0x00, 0xff}, //GREEN
	color.RGBA{0x00, 0x00, 0xff, 0xff}, //BLUE
	color.RGBA{0xff, 0xff, 0x00, 0xff}, //YELLOW
	color.RGBA{0xff, 0x00, 0xff, 0xff}, //MAGENTA
	color.RGBA{0x00, 0xff, 0xff, 0xff}, //CYAN
	color.RGBA{0xff, 0xff, 0xff, 0xff}, //WHITE
}

const (
	blackIndex = uint8(iota) // first color in palette
	redIndex
	greenIndex
	blueIndex
	yellowIndex
	magentaIndex
	cyanIndex
	whiteIndex
	indexSize
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(out)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	color := uint8(rand.Intn(int(indexSize-1))) + 1 // avoid black on black
	freq := rand.Float64() * 3.0                    // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), color)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
