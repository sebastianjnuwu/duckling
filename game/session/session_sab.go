package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_sab.go — sabotagem do impostor: apagar a lamparina ou
// fechar as porteiras, com um cooldown compartilhado.

const (
	// SabLamp apaga a luz (o mapa escurece).
	SabLamp = "lamp"
	// SabGate fecha as porteiras (bloqueia as passagens).
	SabGate = "gate"
	// SabCooldown é o tempo mínimo entre sabotagens.
	SabCooldown = 30 * time.Second
	// LampDuration é quanto tempo a lamparina fica apagada.
	LampDuration = 15 * time.Second
	// GateDuration é quanto tempo as porteiras ficam fechadas.
	GateDuration = 30 * time.Second
)

// Sabotage ativa a sabotagem kind se by for impostor vivo (host).
func (s *Session) Sabotage(by int, kind string) bool {
	s.mu.Lock()
	if !s.canSabLocked(by, kind) {
		s.mu.Unlock()
		return false
	}
	s.sabCd = SabCooldown
	switch kind {
	case SabLamp:
		s.lampLeft = LampDuration
	case SabGate:
		s.gateLeft = GateDuration
	}
	s.mu.Unlock()
	s.broadcast(network.Msg{Ver: network.Version, Kind: "sabotage", Sab: kind, ID: by})
	return true
}

func (s *Session) canSabLocked(by int, kind string) bool {
	if s.phase != match.Playing || s.plan.Roles[by] != match.Imp || s.ghosts[by] || s.sabCd > 0 {
		return false
	}
	switch kind {
	case SabLamp:
		return s.lampLeft <= 0
	case SabGate:
		return s.gateLeft <= 0
	}
	return false
}

// SendSab pede a sabotagem pro host (cliente).
func (s *Session) SendSab(kind string) {
	if s.host || s.ep == nil {
		return
	}
	data := network.Encode(network.Msg{Ver: network.Version, Kind: "sabotage", Sab: kind})
	go func() { _ = s.ep.Send(data) }()
}
