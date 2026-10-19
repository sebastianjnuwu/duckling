package session

import (
	"sync"
	"time"

	"duckling/engine/network"
	"duckling/game/match"
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

type peer struct {
	id   int
	name string
	in   uint8 // bits segurados
	prev uint8 // bits do pacote anterior
	edge uint8 // borda pendente (consumida no HostInputs)
	last time.Time
}

type Session struct {
	host bool
	name string
	room string
	pass string
	min  int
	max  int
	imps int
	id   int
	ep   *network.Endpoint
	stop chan struct{}
	once sync.Once

	mu       sync.Mutex
	nextID   int
	peers    map[string]*peer
	states   map[int]network.Player
	events   []string
	fatal    string
	lastRecv time.Time
	sent     time.Time
	sentIn   time.Time
	joined   bool
	phase    match.Phase
	plan     match.Plan
	ghosts   map[int]bool
	tasks    map[int]map[int]bool
	done     int
	total    int
	bodies   []network.Body
	kills    []KillReq
	killCd   time.Duration
	cdAt     time.Time
	cdMS     int
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
