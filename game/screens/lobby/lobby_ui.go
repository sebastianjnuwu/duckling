package lobby

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// lobby_ui.go — miudezas de desenho: pílula, número ouro,
// texto à esquerda e as duas tintas.

func gold() color.Color { return color.RGBA{232, 190, 72, 255} }

func cream() color.Color { return color.RGBA{235, 240, 235, 255} }

func pill(img *ebiten.Image, x, y, w, h float64) {
	ebitenutil.DrawRect(img, x, y, w, h, color.RGBA{48, 66, 100, 255})
	ebitenutil.DrawRect(img, x+4, y+4, w-8, h-8, color.RGBA{16, 16, 30, 255})
}

func num(img *ebiten.Image, s string, f *text.GoTextFace, cx, y float64) {
	w, _ := text.Measure(s, f, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate(cx-w/2, y)
	op.ColorScale.ScaleWithColor(gold())
	text.Draw(img, s, f, op)
}

func left(img *ebiten.Image, s string, f *text.GoTextFace, x, y float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	text.Draw(img, s, f, op)
}
