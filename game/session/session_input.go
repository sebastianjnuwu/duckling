package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/player"
)

// session_input.go — input do convidado: bits segurados + bordas
// (pulo/voo) que o cliente reenvia; o host filtra a repetição.

const edgeDup = 120 * time.Millisecond

func (s *Session) hostInput(addr string, m network.Msg) {
	s.mu.Lock()
	p := s.peers[addr]
	if p == nil {
		s.mu.Unlock()
		return
	}
	now := time.Now()
	p.edge |= m.In &^ p.in
	if m.Edge&actor.BitJump != 0 && now.Sub(p.jumpAt) > edgeDup {
		p.edge |= actor.BitJump
		p.jumpAt = now
	}
	if m.Edge&actor.BitFly != 0 && now.Sub(p.flyAt) > edgeDup {
		p.edge |= actor.BitFly
		p.flyAt = now
	}
	p.prev, p.in = p.in, m.In
	p.last = now
	s.mu.Unlock()
}
