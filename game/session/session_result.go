package session

import (
	"duckling/engine/network"
	"duckling/game/match"
)

// session_result.go — o desfecho da partida: quem venceu, por quê e
// os números que a tela de fim mostra. O host decide.

// Result devolve o desfecho (vazio se a partida não acabou).
func (s *Session) Result() match.Result {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.result
}

// Kills conta as mortes da partida.
func (s *Session) Kills() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.killsN
}

// LastVotes devolve os votos da última reunião (quem votou em quem).
func (s *Session) LastVotes() []network.Vote {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]network.Vote(nil), s.lastVotes...)
}

// Stop encerra sem vencedor (o dono desistiu da partida).
func (s *Session) Stop() { s.finish(match.Result{}) }

// win monta o resultado a partir da mesa (host).
func (s *Session) win() match.Result {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.winLocked()
}

// winLocked: sem impostores vivos = tripulantes; tarefas 100% =
// tripulantes; impostores em número igual ou maior = impostores.
func (s *Session) winLocked() match.Result {
	imps, crew := 0, 0
	for id, r := range s.plan.Roles {
		if s.ghosts[id] {
			continue
		}
		if r == match.Imp {
			imps++
		} else {
			crew++
		}
	}
	switch {
	case imps == 0:
		return match.Result{Winner: match.WinCrew, Reason: match.WhyEject}
	case s.total > 0 && s.done >= s.total:
		return match.Result{Winner: match.WinCrew, Reason: match.WhyTasks}
	case imps >= crew:
		return match.Result{Winner: match.WinImp, Reason: match.WhyParity}
	}
	return match.Result{}
}

// finish fecha a partida e anuncia o resultado (host).
func (s *Session) finish(r match.Result) {
	s.mu.Lock()
	s.phase, s.result = match.End, r
	s.mu.Unlock()
	m := network.Msg{
		Ver: network.Version, Kind: "phase", Phase: int(match.End),
		Win: r.Winner, Why: r.Reason,
	}
	for i := 0; i < 3; i++ {
		s.broadcast(m)
	}
}
