// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var stdout io.Writer = os.Stdout

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var typ = 0

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			typ = rand.Intn(int(4))
			w.Header().Set("Content-Type", "image/svg+xml")
			surface(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	surface(stdout)
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j)
			bx, by, berr := corner(i, j)
			cx, cy, cerr := corner(i, j+1)
			dx, dy, derr := corner(i+1, j+1)
			if aerr == nil &&
				berr == nil &&
				cerr == nil &&
				derr == nil {
				r, b := color(i, j)
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%d,0,%d)\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, b)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, error) {
	x, y, z := xyz(i, j)

	if math.IsNaN(z) {
		return 0, 0, fmt.Errorf("f(%g, %g) was NaN. corner(%d, %d)", x, y, i, j)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func color(i, j int) (int, int) {
	_, _, z := xyz(i, j)
	r, b := 0.0, 0.0
	if !math.IsNaN(z) {
		if z >= 0 {
			r = 255 * z
		} else {
			b = 255 * math.Abs(z)
		}
	}
	return int(r), int(b)
}

func xyz(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	return x, y, z
}

func f(x, y float64) float64 {
	switch typ {
	case 1:
		return feggbox(x, y)
	case 2:
		return fmoguls(x, y)
	case 3:
		return fsaddle(x, y)
	default:
		return fpersp(x, y)
		//return ftest(x, y)
	}
}

func ftest(x, y float64) float64 {
	return (math.Pow(x, 2) - math.Pow(y, 2)) / 500
}

func fpersp(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func feggbox(x, y float64) float64 {
	return (math.Sin(x) * math.Sin(y)) / 10
}

func fmoguls(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 30
}

func fsaddle(x, y float64) float64 {
	return (math.Pow(x, 2) - math.Pow(y, 2)) / 500
}
