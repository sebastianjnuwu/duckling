package session

import (
	"sort"
)

// session_roster.go — a lista da reunião (vivos e mortos) e o
// resultado da ejeção pra tela avisar.

// RosterEntry é um jogador na lista da reunião.
type RosterEntry struct {
	ID    int
	Name  string
	Color int
	Ghost bool
}

// Roster devolve os jogadores ordenados por id.
func (s *Session) Roster() []RosterEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]RosterEntry, 0, len(s.peers)+1)
	if s.host {
		out = append(out, RosterEntry{
			ID: 0, Name: s.name, Color: s.colorAt(0), Ghost: s.ghosts[0],
		})
		for _, p := range s.peers {
			out = append(out, RosterEntry{
				ID: p.id, Name: p.name, Color: s.colorAt(p.id), Ghost: s.ghosts[p.id],
			})
		}
	} else {
		for _, pl := range s.states {
			out = append(out, RosterEntry{
				ID: pl.ID, Name: pl.Name, Color: pl.Color, Ghost: pl.Ghost,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Eject é o resultado da votação.
type Eject struct {
	ID   int
	Skip bool
	Imp  bool
}

// TakeEject entrega e limpa o resultado pendente (uma vez só).
func (s *Session) TakeEject() (Eject, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.ejectDone {
		return Eject{}, false
	}
	s.ejectDone = false
	return Eject{ID: s.ejectID, Skip: s.ejectSkip, Imp: s.ejectImp}, true
}

// colorAt lê a cor no mapa do host (chamado com o lock).
func (s *Session) colorAt(id int) int {
	if c, ok := s.colors[id]; ok {
		return c
	}
	return 0
}
