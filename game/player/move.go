package actor

import (
	"math"

	"duckling/engine/input"
	"duckling/game/map"
)

// move.go — física do personagem: aceleração, animação e trava.
// Anda com peso (acelera/freia), diagonal não corre mais que reto.

// Aceleração por estado (px/frame²): chão firme, ar obediente,
// voo macio. Chegar a 2.6px/f leva ~3 frames no chão.
const (
	groundAcc = 0.9
	airAcc    = 0.4
	flyAcc    = 0.55
	diagK     = 0.7071 // 1/√2: diagonal anda igual ao reto
)

// Update anda conforme WASD/setas, desliza nos sólidos e
// prende nos limites (no alto, com folga pra cabeça não
// sair da tela). Passa por cima do que for mais baixo.
func (d *Duck) Update(x0, y0, x1, y1 int, solids []world.Solid) {
	sp := d.def.Speed * d.air(solids)
	var tx, ty float64
	if input.Held(input.Left) {
		tx -= sp
	}
	if input.Held(input.Right) {
		tx += sp
	}
	if input.Held(input.Up) {
		ty -= sp
	}
	if input.Held(input.Down) {
		ty += sp
	}
	if tx != 0 && ty != 0 {
		tx, ty = tx*diagK, ty*diagK
	}
	acc := groundAcc
	if d.flying {
		acc = flyAcc
	} else if d.z > d.ground+0.1 {
		acc = airAcc
	}
	d.vx = approach(d.vx, tx, acc)
	d.vy = approach(d.vy, ty, acc)
	d.step(d.vx, d.vy, solids)
	d.Px = min(max(d.Px, float64(x0)), float64(x1))
	top := y0
	if d.z > 0 {
		top = min(y0+int(d.z)+20, y1)
	}
	d.Py = min(max(d.Py, float64(top)), float64(y1))
}

// approach leva v até t no máximo a por frame (sem passar do alvo).
func approach(v, t, a float64) float64 {
	if v < t {
		return min(v+a, t)
	}
	return max(v-a, t)
}

// step anda em fatias de 1px; desliza suavemente em paredes e quinas.
func (d *Duck) step(dx, dy float64, solids []world.Solid) {
	n := max(1, int(math.Ceil(math.Abs(dx))), int(math.Ceil(math.Abs(dy))))
	sx, sy := dx/float64(n), dy/float64(n)
	for i := 0; i < n; i++ {
		if !hit(d.Px+sx, d.Py, d.z, solids) {
			d.Px += sx
		}
		if !hit(d.Px, d.Py+sy, d.z, solids) {
			d.Py += sy
		}
	}
	d.Anim(dx, dy)
}
