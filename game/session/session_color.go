package session

import (
	"duckling/engine/network"
	"duckling/game/player"
)

// session_color.go — cor do pato no lobby: o host guarda o mapa e
// trava repetido; o convidado só pede e espera o snapshot.

// Color devolve a cor do jogador id (0 = amarelo).
func (s *Session) Color(id int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		if c, ok := s.colors[id]; ok {
			return c
		}
		return 0
	}
	if id == s.id {
		return s.color
	}
	return s.states[id].Color
}

// Taken diz se a cor c já é de alguém na sala.
func (s *Session) Taken(c int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.host {
		for _, v := range s.colors {
			if v == c {
				return true
			}
		}
		return false
	}
	for _, pl := range s.states {
		if pl.Color == c {
			return true
		}
	}
	return false
}

// PickColor pede a cor (host aplica na hora; convidado manda).
func (s *Session) PickColor(c int) {
	if !actor.ColorValid(c) {
		return
	}
	if s.host {
		s.setColor(0, c)
		return
	}
	if s.ep != nil {
		_ = s.ep.Send(network.Encode(network.Msg{
			Ver: network.Version, Kind: "color", Color: c,
		}))
	}
}

// setColor troca a cor de id se estiver livre (host).
func (s *Session) setColor(id, c int) bool {
	if !actor.ColorValid(c) {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for pid, v := range s.colors {
		if pid != id && v == c {
			return false
		}
	}
	s.colors[id] = c
	return true
}

// freeColorLocked dá a primeira cor livre (0 se lotou).
func (s *Session) freeColorLocked() int {
	var used [actor.ColorCount]bool
	for _, c := range s.colors {
		if actor.ColorValid(c) {
			used[c] = true
		}
	}
	for i, u := range used {
		if !u {
			return i
		}
	}
	return 0
}
