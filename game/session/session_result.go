package session

import (
	"sort"

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
	s.reveal = s.revealLocked()
	m := network.Msg{
		Ver: network.Version, Kind: "phase", Phase: int(match.End),
		Win: r.Winner, Why: r.Reason, Roles: s.reveal,
	}
	s.mu.Unlock()
	for i := 0; i < 3; i++ {
		s.broadcast(m)
	}
}

// Reveal devolve os papéis de todos (host: do plano; cliente: do fim).
func (s *Session) Reveal() []network.RoleLine {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		return s.revealLocked()
	}
	return append([]network.RoleLine(nil), s.reveal...)
}

// revealLocked monta a lista de papéis ordenada por id (host).
func (s *Session) revealLocked() []network.RoleLine {
	out := make([]network.RoleLine, 0, len(s.plan.Roles))
	for id, r := range s.plan.Roles {
		role := "crew"
		if r == match.Imp {
			role = "imp"
		}
		out = append(out, network.RoleLine{ID: id, Role: role})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
