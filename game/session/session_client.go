package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_client.go — o convidado: pede entrada (esperando a
// resposta na hora), manda o input e guarda o último estado.

// Join conecta (bloqueia até welcome/deny) e devolve a sessão.
// addr é "ip:porta" do host. Erro traz chave i18n.
func Join(addr, name, pass string) (*Session, error) {
	ep, err := network.DialHost(addr)
	if err != nil {
		return nil, errString("net.fail")
	}
	s := &Session{
		name: name, pass: pass, ep: ep, stop: make(chan struct{}),
		states: map[int]network.Player{}, lastRecv: time.Now(),
	}
	_ = s.ep.Send(network.Encode(network.Msg{Ver: network.Version, Kind: "join", Name: name, Pass: pass}))
	deadline := time.After(1200 * time.Millisecond)
	for {
		select {
		case <-deadline:
			s.ep.Close()
			return nil, errString("net.timeout")
		case pkt := <-ep.C():
			m, err := network.Decode(pkt.Data)
			if err != nil {
				continue
			}
			switch m.Kind {
			case "welcome":
				s.id, s.room, s.max = m.ID, m.Room, m.Max
				s.min = m.Min
				s.joined, s.lastRecv = true, time.Now()
				go s.clientLoop()
				return s, nil
			case "deny":
				s.ep.Close()
				return nil, errString(m.Reason)
			}
		}
	}
}

func (s *Session) clientLoop() {
	for {
		select {
		case <-s.stop:
			return
		case pkt := <-s.ep.C():
			m, err := network.Decode(pkt.Data)
			if err != nil {
				continue
			}
			s.clientHandle(m)
		}
	}
}

func (s *Session) clientHandle(m network.Msg) {
	switch m.Kind {
	case "state":
		fresh := make(map[int]network.Player, len(m.Players))
		for _, p := range m.Players {
			fresh[p.ID] = p
		}
		s.mu.Lock()
		s.states, s.lastRecv = fresh, time.Now()
		s.mu.Unlock()
	case "role":
		role := match.Crew
		if m.Role == "imp" {
			role = match.Imp
		}
		s.mu.Lock()
		s.plan.Put(s.id, role, m.Mates, m.Tasks, match.Pt{X: m.X, Y: m.Y})
		s.phase = match.Phase(m.Phase)
		s.mu.Unlock()
	case "start", "phase":
		s.mu.Lock()
		s.phase = match.Phase(m.Phase)
		if s.phase == match.Lobby {
			s.plan.Clear()
			s.done, s.total = 0, 0
		}
		s.mu.Unlock()
	case "tasks":
		s.mu.Lock()
		s.done, s.total = m.Done, m.Total
		s.mu.Unlock()
	case "deny":
		s.mu.Lock()
		s.fatal = m.Reason
		s.mu.Unlock()
	}
}
