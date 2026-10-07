package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// slider.go — barra arrastável: clica ou desliza o pino.
// Update retorna true quando Pct muda.

type Slider struct {
	X, Y, W, H float64
	Pct        int
	drag       bool
}

func (s *Slider) Update() bool {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.drag = false
		return false
	}
	mx, my := ebiten.CursorPosition()
	over := float64(mx) >= s.X-8 && float64(mx) <= s.X+s.W+8 &&
		float64(my) >= s.Y-12 && float64(my) <= s.Y+s.H+12
	if !s.drag && !(over && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)) {
		return false
	}
	s.drag = true
	p := max(0, min(100, int((float64(mx)-s.X)/s.W*100)))
	if p == s.Pct {
		return false
	}
	s.Pct = p
	return true
}

func (s *Slider) Draw(img *ebiten.Image) {
	ink := color.RGBA{16, 16, 30, 255}
	notch(img, s.X, s.Y, s.W, s.H, 6, ink)
	fw := (s.W - 8) * float64(s.Pct) / 100
	if fw > 0 {
		ebitenutil.DrawRect(img, s.X+4, s.Y+4, fw, s.H-8, color.RGBA{232, 190, 72, 255})
	}
	kx := min(max(s.X+4+fw-8, s.X+4), s.X+s.W-20)
	notch(img, kx, s.Y-4, 16, s.H+8, 4, color.RGBA{235, 240, 235, 255})
}
