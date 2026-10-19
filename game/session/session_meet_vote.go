package session

import (
	"sort"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_meet_vote.go — voto do convidado (o host guarda) e a
// contagem que decide quem sai: empate ou skip não ejetam ninguém.

// Vote registra o voto de by em target (SkipVote = pular).
func (s *Session) Vote(by, target int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.phase != match.Meeting || !s.voteOpen || s.ghosts[by] {
		return false
	}
	if target != SkipVote && (target == by || s.ghosts[target]) {
		return false
	}
	if s.votes == nil {
		s.votes = map[int]int{}
	}
	s.votes[by] = target
	return true
}

// tallyLocked conta os votos: alvo, se pulou e se era impostor.
func (s *Session) tallyLocked() (int, bool, bool) {
	count := map[int]int{}
	for by, id := range s.votes {
		if !s.ghosts[by] {
			count[id]++
		}
	}
	best, bestN, tie := -1, 0, false
	for id, n := range count {
		switch {
		case n > bestN:
			best, bestN, tie = id, n, false
		case n == bestN:
			tie = true
		}
	}
	if best < 0 || tie || best == SkipVote {
		return -1, true, false
	}
	return best, false, s.plan.Roles[best] == match.Imp
}

// crewWinLocked: nenhum impostor vivo (ejetaram todos).
func (s *Session) crewWinLocked() bool {
	for id, r := range s.plan.Roles {
		if r == match.Imp && !s.ghosts[id] {
			return false
		}
	}
	return true
}

// meetVotesLocked copia (ordenada) a votação atual.
func (s *Session) meetVotesLocked() []network.Vote {
	out := make([]network.Vote, 0, len(s.votes))
	for by, id := range s.votes {
		out = append(out, network.Vote{By: by, ID: id})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].By < out[j].By })
	return out
}
