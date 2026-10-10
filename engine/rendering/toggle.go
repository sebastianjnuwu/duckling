package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// toggle.go — interruptor pixel: trilho + pino dourado.
// Clique alterna On e retorna true.

type Toggle struct {
	X, Y, W, H float64
	On         bool
}

func (t *Toggle) Update() bool {
	mx, my := ebiten.CursorPosition()
	over := float64(mx) >= t.X && float64(mx) <= t.X+t.W &&
		float64(my) >= t.Y && float64(my) <= t.Y+t.H
	if over && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		t.On = !t.On
		return true
	}
	return false
}

func (t *Toggle) Draw(img *ebiten.Image) {
	notch(img, t.X, t.Y, t.W, t.H, 8, color.RGBA{16, 16, 30, 255})
	notch(img, t.X+4, t.Y+4, t.W-8, t.H-8, 4, color.RGBA{34, 48, 78, 255})
	kx := t.X + 8
	kc := color.RGBA{90, 100, 130, 255}
	if t.On {
		kx = t.X + t.W - 40
		kc = color.RGBA{232, 190, 72, 255}
	}
	notch(img, kx, t.Y+8, 32, 32, 6, kc)
	ebitenutil.DrawRect(img, kx+8, t.Y+8, 16, 4, color.RGBA{255, 255, 255, 60})
}
