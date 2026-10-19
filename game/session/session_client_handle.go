package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_client_handle.go — o que o convidado faz com cada
// mensagem: snapshot, papel, fase, tarefas e recusa.

func (s *Session) clientHandle(m network.Msg) {
	switch m.Kind {
	case "state":
		fresh := make(map[int]network.Player, len(m.Players))
		for _, p := range m.Players {
			fresh[p.ID] = p
		}
		s.mu.Lock()
		s.states, s.lastRecv = fresh, time.Now()
		s.cdMS, s.bodies = m.Cd, m.Bodies
		s.sabMS, s.lampMS, s.gateMS = m.SabCd, m.Lamp, m.Gate
		// O snapshot também carrega a fase: se um "phase" se perder,
		// reunião/fim/volta chegam do mesmo jeito (lobby fica com a
		// mensagem explícita pra não voltar por pacote atrasado).
		if m.Phase != int(match.Lobby) && s.phase != match.End {
			s.phase = match.Phase(m.Phase)
		}
		if m.Phase == int(match.End) {
			s.result = match.Result{Winner: m.Win, Reason: m.Why}
			s.reveal = m.Roles
		}
		s.meetMS, s.meetOpen = m.MeetLeft, m.MeetOpen
		s.meetReason, s.meetBy, s.meetOf = m.Meet, m.MeetBy, m.MeetOf
		s.meetAll = m.Votes
		s.chats = m.Chats
		if p, ok := fresh[s.id]; ok {
			s.color = p.Color
		}
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
		if m.Win != "" || m.Phase == int(match.End) {
			s.result = match.Result{Winner: m.Win, Reason: m.Why}
			s.reveal = m.Roles
		}
		if m.Meet != "" {
			s.meetMS, s.meetOpen = m.MeetLeft, m.MeetOpen
			s.meetReason, s.meetBy, s.meetOf = m.Meet, m.MeetBy, m.MeetOf
			s.meetAll = m.Votes
		}
		if s.phase == match.Lobby {
			s.plan.Clear()
			s.done, s.total = 0, 0
			s.bodies, s.cdMS, s.killCd = nil, 0, 0
			s.sabMS, s.lampMS, s.gateMS = 0, 0, 0
			s.meetMS, s.meetOpen, s.meetAll = 0, false, nil
			s.meetReason, s.meetBy, s.meetOf = "", 0, 0
			s.result, s.reveal = match.Result{}, nil
		}
		s.mu.Unlock()
	case "tasks":
		s.mu.Lock()
		s.done, s.total = m.Done, m.Total
		s.mu.Unlock()
	case "eject":
		s.mu.Lock()
		s.ejectID, s.ejectSkip, s.ejectImp = m.ID, m.Skip, m.Imp
		s.ejectDone = true
		s.mu.Unlock()
	case "deny":
		s.mu.Lock()
		s.fatal = m.Reason
		s.mu.Unlock()
	}
}
