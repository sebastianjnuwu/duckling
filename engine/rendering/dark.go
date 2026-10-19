package ui

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// dark.go — o escuro da lamparina: um preto por cima da cena com
// buracos circulares suaves (o raio de visão de cada pato).

type Dark struct {
	buf  *ebiten.Image
	hole *ebiten.Image
}

// NewDark monta o buffer do tamanho lógico e a textura do buraco.
func NewDark(w, h, hole int) *Dark {
	d := &Dark{buf: ebiten.NewImage(w, h)}
	img := image.NewRGBA(image.Rect(0, 0, hole, hole))
	c := float64(hole) / 2
	for y := 0; y < hole; y++ {
		for x := 0; x < hole; x++ {
			dx, dy := float64(x)+0.5-c, float64(y)+0.5-c
			a := 1 - (math.Hypot(dx, dy)/c-0.5)/0.5
			if a < 0 {
				a = 0
			}
			if a > 1 {
				a = 1
			}
			img.SetRGBA(x, y, color.RGBA{0, 0, 0, uint8(a * 255)})
		}
	}
	d.hole = ebiten.NewImageFromImage(img)
	return d
}

// Draw escurece a tela e reabre um círculo em cada foco (px de tela).
func (d *Dark) Draw(dst *ebiten.Image, focus [][2]float64, radius float64) {
	d.buf.Clear()
	d.buf.Fill(color.RGBA{6, 8, 14, 238})
	size := float64(d.hole.Bounds().Dx())
	for _, f := range focus {
		op := &ebiten.DrawImageOptions{Blend: ebiten.BlendDestinationOut}
		op.GeoM.Translate(-size/2, -size/2)
		op.GeoM.Scale(radius*2/size, radius*2/size)
		op.GeoM.Translate(f[0], f[1])
		d.buf.DrawImage(d.hole, op)
	}
	dst.DrawImage(d.buf, nil)
}
