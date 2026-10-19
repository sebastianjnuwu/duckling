package play

import (
	"math"

	"duckling/engine/input"
	"duckling/game/match"
	"duckling/game/session"
)

// play_kill.go — o impostor local: acha o alvo no alcance e manda
// a intenção; o host confere posição, papel e cooldown.

const killLock = 0.5 // segundos parado depois de matar

// killUpdate lê a tecla MATAR (só impostor vivo, jogando).
func (p *Play) killUpdate() {
	if !p.canKillNow() || !input.JustPressed(input.Kill) {
		return
	}
	t := p.killTarget()
	if t < 0 {
		return
	}
	if p.sess.IsHost() {
		p.tryKill(0, t)
		return
	}
	p.sess.SendKill(t)
	p.lock[p.sess.ID()] = killLock
}

// canKillNow junta as condições locais (o host recheca as dele).
func (p *Play) canKillNow() bool {
	return p.sess != nil && p.phase == match.Playing && p.imp() &&
		!p.ghost() && p.sess.Cooldown() <= 0 && p.lock[p.localID()] <= 0
}

// killTarget é o pato vivo mais perto dentro do alcance (-1 = nenhum).
func (p *Play) killTarget() int {
	best, bd := -1, session.KillReach
	for id, d := range p.rem {
		if p.rghost[id] {
			continue
		}
		if v := math.Hypot(p.d.Px-d.Px, p.d.Py-d.Py); v < bd {
			best, bd = id, v
		}
	}
	return best
}

// localID é o número do jogador desta máquina (host = 0).
func (p *Play) localID() int {
	if p.sess == nil {
		return 0
	}
	return p.sess.ID()
}
