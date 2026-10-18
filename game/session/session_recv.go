package session

import (
	"time"

	"duckling/engine/network"
)

// session_recv.go — o lado do convidado na partida: ler o
// snapshot recebido e mandar o input pro host.

// States devolve o último estado conhecido de cada pato.
func (s *Session) States() []network.Player {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]network.Player, 0, len(s.states))
	for _, p := range s.states {
		out = append(out, p)
	}
	return out
}

// SendInput manda a entrada segurada pro host (no máximo 30/s).
func (s *Session) SendInput(bits uint8) {
	now := time.Now()
	s.mu.Lock()
	if now.Sub(s.sentIn) < inputEvery {
		s.mu.Unlock()
		return
	}
	s.sentIn = now
	s.mu.Unlock()
	_ = s.ep.Send(network.Encode(network.Msg{Ver: network.Version, Kind: "input", In: bits}))
}
