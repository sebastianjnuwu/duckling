package shop

import (
	"image/color"

	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// shop_cards.go — geometria dos 4 cartões e o quadro de destaque.

const (
	cardY   = 190.0
	cardW   = 256.0
	cardH   = 360.0
	cardGap = 20.0
)

// cardX devolve o x do cartão i (fileira centrada).
func cardX(i int) float64 {
	w := float64(actor.HatCount)*cardW + float64(actor.HatCount-1)*cardGap
	return (1280-w)/2 + float64(i)*(cardW+cardGap)
}

// cardHit diz se o mouse está no cartão i (clique e hover).
func cardHit(i int) bool {
	x, y := ebiten.CursorPosition()
	cx := cardX(i)
	return float64(x) >= cx && float64(x) <= cx+cardW &&
		float64(y) >= cardY && float64(y) <= cardY+cardH
}

// frame desenha a borda de espessura t em volta do retângulo.
func frame(img *ebiten.Image, x, y, w, h, t float64, c color.Color) {
	ebitenutil.DrawRect(img, x, y, w, t, c)
	ebitenutil.DrawRect(img, x, y+h-t, w, t, c)
	ebitenutil.DrawRect(img, x, y, t, h, c)
	ebitenutil.DrawRect(img, x+w-t, y, t, h, c)
}
