package session

import "time"

// session_tick.go — manutenção por frame e o encerramento:
// derruba quem ficou mudo e avisa quando o host some.

// Tick faz a manutenção por frame: derruba quem sumiu e marca
// host perdido. Chamado pelo play.
func (s *Session) Tick() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for addr, p := range s.peers {
		if now.Sub(p.last) > peerTimeout {
			delete(s.peers, addr)
			s.events = append(s.events, "net.left")
		}
	}
	if !s.host && s.joined && s.fatal == "" && now.Sub(s.lastRecv) > peerTimeout {
		s.fatal = "net.hostgone"
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
