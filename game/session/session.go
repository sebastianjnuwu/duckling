package session

import (
	"sync"
	"time"
)

// session.go — partida na LAN: o host simula e distribui, o
// cliente manda input e interpola os remotos. Sem sessão = solo.

const (
	stateEvery  = 66 * time.Millisecond // 15 snapshots/s
	inputEvery  = 33 * time.Millisecond // 30 inputs/s
	peerTimeout = 5 * time.Second       // 5s mudo = caiu
)

var (
	regMu   sync.Mutex
	current *Session
)

// Set troca a sessão atual (nil = solo); a antiga é fechada.
func Set(s *Session) {
	regMu.Lock()
	old := current
	current = s
	regMu.Unlock()
	if old != nil && old != s {
		old.Close()
	}
}

// Get devolve a sessão atual (nil = solo).
func Get() *Session {
	regMu.Lock()
	defer regMu.Unlock()
	return current
}

// IsHost diz se esta máquina é a dona da sala.
func (s *Session) IsHost() bool { return s.host }

// ID é o número do jogador local (host = 0).
func (s *Session) ID() int { return s.id }

// Name é o nome do jogador local.
func (s *Session) Name() string { return s.name }

// Room é o nome da sala.
func (s *Session) Room() string { return s.room }

// Max é o teto de jogadores.
func (s *Session) Max() int { return s.max }

// errString é um erro que carrega uma chave i18n.
type errString string

func (e errString) Error() string { return string(e) }
