package ui

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// toast.go — aviso temporário centralizado embaixo: Show
// liga por 2.2s, some com fade. Sem estado de jogo aqui.

type Toast struct {
	t float64
}

// Show liga o aviso por 2.2s.
func (t *Toast) Show() { t.t = 2.2 }

func (t *Toast) Update() { t.t = max(0, t.t-1.0/60.0) }

func (t *Toast) Draw(img *ebiten.Image, msg string, f *text.GoTextFace) {
	if t.t <= 0 {
		return
	}
	lines := wrapToast(msg, f, 900)
	_, lh := text.Measure("A", f, 0)
	bw := 0.0
	for _, ln := range lines {
		if w, _ := text.Measure(ln, f, 0); w > bw {
			bw = w
		}
	}
	bw += 48
	bh := float64(len(lines))*(lh+6) - 6 + 28
	a := min(1, t.t/0.4)
	bg := color.RGBA{16, 16, 30, uint8(230 * a)}
	fg := color.RGBA{235, 240, 235, uint8(255 * a)}
	W, H := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	bx, by := W/2-bw/2, H-bh-32
	ebitenutil.DrawRect(img, bx, by, bw, bh, bg)
	ebitenutil.DrawRect(img, bx, by, bw, 4, color.RGBA{232, 190, 72, uint8(255 * a)})
	for i, ln := range lines {
		w, _ := text.Measure(ln, f, 0)
		op := &text.DrawOptions{}
		op.GeoM.Translate(bx+(bw-w)/2, by+14+float64(i)*(lh+6))
		op.ColorScale.ScaleWithColor(fg)
		text.Draw(img, ln, f, op)
	}
}

// wrapToast quebra em linhas de até maxW.
func wrapToast(msg string, f *text.GoTextFace, maxW float64) []string {
	var out []string
	cur := ""
	for _, w := range strings.Fields(msg) {
		t := strings.TrimSpace(cur + " " + w)
		if tw, _ := text.Measure(t, f, 0); tw > maxW && cur != "" {
			out = append(out, cur)
			cur = w
		} else {
			cur = t
		}
	}
	return append(out, cur)
}
