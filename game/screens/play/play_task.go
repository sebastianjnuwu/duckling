package play

import (
	"duckling/engine/input"
	"duckling/game/match"
	"duckling/game/task"
)

// play_task.go — interação local: qual ponto está no alcance e o
// que a tecla segurada faz nele (a mecânica pura é de game/task).

// syncTasks refaz o estado local quando a lista muda (o "role"
// pode chegar antes ou depois do "start").
func (p *Play) syncTasks() {
	var ids []int
	if p.sess != nil && p.phase != match.Lobby {
		ids = p.sess.Tasks(p.sess.ID())
	}
	if sameIDs(ids, p.tids) {
		return
	}
	p.tids = append([]int(nil), ids...)
	p.ts = map[int]*task.State{}
	p.done = map[int]bool{}
	for _, i := range ids {
		if i >= 0 && i < len(p.mp.Tasks) {
			p.ts[i] = task.New(p.mp.Tasks[i])
		}
	}
}

func sameIDs(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// taskUpdate roda a interação do jogador local (jogando e vivo).
func (p *Play) taskUpdate(dt float64) {
	if p.sess == nil || p.phase != match.Playing || p.ghost() {
		return
	}
	held := input.Held(input.Interact)
	pressed := input.JustPressed(input.Interact)

	// Carregando: só falta entregar no destino.
	if i := p.carrying(); i >= 0 {
		if pressed && p.near(p.mp.Tasks[i].Target()) {
			p.ts[i].Tap(true)
		}
		p.finishTasks()
		return
	}

	idx := p.nearestTask()
	for i, st := range p.ts {
		st.Hold(dt, held && i == idx)
	}
	if idx >= 0 && pressed {
		p.ts[idx].Tap(false)
	}
	p.finishTasks()
}
