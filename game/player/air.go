package actor

import (
	"image"

	"duckling/engine/audio"
	"duckling/engine/input"
	"duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// air.go — pulo e voo: ESPAÇO pula, E ou X voa (pousa em cercas/pedras).
// Alturas de referência no Tiled: h≤25 pulo, h≤34 voo, h>34 muro.

const (
	jumpV    = 5.2
	jumpG    = 0.5
	jumpCut  = 2.5  // soltou o pulo subindo: corta (toque = pulinho)
	jumpBuf  = 0.12 // buffer: apertou um pouco antes, vale
	coyote   = 0.10 // saiu da borda agora: ainda pode pular
	flyMax   = 5.0
	flyBoost = 1.25
	flyH     = 38.0
	flyHold  = 0.2
	flyRise  = 1.8 // subida do voo (decolagem suave)
	flyDive  = 3.0 // segurando baixo no voo: mergulha
	flyRegen = 2.0 // estamina de volta no chão, por s

	// JumpPeak é o teto do pulo (v²/2g; ~30 na integração do jogo).
	JumpPeak = jumpV * jumpV / (2 * jumpG) // ~27 contínuo

	// FlyPeak é o teto do voo.
	FlyPeak = flyH // 38
)

// air atualiza piso, gravidade, pulo e voo; devolve multiplicador de vel.
func (d *Duck) air(solids []world.Solid) float64 {
	d.ground = Floor(d.Px, d.Py, d.z, solids)
	onGround := d.z <= d.ground+0.1
	if onGround {
		d.coyT = coyote
	} else {
		d.coyT -= 1.0 / 60.0
	}
	if d.bufT > 0 {
		d.bufT -= 1.0 / 60.0
	}
	if input.JustPressed(input.Jump) {
		d.bufT = jumpBuf
	}
	held := input.Held(input.Jump)
	// Soltou subindo: corta o pulo (toque rápido = pulinho).
	if !held && d.jumpHeld && d.vz > jumpCut && !d.flying {
		d.vz = jumpCut
	}
	d.jumpHeld = held
	if d.bufT > 0 && (onGround || d.coyT > 0) && !d.flying {
		d.bufT, d.coyT, d.vz = 0, 0, jumpV
		_ = audio.Play("jump")
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		d.holdT += 1.0 / 60.0
	} else {
		d.holdT = 0
	}
	if input.JustPressed(input.Fly) || inpututil.IsKeyJustPressed(ebiten.KeyE) {
		d.wantFly = !d.wantFly
	}
	// Paredão colado corta o voo (e nem deixa levantar).
	wall := d.walled(solids)
	if wall {
		d.wantFly = false
	}
	d.flying = (d.wantFly || input.Held(input.Fly) || ebiten.IsKeyPressed(ebiten.KeyE) || d.holdT > flyHold) && d.flyT > 0 && !wall
	if !d.flying {
		d.wantFly = false
	}
	if d.flying {
		d.flyT -= 1.0 / 60.0
		// Teto desce no último 1s: cansou, plana até o piso.
		ceiling := d.ground + flyH*min(max(d.flyT, 0), 1)
		if input.Held(input.Down) {
			// Mergulha: desce rápido até o piso.
			d.z = max(d.ground, d.z-flyDive)
		} else if up := min(2.5, (ceiling-d.z)*0.3+0.3); up > 0 {
			d.z = min(ceiling, d.z+min(up, flyRise))
		}
		d.vz = 0
		return flyBoost
	}
	if onGround && d.flyT < flyMax {
		d.flyT += flyRegen / 60.0
	}
	if d.z > d.ground || d.vz != 0 {
		d.z += d.vz
		d.vz = max(d.vz-jumpG, -4.5)
		if d.vz <= 0 && d.z <= d.ground {
			d.z, d.vz = d.ground, 0
		}
	} else if d.z < d.ground {
		d.z = d.ground
	}
	return 1
}

// walled diz se o pato tá colado em paredão mais alto que o voo.
func (d *Duck) walled(solids []world.Solid) bool {
	r := image.Rect(int(d.Px-11), int(d.Py-9), int(d.Px+11), int(d.Py+9))
	for _, s := range solids {
		if s.H > d.z+flyH && r.Overlaps(s.Rect) {
			return true
		}
	}
	return false
}
