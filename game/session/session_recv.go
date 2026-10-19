package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/player"
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

// Me devolve o meu estado no último snapshot (ok = achado).
func (s *Session) Me() (network.Player, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.states[s.id]
	return p, ok
}

// SendInput manda a entrada segurada + as bordas reenviadas (30/s).
func (s *Session) SendInput(in actor.In) {
	if s.host || s.ep == nil {
		return
	}
	now := time.Now()
	s.mu.Lock()
	if now.Sub(s.sentIn) < inputEvery {
		s.mu.Unlock()
		return
	}
	s.sentIn = now
	edge := in.Edges() | s.edgeLast
	s.edgeLast = in.Edges()
	s.mu.Unlock()
	_ = s.ep.Send(network.Encode(network.Msg{
		Ver: network.Version, Kind: "input", In: in.Encode(), Edge: edge,
	}))
}
