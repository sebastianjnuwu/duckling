package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_meet.go — a reunião: corpo achado ou botão de emergência
// congela o mundo, abre a discussão, depois o voto, e ejetar revela
// se era o impostor. Tudo decidido no host.

const (
	// DiscussTime é a discussão antes do voto abrir.
	DiscussTime = 60 * time.Second
	// VoteTime é a janela de votação depois da discussão.
	VoteTime = 30 * time.Second
	// SkipVote é o voto de pular (ninguém sai).
	SkipVote = -1
)

// Report começa a reunião por corpo achado (host): by vivo e of
// tem que ser um corpo caído.
func (s *Session) Report(by, of int) bool {
	s.mu.Lock()
	if !s.canMeetLocked(by) || !s.hasBodyLocked(of) {
		s.mu.Unlock()
		return false
	}
	s.startMeetLocked("body", by, of)
	s.mu.Unlock()
	s.broadcast(phaseMsg(match.Meeting))
	return true
}

// CallBell começa a reunião pelo botão de emergência: um por
// jogador por partida.
func (s *Session) CallBell(by int) bool {
	s.mu.Lock()
	if !s.canMeetLocked(by) || s.calls[by] {
		s.mu.Unlock()
		return false
	}
	if s.calls == nil {
		s.calls = map[int]bool{}
	}
	s.calls[by] = true
	s.startMeetLocked("bell", by, -1)
	s.mu.Unlock()
	s.broadcast(phaseMsg(match.Meeting))
	return true
}

// MeetStep toca o relógio da reunião e resolve no fim (host).
func (s *Session) MeetStep() {
	s.mu.Lock()
	if s.phase != match.Meeting {
		s.mu.Unlock()
		return
	}
	el := time.Since(s.meetAt)
	if !s.voteOpen && el >= DiscussTime {
		s.voteOpen = true
		s.mu.Unlock()
		return
	}
	if !s.voteOpen || el < DiscussTime+VoteTime {
		s.mu.Unlock()
		return
	}
	id, skip, imp := s.tallyLocked()
	s.clearMeetLocked()
	s.phase = match.Playing
	end := false
	if !skip {
		s.ghosts[id] = true
		s.recountLocked()
		end = s.impWinLocked() || s.crewWinLocked()
	}
	s.ejectID, s.ejectSkip, s.ejectImp, s.ejectDone = id, skip, imp, true
	s.mu.Unlock()
	s.broadcast(network.Msg{
		Ver: network.Version, Kind: "eject", ID: id, Skip: skip, Imp: imp,
	})
	if end {
		s.Finish()
		return
	}
	s.broadcast(phaseMsg(match.Playing))
}
