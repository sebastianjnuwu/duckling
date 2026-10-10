package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// flag.go — bandeiras pixel e botão todo de bandeira.

// Larguras do losango da BR por fileira.
var flagFW = []float64{.12, .3, .48, .66, .66, .48, .3, .12}

// DrawFlagSize bandeira pixel na caixa (x,y,w,h); apagada se !active.
func DrawFlagSize(img *ebiten.Image, code string, x, y, w, h float64, active bool) {
	r := func(f float64) float64 { return math.Round(f) }
	rect := func(fx, fy, fw, fh float64, c color.Color) {
		ebitenutil.DrawRect(img, x+r(fx), y+r(fy), r(fw), r(fh), c)
	}
	switch code {
	case "br":
		rect(0, 0, w, h, color.RGBA{0, 150, 60, 255})
		for i := 0.0; i < 8; i++ {
			rect((w-flagFW[int(i)]*w)/2, i*h/8, flagFW[int(i)]*w, h/8, color.RGBA{255, 220, 0, 255})
		}
		s := r(h * .4)
		rect((w-s)/2, (h-s)/2, s, s, color.RGBA{30, 60, 180, 255})
	case "us":
		rect(0, 0, w, h, color.RGBA{235, 235, 235, 255})
		for i := 0.0; i < 5; i++ {
			rect(0, i*h*2/9, w, h/9, color.RGBA{190, 40, 50, 255})
		}
		rect(0, 0, w*.47, h*.56, color.RGBA{40, 60, 160, 255})
	default:
		rect(0, 0, w, h, color.RGBA{190, 40, 45, 255})
		rect(0, h*.25, w, h*.5, color.RGBA{255, 210, 60, 255})
	}
	if !active {
		ebitenutil.DrawRect(img, x, y, w, h, color.RGBA{0, 0, 0, 140})
	}
}

// FlagButton botão todo de bandeira: borda, arte cheia,
// apagado se !active, ouro no hover, afunda no clique.
func FlagButton(img *ebiten.Image, b Btn, code string, active bool) {
	edge, y := color.RGBA{16, 16, 30, 255}, b.Y+b.Press*4
	if b.Hover && b.Press == 0 {
		edge, y = color.RGBA{232, 190, 72, 255}, y-2
	}
	notch(img, b.X, y, b.W, b.H, 8, edge)
	DrawFlagSize(img, code, b.X+4, y+4, b.W-8, b.H-8, active)
	if b.Press > 0.5 {
		ebitenutil.DrawRect(img, b.X+4, y+4, b.W-8, b.H-8, color.RGBA{0, 0, 0, 110})
	}
}
