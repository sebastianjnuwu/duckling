package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// icon.go — ícones pixel 24px, em blocos, sem alocar por
// frame (tabelas em var de pacote). Cor acompanha o texto.

type Icon int

const (
	IconNone Icon = iota
	IconPlay
	IconShop
	IconGear
	IconDoor
	IconNote
	IconSpk
	IconGlobe
	IconLan
	IconLock
	IconPlus
	IconMinus
	IconUser
)

// Posição x dos cursores do IconGear por fileira.
var gearKX = []float64{2, 12, 7}

// DrawIcon desenha o ícone numa caixa 24x24 em (x,y).
// Ordem: PNG custom, fonte, procedural.
func DrawIcon(img *ebiten.Image, ic Icon, x, y float64, c color.Color) {
	if src := customIcon(ic); src != nil {
		drawCustom(img, src, x, y, 24, c)
		return
	}
	if drawFont(img, ic, x, y, 24, c) {
		return
	}
	switch ic {
	case IconPlay:
		for i := 0.0; i < 7; i++ {
			w := (i + 1) * 4
			if i > 3 {
				w = (7 - i) * 4
			}
			ebitenutil.DrawRect(img, x+2, y+1+i*3, w, 3, c)
		}
	case IconShop:
		ebitenutil.DrawRect(img, x+2, y+2, 5, 3, c)
		ebitenutil.DrawRect(img, x+2, y+5, 20, 11, c)
		ebitenutil.DrawRect(img, x+5, y+18, 3, 3, c)
		ebitenutil.DrawRect(img, x+16, y+18, 3, 3, c)
	case IconGear:
		for r := 0.0; r < 3; r++ {
			ebitenutil.DrawRect(img, x, y+r*8+2, 24, 3, c)
			ebitenutil.DrawRect(img, x+gearKX[int(r)], y+r*8, 6, 7, c)
		}
	case IconDoor:
		ebitenutil.DrawRect(img, x+2, y, 20, 3, c)
		ebitenutil.DrawRect(img, x+2, y+21, 20, 3, c)
		ebitenutil.DrawRect(img, x+2, y, 3, 24, c)
		ebitenutil.DrawRect(img, x+19, y, 3, 24, c)
		ebitenutil.DrawRect(img, x+15, y+11, 3, 3, c)
	case IconNote:
		ebitenutil.DrawRect(img, x+15, y, 4, 18, c)
		ebitenutil.DrawRect(img, x+15, y, 10, 4, c)
		ebitenutil.DrawRect(img, x+20, y+4, 5, 6, c)
		ebitenutil.DrawRect(img, x+2, y+15, 13, 4, c)
		ebitenutil.DrawRect(img, x+2, y+19, 11, 5, c)
	case IconSpk:
		ebitenutil.DrawRect(img, x, y+8, 6, 8, c)
		ebitenutil.DrawRect(img, x+6, y+4, 6, 16, c)
		ebitenutil.DrawRect(img, x+12, y, 6, 24, c)
		ebitenutil.DrawRect(img, x+18, y+6, 3, 4, c)
		ebitenutil.DrawRect(img, x+18, y+14, 3, 4, c)
	case IconGlobe:
		ebitenutil.DrawRect(img, x+6, y, 12, 3, c)
		ebitenutil.DrawRect(img, x+6, y+21, 12, 3, c)
		ebitenutil.DrawRect(img, x, y+6, 3, 12, c)
		ebitenutil.DrawRect(img, x+21, y+6, 3, 12, c)
		ebitenutil.DrawRect(img, x+10, y+3, 4, 18, c)
		ebitenutil.DrawRect(img, x+3, y+10, 18, 4, c)
	case IconLan:
		ebitenutil.DrawRect(img, x, y+6, 8, 8, c)
		ebitenutil.DrawRect(img, x+16, y+6, 8, 8, c)
		ebitenutil.DrawRect(img, x+8, y+10, 8, 4, c)
	case IconLock:
		ebitenutil.DrawRect(img, x+4, y+10, 16, 11, c)
		ebitenutil.DrawRect(img, x+7, y+4, 3, 9, c)
		ebitenutil.DrawRect(img, x+14, y+4, 3, 9, c)
		ebitenutil.DrawRect(img, x+7, y+4, 10, 3, c)
	case IconPlus:
		ebitenutil.DrawRect(img, x+10, y+4, 4, 16, c)
		ebitenutil.DrawRect(img, x+4, y+10, 16, 4, c)
	case IconMinus:
		ebitenutil.DrawRect(img, x+4, y+10, 16, 4, c)
	case IconUser:
		ebitenutil.DrawRect(img, x+9, y+2, 6, 6, c)
		ebitenutil.DrawRect(img, x+6, y+10, 12, 10, c)
	}
}
