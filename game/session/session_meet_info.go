package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_meet_info.go — a reunião como a tela vê: motivo, who,
// relógio da fase atual e a votação.

// NoVote é o retorno de MyVote quando ainda não votei.
const NoVote = -2

// MeetInfo é o retrato da reunião (Left em segundos da fase atual).
type MeetInfo struct {
	Active bool
	Reason string
	By, Of int
	Left   float64
	Open   bool
	Votes  []network.Vote
}

// Meet devolve o estado da reunião (Active=false fora dela).
func (s *Session) Meet() MeetInfo {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.phase != match.Meeting {
		return MeetInfo{}
	}
	m := MeetInfo{Active: true, Reason: s.meetReason, By: s.meetBy, Of: s.meetOf}
	if s.host {
		el := time.Since(s.meetAt)
		m.Open = s.voteOpen
		m.Left = (DiscussTime - el).Seconds()
		if s.voteOpen {
			m.Left = (DiscussTime + VoteTime - el).Seconds()
		}
		m.Votes = s.meetVotesLocked()
	} else {
		m.Open, m.Left = s.meetOpen, float64(s.meetMS)/1000
		m.Votes = append([]network.Vote(nil), s.meetAll...)
	}
	if m.Left < 0 {
		m.Left = 0
	}
	return m
}

// MyVote devolve em quem votei (NoVote = ainda não votei).
func (s *Session) MyVote() int {
	m := s.Meet()
	for _, v := range m.Votes {
		if v.By == s.id {
			return v.ID
		}
	}
	return NoVote
}
