package menu

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// menu_ground.go — gramado + sombra pixelada sob o pato.
// A grama é fixa; a sombra acompanha o pato (x).

func (m *Menu) drawGround(img *ebiten.Image, x, bob, p float64) {
	grow := min(1, p*1.5)
	blob(img, 340, 548, 240*grow, 36, 6, color.RGBA{30, 120, 50, 255}, false)
	blob(img, 340, 542, 210*grow, 30, 6, color.RGBA{60, 170, 80, 255}, false)
}

// blob elipse em blocos px; xadrez na borda se dither.
func blob(img *ebiten.Image, cx, cy, rx, ry, px float64, c color.Color, dither bool) {
	if rx < px || ry < px {
		return
	}
	cx, rx = math.Round(cx/px)*px, math.Round(rx/px)*px
	cy, ry = math.Round(cy/px)*px, math.Round(ry/px)*px
	W, H := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	for by := -ry; by < ry; by += px {
		for bx := -rx; bx < rx; bx += px {
			if cx+bx < 0 || cy+by < 0 || cx+bx+px > W || cy+by+px > H {
				continue
			}
			dx, dy := bx+px/2, by+px/2
			d := dx*dx/(rx*rx) + dy*dy/(ry*ry)
			if d > 1 || (dither && d > 0.6 && int((bx+by)/px)%2 == 0) {
				continue
			}
			for j := 0.0; j < px; j++ {
				for i := 0.0; i < px; i++ {
					img.Set(int(cx+bx+i), int(cy+by+j), c)
				}
			}
		}
	}
}
