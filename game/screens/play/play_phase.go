package play

import (
	"duckling/game/match"
	"duckling/game/player"
)

// play_phase.go — a fase manda: ao virar jogando, cada um nasce
// no ponto sorteado; ao voltar pro lobby, o elenco se refaz.

func (p *Play) syncPhase() {
	ph := match.Lobby
	if p.sess != nil {
		ph = p.sess.Phase()
	}
	if ph != p.phase {
		p.phase, p.placed = ph, false
		if ph == match.Lobby {
			p.resetCast()
		}
	}
	// Espera o "role" chegar (o spawn do convidado vem nele).
	if !p.placed && p.sess != nil {
		if x, y, ok := p.sess.Spawn(p.sess.ID()); ok {
			p.d.Snap(x, y, 0, 0, false)
			p.placed = true
		}
	}
}

// canMove trava o pato no fim (reunião entra na Fase 5).
func (p *Play) canMove() bool {
	return p.phase != match.End && p.phase != match.Meeting
}

// resetCast devolve o pato e os remotos aos nascimentos locais.
func (p *Play) resetCast() {
	id := 0
	if p.sess != nil {
		id = p.sess.ID()
	}
	p.d.Px, p.d.Py = p.spawnAt(id)
	p.rem = map[int]*actor.Duck{}
	p.rghost = map[int]bool{}
}
