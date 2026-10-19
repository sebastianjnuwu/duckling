package session

import (
	"time"

	"duckling/engine/network"
)

// session_host.go — o dono da sala: aceita join, guarda o input
// dos clientes, simula todo mundo (no play) e distribui snapshot.

// Host abre a sala e passa a anunciar na LAN.
func Host(name, room, pass string, min, max, imps int) (*Session, error) {
	ep, err := network.ListenHost()
	if err != nil {
		return nil, err
	}
	s := &Session{
		host: true, name: name, room: room, pass: pass,
		min: min, max: max, imps: imps,
		ep: ep, stop: make(chan struct{}), nextID: 1,
		peers: map[string]*peer{}, ghosts: map[int]bool{},
		tasks: map[int]map[int]bool{},
	}
	go s.announce()
	go s.hostLoop()
	return s, nil
}

func (s *Session) announce() {
	_ = network.AnnounceLoop(s.room, s.Count, s.max, s.pass != "", s.stop)
}

// Count é quantos estão na sala (host + clientes).
func (s *Session) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.peers) + 1
}

func (s *Session) hostLoop() {
	for {
		select {
		case <-s.stop:
			return
		case pkt := <-s.ep.C():
			m, err := network.Decode(pkt.Data)
			if err != nil {
				continue
			}
			s.hostHandle(pkt.From, m)
		}
	}
}

func (s *Session) hostHandle(addr string, m network.Msg) {
	switch m.Kind {
	case "join":
		s.hostJoin(addr, m)
	case "input":
		s.hostInput(addr, m)
	case "leave":
		s.drop(addr)
	case "task":
		s.mu.Lock()
		id := 0
		if p := s.peers[addr]; p != nil {
			id = p.id
			p.last = time.Now()
		}
		s.mu.Unlock()
		if id > 0 {
			s.DoneTask(id, m.Task)
		}
	case "kill":
		s.mu.Lock()
		id := 0
		if p := s.peers[addr]; p != nil {
			id = p.id
			p.last = time.Now()
		}
		s.mu.Unlock()
		if id > 0 {
			s.AskKill(id, m.ID)
		}
	case "sabotage":
		s.mu.Lock()
		id := 0
		if p := s.peers[addr]; p != nil {
			id = p.id
			p.last = time.Now()
		}
		s.mu.Unlock()
		if id > 0 {
			s.Sabotage(id, m.Sab)
		}
	}
}
