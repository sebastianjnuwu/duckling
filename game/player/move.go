package actor

import (
	"image"
	"math"

	"duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2"
)

// move.go — física do personagem: entrada, animação e trava.
// A sheet só tem voo de lado, então no chão ele anda de frente
// (espelha no A/D). Parado, congela no frame base.

// Update anda conforme WASD/setas, desliza nos sólidos e
// prende nos limites (no alto, com folga pra cabeça não
// sair da tela). Passa por cima do que for mais baixo.
func (d *Duck) Update(x0, y0, x1, y1 int, solids []world.Solid) {
	sp := d.def.Speed * d.air(solids)
	// Trombou em NPC: meio passo por 1.2s (com carência).
	if d.slowT > 0 {
		d.slowT -= 1.0 / 60.0
		sp *= 0.5
	}
	if d.bumpT > 0 {
		d.bumpT -= 1.0 / 60.0
	}
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
	d.step(dx, dy, solids)
	d.Px = min(max(d.Px, float64(x0)), float64(x1))
	top := y0
	if d.z > 0 {
		top = min(y0+int(d.z)+20, y1)
	}
	d.Py = min(max(d.Py, float64(top)), float64(y1))
}

// step anda em fatias de 1px; bloqueou, quica pra trás com atrito.
func (d *Duck) step(dx, dy float64, solids []world.Solid) {
	n := max(1, int(math.Ceil(math.Abs(dx))), int(math.Ceil(math.Abs(dy))))
	sx, sy := dx/float64(n), dy/float64(n)
	bx, by := false, false
	for i := 0; i < n; i++ {
		if !hit(d.Px+sx, d.Py, d.z, solids) {
			d.Px += sx
		} else {
			bx = true
		}
		if !hit(d.Px, d.Py+sy, d.z, solids) {
			d.Py += sy
		} else {
			by = true
		}
	}
	// Quica só no 1º encosto (segurando contra o muro, descansa).
	if bx && !d.hx {
		d.kx = -dx * 1.5
	}
	if by && !d.hy {
		d.ky = -dy * 1.5
	}
	if !hit(d.Px+d.kx, d.Py, d.z, solids) {
		d.Px += d.kx
	} else {
		d.kx = 0
	}
	if !hit(d.Px, d.Py+d.ky, d.z, solids) {
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
	d.hx, d.hy = bx, by
	d.Anim(dx, dy)
}

// hit diz se a caixa dos pés (20x16) encosta em algum sólido
// mais alto que o pato (por cima da altura, passa direto).
func hit(x, y, z float64, solids []world.Solid) bool {
	r := image.Rect(int(x-10), int(y-8), int(x+10), int(y+8))
	for _, s := range solids {
		if z < s.H && r.Overlaps(s.Rect) {
			return true
		}
	}
	return false
}
