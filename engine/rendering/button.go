package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// button.go — botão pixel-art: canto recortado, paleta
// fria de lagoa com face dourada no hover/clique.
// Texto e ícone acompanham (creme, tinta no ouro).

// Press vai a 1 no clique e decai até 0 (chamar Update todo frame).
type Btn struct {
	X, Y, W, H float64
	Hover      bool
	Press      float64
	Icon       Icon
}

func (b *Btn) Update() bool {
	mx, my := ebiten.CursorPosition()
	b.Hover = float64(mx) >= b.X && float64(mx) <= b.X+b.W &&
		float64(my) >= b.Y && float64(my) <= b.Y+b.H
	clicked := b.Hover && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	if clicked {
		b.Press = 1
	}
	b.Press = max(0, b.Press-1.0/12.0)
	return clicked
}

func (b *Btn) Draw(img *ebiten.Image, label string, f *text.GoTextFace) {
	dy := b.Press * 4
	face := color.RGBA{48, 66, 100, 255}
	if b.Hover && b.Press == 0 {
		dy = -2
		face = color.RGBA{232, 190, 72, 255}
	}
	if b.Press > 0.5 {
		face = color.RGBA{186, 148, 50, 255}
	}
	x, y := b.X, b.Y+dy
	hi, _, lo := bevel(face)
	ebitenutil.DrawRect(img, x+6, y+8, b.W, b.H, color.RGBA{0, 0, 0, 130})
	ink := color.RGBA{16, 16, 30, 255}
	notch(img, x, y, b.W, b.H, 8, ink)
	notch(img, x+4, y+4, b.W-8, b.H-8, 4, face)
	ebitenutil.DrawRect(img, x+12, y+4, b.W-24, 4, hi)
	ebitenutil.DrawRect(img, x+12, y+b.H-8, b.W-24, 4, lo)
	if f == nil {
		return
	}
	w, h := text.Measure(label, f, 0)
	tc := color.RGBA{235, 240, 235, 255}
	if b.Hover && b.Press <= 0.5 {
		tc = color.RGBA{16, 16, 30, 255}
	}
	tx := x + b.W/2 - w/2
	if b.Icon != IconNone {
		tx = x + (b.W-(24+12+w))/2 + 36
		DrawIcon(img, b.Icon, tx-36, y+b.H/2-12, tc)
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(tx, y+b.H/2-h/2)
	op.ColorScale.ScaleWithColor(tc)
	text.Draw(img, label, f, op)
}

// notch desenha retângulo de canto recortado (2 faixas).
func notch(img *ebiten.Image, x, y, w, h, n float64, c color.Color) {
	ebitenutil.DrawRect(img, x+n, y, w-2*n, h, c)
	ebitenutil.DrawRect(img, x, y+n, w, h-2*n, c)
}

// bevel deriva claro/escuro da cor da face.
func bevel(face color.RGBA) (hi, mid, lo color.Color) {
	hi = color.RGBA{uint8(min(255, int(face.R)+90)), uint8(min(255, int(face.G)+90)), uint8(min(255, int(face.B)+90)), 255}
	mid = color.RGBA{uint8(max(0, int(face.R)-30)), uint8(max(0, int(face.G)-30)), uint8(max(0, int(face.B)-30)), 255}
	lo = color.RGBA{20, 20, 30, 255}
	return hi, mid, lo
}
