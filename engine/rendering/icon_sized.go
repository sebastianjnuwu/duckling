package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// icon_sized.go — ícone em caixa arbitrária (botão de destaque).
// PNG e fonte escalam; procedural cai no 24 centralizado.

func DrawIconSized(img *ebiten.Image, ic Icon, x, y, size float64, c color.Color) {
	if size == 24 {
		DrawIcon(img, ic, x, y, c)
		return
	}
	if src := customIcon(ic); src != nil {
		drawCustom(img, src, x, y, size, c)
		return
	}
	if drawFont(img, ic, x, y, size, c) {
		return
	}
	d := (size - 24) / 2
	DrawIcon(img, ic, x+d, y+d, c)
}
