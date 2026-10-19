package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_meet_host.go — miudezas da reunião no host: quem pode
// começar, abrir/fechar o estado e a mensagem de fase.

func (s *Session) canMeetLocked(by int) bool {
	return s.phase == match.Playing && !s.ghosts[by]
}

func (s *Session) hasBodyLocked(of int) bool {
	for _, b := range s.bodies {
		if b.ID == of {
			return true
		}
	}
	return false
}

func (s *Session) startMeetLocked(reason string, by, of int) {
	s.meetReason, s.meetBy, s.meetOf = reason, by, of
	s.meetAt, s.voteOpen = time.Now(), false
	s.votes, s.chats = map[int]int{}, nil
	s.phase = match.Meeting
}

func (s *Session) clearMeetLocked() {
	s.meetReason, s.meetBy, s.meetOf = "", 0, 0
	s.voteOpen = false
	s.votes, s.chats = nil, nil
}

func phaseMsg(ph match.Phase) network.Msg {
	return network.Msg{Ver: network.Version, Kind: "phase", Phase: int(ph)}
}
