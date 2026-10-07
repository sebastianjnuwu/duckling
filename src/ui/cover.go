package ui

import "github.com/hajimehoshi/ebiten/v2"

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
