package tutorial

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // primeira cor da paleta
	blackIndex = 1 // próxima cor da paleta
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre de [-size..+size]
		nframes = 64    // número de frames da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // frequencia relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // ignorando erros de codificação
}
