package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/player"
)

// session_pub.go — o host distribui: snapshot dos patos (15/s)
// e leitura do input que cada cliente mandou.

// Publish manda o snapshot pros clientes (no máximo 15/s).
func (s *Session) Publish(ps []network.Player) {
	now := time.Now()
	s.mu.Lock()
	if now.Sub(s.sent) < stateEvery {
		s.mu.Unlock()
		return
	}
	s.sent = now
	s.mu.Unlock()
	s.broadcast(network.Msg{Ver: network.Version, Kind: "state", Players: ps})
}

// broadcast manda a mesma mensagem pra cada convidado.
func (s *Session) broadcast(m network.Msg) {
	s.mu.Lock()
	addrs := make([]string, 0, len(s.peers))
	for addr := range s.peers {
		addrs = append(addrs, addr)
	}
	s.mu.Unlock()
	data := network.Encode(m)
	for _, addr := range addrs {
		_ = s.ep.SendTo(addr, data)
	}
}

// HostInputs devolve a entrada de cada cliente, com a borda de
// pulo/voo já derivada e consumida.
func (s *Session) HostInputs() map[int]actor.In {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[int]actor.In, len(s.peers))
	for _, p := range s.peers {
		out[p.id] = actor.DecodeIn(p.in, p.edge)
		p.edge = 0
	}
	return out
}
