package actor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// feathers.go — penas amarelas+vermelhas (as cores dos dois patos).

// feather solta n penas nos pés.
func (d *Duck) feather(n int) {
	d.fx.Burst(d.Px, d.Py+8, n,
		color.RGBA{245, 220, 90, 255},
		color.RGBA{235, 240, 235, 255},
		color.RGBA{230, 120, 60, 255})
}

// DrawFx desenha as penas com câmera.
func (d *Duck) DrawFx(img *ebiten.Image, camX, camY, zoom float64) {
	d.fx.DrawAt(img, camX, camY, zoom)
}
