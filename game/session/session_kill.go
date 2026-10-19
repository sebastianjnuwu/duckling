package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_kill.go — a morte: pedido do convidado, aval do host
// (papel, alvo vivo, cooldown) e o corpo deixado no chão.

const (
	// KillCooldown é o tempo mínimo entre mortes (compartilhado).
	KillCooldown = 25 * time.Second
	// KillReach é o alcance máximo (px) que o host aceita.
	KillReach = 46.0
)

// KillReq é um pedido de morte vindo de um convidado.
type KillReq struct{ By, Target int }

// AskKill enfileira o pedido (o play do host avalia o alcance).
func (s *Session) AskKill(by, target int) {
	s.mu.Lock()
	if len(s.kills) < 8 {
		s.kills = append(s.kills, KillReq{By: by, Target: target})
	}
	s.mu.Unlock()
}

// TakeKills entrega e limpa os pedidos pendentes (host).
func (s *Session) TakeKills() []KillReq {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := s.kills
	s.kills = nil
	return k
}

// CanKill diz se by pode matar target agora (sem mexer em nada).
func (s *Session) CanKill(by, target int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.canKillLocked(by, target)
}

func (s *Session) canKillLocked(by, target int) bool {
	return s.phase == match.Playing && s.plan.Roles[by] == match.Imp &&
		!s.ghosts[by] && !s.ghosts[target] && by != target && s.killCd <= 0
}

// Kill mata target e deixa o corpo em x,y (host). Devolve se matou.
func (s *Session) Kill(by, target int, x, y float64) bool {
	s.mu.Lock()
	if !s.canKillLocked(by, target) {
		s.mu.Unlock()
		return false
	}
	s.ghosts[target] = true
	s.bodies = append(s.bodies, network.Body{ID: target, X: x, Y: y})
	s.killCd, s.cdAt = KillCooldown, time.Now()
	s.recountLocked()
	m := tasksMsg(s.done, s.total)
	end := s.impWinLocked()
	s.mu.Unlock()
	s.broadcast(m)
	if end {
		s.Finish()
	}
	return true
}

// SendKill pede a morte do alvo pro host (cliente).
func (s *Session) SendKill(target int) {
	if s.host || s.ep == nil {
		return
	}
	data := network.Encode(network.Msg{Ver: network.Version, Kind: "kill", ID: target})
	go func() { _ = s.ep.Send(data) }()
}
