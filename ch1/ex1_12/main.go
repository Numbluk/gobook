// handler returns a lissajous gif and accepts parameters
// to change attributes about the gif
package main

import (
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	var cycles = 5   // number of complete x oscillator revolutions
	var res = 0.001  // angular resolution
	var size = 100   // image canvas covers [-size..+size]
	var nframes = 64 // number of animation frames
	var delay = 8    // delay between frames in 10ms units

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		if k == "cycles" {
			cycles, _ = strconv.Atoi(v[0])
		} else if k == "res" {
			res, _ = strconv.ParseFloat(v[0], 64)
		} else if k == "size" {
			size, _ = strconv.Atoi(v[0])
		} else if k == "nframes" {
			nframes, _ = strconv.Atoi(v[0])
		} else if k == "delay" {
			delay, _ = strconv.Atoi(v[0])
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(float64(y)*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
