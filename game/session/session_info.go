package session

import (
	"duckling/game/match"
)

// session_info.go — leitura do estado da partida (fase, papel,
// spawn, tarefas e fantasmas). O host é a autoridade; o
// convidado só recebe a própria parte.

func (s *Session) Phase() match.Phase {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.phase
}

// Min é o mínimo de jogadores que o dono configurou.
func (s *Session) Min() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.min
}

// Impostors é o override do dono (0 = automático).
func (s *Session) Impostors() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.imps
}

// Role devolve o papel de id (ok = já sorteado/avisado).
func (s *Session) Role(id int) (match.Role, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.plan.Roles[id]
	return r, ok
}

// Mates são os outros impostores (só o impostor recebe).
func (s *Session) Mates(id int) []int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]int(nil), s.plan.Mates[id]...)
}

// Tasks são os índices das tarefas de id.
func (s *Session) Tasks(id int) []int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]int(nil), s.plan.Tasks[id]...)
}

// Spawn devolve o nascimento sorteado de id.
func (s *Session) Spawn(id int) (float64, float64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.plan.Spawn[id]
	return p.X, p.Y, ok
}

// Ghost diz se id virou fantasma (caiu ou morreu).
func (s *Session) Ghost(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ghosts[id]
}
