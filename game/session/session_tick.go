package session

import (
	"time"

	"duckling/game/match"
)

// session_tick.go — manutenção por frame e o encerramento:
// derruba quem ficou mudo e avisa quando o host some.

// Tick faz a manutenção por frame: derruba quem sumiu e marca
// host perdido. Chamado pelo play.
func (s *Session) Tick() {
	s.mu.Lock()
	now := time.Now()
	s.tickCooldown(now)
	lost := false
	for addr, p := range s.peers {
		if now.Sub(p.last) > peerTimeout {
			if s.phase == match.Playing {
				s.ghosts[p.id] = true
				lost = true
			}
			delete(s.peers, addr)
			s.events = append(s.events, "net.left")
		}
	}
	if !s.host && s.joined && s.fatal == "" && now.Sub(s.lastRecv) > peerTimeout {
		s.fatal = "net.hostgone"
	}
	s.mu.Unlock()
	if lost {
		s.refreshTasks()
	}
}

// tickCooldown corre só jogando: reunião/lobby pausam o relógio.
func (s *Session) tickCooldown(now time.Time) {
	if !s.host {
		return
	}
	d := time.Duration(0)
	if !s.cdAt.IsZero() {
		d = now.Sub(s.cdAt)
	}
	s.cdAt = now
	if s.phase != match.Playing || s.killCd <= 0 {
		return
	}
	if s.killCd -= d; s.killCd < 0 {
		s.killCd = 0
	}
}

// TakeEvents entrega e limpa os avisos pendentes (chaves i18n).
func (s *Session) TakeEvents() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	e := s.events
	s.events = nil
	return e
}

// Fatal devolve o motivo de fim (chave i18n) ou "".
func (s *Session) Fatal() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.fatal
}

// Close encerra sockets e goroutines (idempotente).
func (s *Session) Close() {
	s.once.Do(func() {
		close(s.stop)
		s.ep.Close()
	})
}
