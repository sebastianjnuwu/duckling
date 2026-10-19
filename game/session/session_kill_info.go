package session

import (
	"duckling/engine/network"
	"duckling/game/match"
)

// session_kill_info.go — leitura do estado de vida: cooldown,
// corpos e a conta que decide a vitória dos impostores.

// Cooldown devolve os segundos que faltam pra matar de novo.
func (s *Session) Cooldown() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		return s.killCd.Seconds()
	}
	return float64(s.cdMS) / 1000
}

// Bodies devolve uma cópia dos corpos caídos.
func (s *Session) Bodies() []network.Body {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]network.Body(nil), s.bodies...)
}

// AliveCount conta os vivos (host): é o número que decide a vitória.
func (s *Session) AliveCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := 0
	if !s.ghosts[0] {
		n++
	}
	for _, p := range s.peers {
		if !s.ghosts[p.id] {
			n++
		}
	}
	return n
}

// impWinLocked: há impostor vivo e não são minoria.
func (s *Session) impWinLocked() bool {
	imp, crew := 0, 0
	for id, r := range s.plan.Roles {
		if s.ghosts[id] {
			continue
		}
		if r == match.Imp {
			imp++
		} else {
			crew++
		}
	}
	return imp > 0 && imp >= crew
}
