package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// toast.go — avisos empilhados embaixo (máx 3): entram
// deslizando de baixo, somem com fade. Sem estado de jogo.

const (
	toastDur = 2.2
	toastMax = 3
)

type toastItem struct {
	msg string
	t   float64
}

type Toast struct {
	items []toastItem
	// Right gruda no canto inferior direito em vez do centro.
	Right bool
}

// Show empilha o aviso (cheio? sai o mais velho).
func (t *Toast) Show(msg string) {
	t.items = append(t.items, toastItem{msg, toastDur})
	if len(t.items) > toastMax {
		t.items = t.items[1:]
	}
}

func (t *Toast) Update() {
	keep := t.items[:0]
	for _, it := range t.items {
		it.t -= 1.0 / 60.0
		if it.t > 0 {
			keep = append(keep, it)
		}
	}
	t.items = keep
}

func (t *Toast) Draw(img *ebiten.Image, f *text.GoTextFace) {
	W, H := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	y := H - 32.0
	for i := len(t.items) - 1; i >= 0; i-- {
		it := t.items[i]
		lines := wrap(it.msg, f, 900)
		_, lh := text.Measure("A", f, 0)
		bw := 0.0
		for _, ln := range lines {
			if w, _ := text.Measure(ln, f, 0); w > bw {
				bw = w
			}
		}
		bw += 48
		bh := float64(len(lines))*(lh+6) - 6 + 28
		in := min(1, (toastDur-it.t)/0.25)
		a := min(in, min(1, it.t/0.4))
		y -= bh
		by := y + (1-in)*20
		bx := W/2 - bw/2
		if t.Right {
			bx = W - bw - 24
		}
		bg := color.RGBA{16, 16, 30, uint8(230 * a)}
		fg := color.RGBA{235, 240, 235, uint8(255 * a)}
		ebitenutil.DrawRect(img, bx, by, bw, bh, bg)
		ebitenutil.DrawRect(img, bx, by, bw, 4, color.RGBA{232, 190, 72, uint8(255 * a)})
		for j, ln := range lines {
			w, _ := text.Measure(ln, f, 0)
			op := &text.DrawOptions{}
			op.GeoM.Translate(bx+(bw-w)/2, by+14+float64(j)*(lh+6))
			op.ColorScale.ScaleWithColor(fg)
			text.Draw(img, ln, f, op)
		}
		y -= 8
	}
}
