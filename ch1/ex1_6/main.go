// Lissajous generates GIF animations of random Lissajous figures
// The color palette is a range of colors.
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

// composite literal: a compact notation for instantiating any of
// Go's composite types from a sequence of element values
var green color.RGBA = color.RGBA{0x00, 0xff, 0x00, 0x00}
var red color.RGBA = color.RGBA{0xff, 0x00, 0x00, 0x00}
var blue color.RGBA = color.RGBA{0x00, 0x00, 0xff, 0x00}
var black color.RGBA = color.RGBA{0xff, 0xff, 0xff, 0x00}
var palette = []color.Color{black, green, red, blue}

const (
	blackIndex = 0 // next color in palette
	greenIndex = 1 // first color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), random_color_index(palette))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func random_color_index(palette []color.Color) uint8 {
	return uint8(rand.Intn(len(palette) - 1))
}