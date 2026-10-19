package session

import (
	"sort"
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_match.go — o dono abre e fecha a partida: trava a sala,
// manda cada papel só pro dono dele e volta todo mundo ao lobby.

// PlanIDs devolve os ids da sala (host = 0), ordenados.
func (s *Session) PlanIDs() []int {
	s.mu.Lock()
	defer s.mu.Unlock()
	ids := make([]int, 0, len(s.peers)+1)
	ids = append(ids, 0)
	for _, p := range s.peers {
		ids = append(ids, p.id)
	}
	sort.Ints(ids)
	return ids
}

// Start trava a sala e distribui o plano (host).
func (s *Session) Start(plan match.Plan) {
	s.mu.Lock()
	s.plan, s.phase = plan, match.Playing
	s.ghosts = map[int]bool{}
	s.tasks = map[int]map[int]bool{}
	s.bodies, s.kills = nil, nil
	s.killCd, s.cdAt = 0, time.Time{}
	s.sabCd, s.lampLeft, s.gateLeft = 0, 0, 0
	s.clearMeetLocked()
	s.calls = map[int]bool{}
	s.recountLocked()
	done, total := s.done, s.total
	type target struct {
		addr string
		id   int
	}
	ts := make([]target, 0, len(s.peers))
	for addr, p := range s.peers {
		ts = append(ts, target{addr, p.id})
	}
	s.mu.Unlock()

	for _, t := range ts {
		at := plan.Spawn[t.id]
		role := "crew"
		if plan.Roles[t.id] == match.Imp {
			role = "imp"
		}
		_ = s.ep.SendTo(t.addr, network.Encode(network.Msg{
			Ver: network.Version, Kind: "role", Phase: int(match.Playing),
			Role: role, Mates: plan.Mates[t.id], Tasks: plan.Tasks[t.id],
			X: at.X, Y: at.Y,
		}))
	}
	s.broadcast(network.Msg{Ver: network.Version, Kind: "start", Phase: int(match.Playing)})
	s.broadcast(tasksMsg(done, total))
}

// Finish encerra a partida (host) e abre a tela de fim.
func (s *Session) Finish() {
	s.mu.Lock()
	s.phase = match.End
	s.mu.Unlock()
	s.broadcast(network.Msg{Ver: network.Version, Kind: "phase", Phase: int(match.End)})
}

// Reset volta todo mundo pro lobby e limpa papéis/fantasmas.
func (s *Session) Reset() {
	s.mu.Lock()
	s.phase = match.Lobby
	s.plan.Clear()
	s.ghosts = map[int]bool{}
	s.tasks = map[int]map[int]bool{}
	s.done, s.total = 0, 0
	s.bodies, s.kills = nil, nil
	s.killCd, s.cdAt = 0, time.Time{}
	s.sabCd, s.lampLeft, s.gateLeft = 0, 0, 0
	s.clearMeetLocked()
	s.calls = map[int]bool{}
	s.mu.Unlock()
	s.broadcast(network.Msg{Ver: network.Version, Kind: "phase", Phase: int(match.Lobby)})
}
