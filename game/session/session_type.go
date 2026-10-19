package session

import (
	"sync"
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_type.go — o estado da sessão e do convidado (o resto do
// pacote só usa; a struct saiu do session.go pra caber no limite).

type peer struct {
	id   int
	name string
	in   uint8 // bits segurados
	prev uint8 // bits do pacote anterior
	edge uint8 // borda pendente (consumida no HostInputs)
	last time.Time
	// Filtro de borda repetida (reenvio do cliente).
	jumpAt, flyAt time.Time
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
	sabCd    time.Duration
	lampLeft time.Duration
	gateLeft time.Duration
	sabMS    int
	lampMS   int
	gateMS   int
	edgeLast uint8
}
