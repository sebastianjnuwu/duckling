package session

import (
	"time"

	"duckling/engine/network"
	"duckling/engine/save"
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
	_ = s.ep.Send(network.Encode(network.Msg{Ver: network.Version, Kind: "join", Name: name, Pass: pass, Hat: save.Hat()}))
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
				s.min, s.color, s.hat = m.Min, m.Color, m.Hat
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
