package actor

import (
	"image"

	"duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// air.go — pulo e voo: ESPAÇO pula, E ou X voa (segura ou
// 1 toque alterna; 5s, recarrega no chão).
// No alto ele passa por cima do que for mais baixo (voo: 34px).
// Colado em paredão (alto além do voo), nem levanta.

const (
	jumpV    = 5.5
	jumpG    = 0.5
	flyMax   = 5.0
	flyBoost = 1.3
	flyH     = 34.0
	flyHold  = 0.2
)

// air atualiza altura e voo; devolve o multiplicador de velocidade.
func (d *Duck) air(solids []world.Solid) float64 {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && d.z == 0 && !d.flying {
		d.vz = jumpV
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		d.holdT += 1.0 / 60.0
	} else {
		d.holdT = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeyE) {
		d.wantFly = !d.wantFly
	}
	// Paredão colado corta o voo (e nem deixa levantar).
	wall := d.walled(solids)
	if wall {
		d.wantFly = false
	}
	d.flying = (d.wantFly || ebiten.IsKeyPressed(ebiten.KeyX) || d.holdT > flyHold) && d.flyT > 0 && !wall
	if !d.flying {
		d.wantFly = false
	}
	if d.flying {
		d.flyT -= 1.0 / 60.0
		// Teto desce no último 1s: cansou, plana até o chão.
		ceiling := flyH * min(max(d.flyT, 0), 1)
		// Sobe rápido e amortece chegando (sem pop no cruzeiro).
		if up := min(2.5, (ceiling-d.z)*0.3+0.3); up > 0 {
			d.z = min(ceiling, d.z+up)
		}
		d.vz = 0
		return flyBoost
	}
	if d.z == 0 && d.flyT < flyMax {
		d.flyT += 1.5 / 60.0
	}
	if d.z > 0 || d.vz != 0 {
		d.z += d.vz
		d.vz = max(d.vz-jumpG, -4.5)
		if d.z <= 0 {
			d.z, d.vz = 0, 0
		}
	}
	return 1
}

// walled diz se o pato tá colado em paredão (alto além do voo).
func (d *Duck) walled(solids []world.Solid) bool {
	r := image.Rect(int(d.Px-11), int(d.Py-9), int(d.Px+11), int(d.Py+9))
	for _, s := range solids {
		if s.H > flyH && r.Overlaps(s.Rect) {
			return true
		}
	}
	return false
}
