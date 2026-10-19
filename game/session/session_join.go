package session

import (
	"time"

	"duckling/engine/network"
	"duckling/engine/save"
	"duckling/game/match"
	"duckling/game/player"
)

// session_join.go — a porta de entrada: valida versão e senha,
// limita a lotação, fecha depois de começar e devolve um id.

func (s *Session) hostJoin(addr string, m network.Msg) {
	reason := ""
	switch {
	case m.Ver != network.Version:
		reason = "net.badversion"
	case s.pass != "" && m.Pass != s.pass:
		reason = "net.badpass"
	}
	s.mu.Lock()
	if reason == "" && s.phase != match.Lobby {
		reason = "net.started"
	}
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
		if s.colors == nil {
			s.colors = map[int]int{0: int(actor.Yellow)}
		}
		if s.hats == nil {
			s.hats = map[int]int{0: save.Hat()}
		}
		p = &peer{id: s.nextID, name: clip(m.Name)}
		s.nextID++
		s.peers[addr] = p
		s.colors[p.id] = s.freeColorLocked()
		s.hats[p.id] = 0
		if actor.HatValid(m.Hat) {
			s.hats[p.id] = m.Hat
		}
		s.events = append(s.events, "net.joined")
	}
	p.last = time.Now()
	id := p.id
	min := s.min
	color := s.colors[id]
	hat := s.hats[id]
	s.mu.Unlock()
	_ = s.ep.SendTo(addr, network.Encode(network.Msg{
		Ver: network.Version, Kind: "welcome", ID: id, Room: s.room,
		Max: s.max, Min: min, Color: color, Hat: hat,
	}))
}

func (s *Session) drop(addr string) {
	s.mu.Lock()
	if p, ok := s.peers[addr]; ok {
		delete(s.peers, addr)
		delete(s.colors, p.id)
		delete(s.hats, p.id)
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
