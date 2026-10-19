package session

// session_hat.go — chapéu do jogador: o host guarda o mapa (veio
// no join), o convidado recebe pelo welcome/snapshot. Como a cor,
// não tem dono: dois podem usar o mesmo.

// Hat devolve o chapéu do jogador id (0 = sem chapéu).
func (s *Session) Hat(id int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		if h, ok := s.hats[id]; ok {
			return h
		}
		return 0
	}
	if id == s.id {
		return s.hat
	}
	return s.states[id].Hat
}

// hatAt lê o chapéu no mapa do host (chamado com o lock).
func (s *Session) hatAt(id int) int {
	if h, ok := s.hats[id]; ok {
		return h
	}
	return 0
}
