package session

import (
	"time"

	"duckling/engine/network"
)

// session_join.go — a porta de entrada: valida versão e senha,
// limita a lotação e devolve um id pro convidado.

func (s *Session) hostJoin(addr string, m network.Msg) {
	reason := ""
	switch {
	case m.Ver != network.Version:
		reason = "net.badversion"
	case s.pass != "" && m.Pass != s.pass:
		reason = "net.badpass"
	}
	s.mu.Lock()
	if reason == "" && len(s.peers)+1 >= s.max {
		if _, ok := s.peers[addr]; !ok {
			reason = "net.full"
		}
	}
	if reason != "" {
		s.mu.Unlock()
		_ = s.ep.SendTo(addr, network.Encode(network.Msg{Ver: network.Version, Kind: "deny", Reason: reason}))
		return
	}
	p := s.peers[addr]
	if p == nil {
		p = &peer{id: s.nextID, name: clip(m.Name)}
		s.nextID++
		s.peers[addr] = p
		s.events = append(s.events, "net.joined")
	}
	p.last = time.Now()
	id := p.id
	s.mu.Unlock()
	_ = s.ep.SendTo(addr, network.Encode(network.Msg{
		Ver: network.Version, Kind: "welcome", ID: id, Room: s.room, Max: s.max,
	}))
}

func (s *Session) drop(addr string) {
	s.mu.Lock()
	if _, ok := s.peers[addr]; ok {
		delete(s.peers, addr)
		s.events = append(s.events, "net.left")
	}
	s.mu.Unlock()
}

// PeerName devolve o nome do jogador id (host: 0 é o dono).
func (s *Session) PeerName(id int) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if id == 0 {
		return s.name
	}
	for _, p := range s.peers {
		if p.id == id {
			return p.name
		}
	}
	return ""
}

// clip corta o nome no limite da UI (3-8) e evita nome vazio.
func clip(s string) string {
	r := []rune(s)
	if len(r) > 8 {
		r = r[:8]
	}
	if len(r) < 3 {
		return "DUCK"
	}
	return string(r)
}
