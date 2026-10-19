package session

import (
	"time"

	"duckling/engine/network"
)

// session_handle.go — o que o host faz com cada mensagem que chega.

func (s *Session) hostHandle(addr string, m network.Msg) {
	switch m.Kind {
	case "join":
		s.hostJoin(addr, m)
	case "input":
		s.hostInput(addr, m)
	case "leave":
		s.drop(addr)
	case "task":
		if id := s.peerID(addr); id > 0 {
			s.DoneTask(id, m.Task)
		}
	case "kill":
		if id := s.peerID(addr); id > 0 {
			s.AskKill(id, m.ID)
		}
	case "color":
		if id := s.peerID(addr); id > 0 {
			s.setColor(id, m.Color)
		}
	case "report":
		if id := s.peerID(addr); id > 0 {
			s.Report(id, m.ID)
		}
	case "call":
		if id := s.peerID(addr); id > 0 {
			s.CallBell(id)
		}
	case "vote":
		if id := s.peerID(addr); id > 0 {
			s.Vote(id, m.ID)
		}
	case "sabotage":
		if id := s.peerID(addr); id > 0 {
			s.Sabotage(id, m.Sab)
		}
	}
}

// peerID acha o remetente e marca que ele está vivo.
func (s *Session) peerID(addr string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := s.peers[addr]
	if p == nil {
		return 0
	}
	p.last = time.Now()
	return p.id
}
