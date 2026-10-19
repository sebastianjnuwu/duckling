package task

import world "duckling/game/map"

// task.go — mecânica das tarefas (sem desenho, sem rede):
//   hold : segurar enche (soltar pausa)
//   order: segurar enche (soltar zera)
//   carry: aperta na origem, aperta de novo no destino

// State é o progresso local de uma tarefa.
type State struct {
	Kind string
	Dur  float64
	Prog float64 // 0..1
	Load bool    // carry: carga na mão
}

// New monta o estado a partir do ponto do mapa.
func New(t world.Task) *State {
	return &State{Kind: t.Kind, Dur: t.Time}
}

// Hold avança hold/order com a tecla segurada por dt segundos.
func (s *State) Hold(dt float64, held bool) {
	if s.Kind == world.KindCarry || s.Done() {
		return
	}
	if !held {
		if s.Kind == world.KindOrder {
			s.Prog = 0
		}
		return
	}
	s.Prog = min(1, s.Prog+dt/dur(s.Dur))
}

// Tap trata o aperto do carry: pega na origem e entrega no
// destino (atDest = pato em cima do destino).
func (s *State) Tap(atDest bool) {
	if s.Kind != world.KindCarry || s.Done() {
		return
	}
	if !s.Load {
		s.Load = true
		return
	}
	if atDest {
		s.Prog = 1
	}
}

// Done diz se a tarefa já terminou.
func (s *State) Done() bool { return s.Prog >= 1 }

// dur garante pelo menos 1s de barra (mapa sem "time").
func dur(d float64) float64 {
	if d < 1 {
		return 1
	}
	return d
}
