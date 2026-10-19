package session

import (
	"time"

	"duckling/engine/network"
)

// session_meet_send.go — o convidado pede: reportar o corpo, tocar o
// botão de emergência e votar (3 envios, porque UDP perde).

// SendReport avisa o host que achamos o corpo de of.
func (s *Session) SendReport(of int) {
	s.sendThrice(network.Msg{Kind: "report", ID: of})
}

// SendBell toca o botão de emergência.
func (s *Session) SendBell() {
	s.sendThrice(network.Msg{Kind: "call"})
}

// SendVote manda o voto (SkipVote = pular).
func (s *Session) SendVote(target int) {
	s.sendThrice(network.Msg{Kind: "vote", ID: target})
}

func (s *Session) sendThrice(m network.Msg) {
	if s.host || s.ep == nil {
		return
	}
	m.Ver = network.Version
	data := network.Encode(m)
	go func() {
		for i := 0; i < 3; i++ {
			_ = s.ep.Send(data)
			time.Sleep(80 * time.Millisecond)
		}
	}()
}
