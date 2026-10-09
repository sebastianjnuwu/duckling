package actor

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// move.go — física do personagem: entrada, animação e trava.
// A sheet só tem voo de lado, então no chão ele anda de frente
// (espelha no A/D). Parado, congela no frame base.

// Update anda conforme WASD/setas, desliza nos sólidos e
// prende nos limites. Voando, passa por cima de tudo.
func (d *Duck) Update(x0, y0, x1, y1 int, solids []image.Rectangle) {
	sp := d.def.Speed * d.air()
	var dx, dy float64
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx -= sp
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx += sp
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy -= sp
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy += sp
	}
	sol := solids
	if d.flying {
		sol = nil
	}
	d.step(dx, dy, sol)
	d.Px = min(max(d.Px, float64(x0)), float64(x1))
	d.Py = min(max(d.Py, float64(y0)), float64(y1))
}

// step anda em fatias de 1px; bloqueou, quica pra trás com atrito.
func (d *Duck) step(dx, dy float64, solids []image.Rectangle) {
	n := max(1, int(math.Ceil(math.Abs(dx))), int(math.Ceil(math.Abs(dy))))
	sx, sy := dx/float64(n), dy/float64(n)
	bx, by := false, false
	for i := 0; i < n; i++ {
		if !hit(d.Px+sx, d.Py, solids) {
			d.Px += sx
		} else {
			bx = true
		}
		if !hit(d.Px, d.Py+sy, solids) {
			d.Py += sy
		} else {
			by = true
		}
	}
	// Quica de volta (só dispara quicado, senão treme colado).
	if bx && math.Abs(d.kx) < 0.5 {
		d.kx = -dx * 1.5
	}
	if by && math.Abs(d.ky) < 0.5 {
		d.ky = -dy * 1.5
	}
	if !hit(d.Px+d.kx, d.Py, solids) {
		d.Px += d.kx
	} else {
		d.kx = 0
	}
	if !hit(d.Px, d.Py+d.ky, solids) {
		d.Py += d.ky
	} else {
		d.ky = 0
	}
	d.kx *= 0.8
	d.ky *= 0.8
	if math.Abs(d.kx) < 0.05 {
		d.kx = 0
	}
	if math.Abs(d.ky) < 0.05 {
		d.ky = 0
	}
	d.Anim(dx, dy)
}

// hit diz se a caixa dos pés (20x16) encosta em algum sólido.
func hit(x, y float64, solids []image.Rectangle) bool {
	r := image.Rect(int(x-10), int(y-8), int(x+10), int(y+8))
	for _, s := range solids {
		if r.Overlaps(s) {
			return true
		}
	}
	return false
}
