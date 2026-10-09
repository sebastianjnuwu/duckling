package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// air.go — pulo e voo: ESPAÇO pula, segurar X voa (3s, recarrega
// no chão). Voando ele passa por cima dos sólidos.

const (
	jumpV    = 5.5
	jumpG    = 0.5
	flyMax   = 3.0
	flyBoost = 1.8
)

// air atualiza altura e voo; devolve o multiplicador de velocidade.
func (d *Duck) air() float64 {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && d.z == 0 && !d.flying {
		d.vz = jumpV
	}
	d.flying = ebiten.IsKeyPressed(ebiten.KeyX) && d.flyT > 0
	if d.flying {
		d.flyT -= 1.0 / 60.0
	} else if d.z == 0 && d.flyT < flyMax {
		d.flyT += 1.5 / 60.0
	}
	if d.z > 0 || d.vz != 0 {
		d.z += d.vz
		d.vz -= jumpG
		if d.z <= 0 {
			d.z, d.vz = 0, 0
		}
	}
	if d.flying {
		return flyBoost
	}
	return 1
}
