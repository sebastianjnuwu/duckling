package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// cover.go — fundo em "cover": preenche a tela mantendo proporção,
// centralizado. Sem cover não há bg visível nas bordas.

func DrawCover(img, bg *ebiten.Image) {
	w, h := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	bw, bh := float64(bg.Bounds().Dx()), float64(bg.Bounds().Dy())
	s := max(w/bw, h/bh)
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	op.GeoM.Scale(s, s)
	op.GeoM.Translate(w/2-bw*s/2, h/2-bh*s/2)
	img.DrawImage(bg, op)
}

// Dim escurece a tela atrás do modal.
func Dim(img *ebiten.Image) {
	ebitenutil.DrawRect(img, 0, 0, 1280, 720, color.RGBA{0, 0, 0, 160})
}

// Panel desenha a caixa do modal: sombra, canto recortado,
// filete de latão e miolo — mesmo idioma dos botões.
func Panel(img *ebiten.Image, x, y, w, h float64) {
	ebitenutil.DrawRect(img, x+6, y+8, w, h, Shadow)
	ink := Ink
	notch(img, x, y, w, h, 8, ink)
	notch(img, x+4, y+4, w-8, h-8, 6, Brass)
	notch(img, x+8, y+8, w-16, h-16, 4, Slate)
	ebitenutil.DrawRect(img, x+16, y+8, w-32, 4, Gold)
}
