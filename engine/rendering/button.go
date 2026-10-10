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
	armed      bool
	Icon       Icon
	// ISize é a caixa do ícone (0 = 24 padrão).
	ISize float64
	// Disabled chumba: sem hover, sem clique, apagado.
	Disabled bool
}

func (b *Btn) Update() bool {
	mx, my := ebiten.CursorPosition()
	b.Hover = float64(mx) >= b.X && float64(mx) <= b.X+b.W &&
		float64(my) >= b.Y && float64(my) <= b.Y+b.H
	if b.Disabled {
		b.Hover = false
		return false
	}
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
		s := 24.0
		if b.ISize > 0 {
			s = b.ISize
		}
		tx = x + (b.W-(s+12+w))/2 + s + 12
		if w == 0 {
			tx = x + (b.W-s)/2 + s + 12
		}
		DrawIconSized(img, b.Icon, tx-s-12, y+b.H/2-s/2, s, tc)
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(tx, y+b.H/2-h/2)
	op.ColorScale.ScaleWithColor(tc)
	text.Draw(img, label, f, op)
	if b.Disabled {
		ebitenutil.DrawRect(img, x, y, b.W, b.H, color.RGBA{10, 10, 18, 80})
	}
}

// notch desenha retângulo de canto recortado (2 faixas).
func notch(img *ebiten.Image, x, y, w, h, n float64, c color.Color) {
	ebitenutil.DrawRect(img, x+n, y, w-2*n, h, c)
	ebitenutil.DrawRect(img, x, y+n, w, h-2*n, c)
}
