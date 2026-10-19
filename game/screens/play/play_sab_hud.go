package play

import (
	"fmt"
	"math"

	"duckling/engine/i18n"
	"duckling/game/match"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_sab_hud.go — os dois botões de sabotagem (embaixo à direita),
// com o tempo restante no rótulo.

func (p *Play) sabDraw(img *ebiten.Image) {
	if p.sess == nil || p.font == nil || p.phase != match.Playing || !p.imp() || p.ghost() {
		return
	}
	cd := p.sess.SabotageCooldown()
	p.sabLamp.Draw(img, sabLabel("sab.lamp", p.sess.LampLeft(), cd), p.font)
	p.sabGate.Draw(img, sabLabel("sab.gate", p.sess.GateLeft(), cd), p.font)
}

// sabLabel mostra o efeito ativo ou o cooldown; sem nada, só o nome.
func sabLabel(key string, left, cd float64) string {
	if left > 0 {
		return fmt.Sprintf("%s %.0f", i18n.T(key), math.Ceil(left))
	}
	if cd > 0 {
		return fmt.Sprintf("%s %.0f", i18n.T(key), math.Ceil(cd))
	}
	return i18n.T(key)
}
