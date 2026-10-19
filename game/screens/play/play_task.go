package play

import (
	"duckling/engine/input"
	"duckling/game/match"
	"duckling/game/task"
)

// play_task.go — interação local: acha o ponto mais perto, roda a
// mecânica e avisa o host quando termina.

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

// finishTasks conta quem completou e avisa o host.
func (p *Play) finishTasks() {
	for i, st := range p.ts {
		if st.Done() && !p.done[i] {
			p.done[i] = true
			p.report(i)
			p.reward(i)
		}
	}
}

// report manda a conclusão: o host conta a dele, o cliente avisa.
func (p *Play) report(i int) {
	if p.sess == nil {
		return
	}
	if p.sess.IsHost() {
		p.sess.DoneTask(p.sess.ID(), i)
		return
	}
	p.sess.SendTask(i)
}

// nearestTask é o ponto incompleto mais perto dentro do alcance.
func (p *Play) nearestTask() int {
	best, bd := -1, taskReach
	for i, st := range p.ts {
		if st.Done() || st.Load {
			continue
		}
		if d := p.dist(p.mp.Tasks[i].Center()); d < bd {
			best, bd = i, d
		}
	}
	return best
}

// carrying devolve a tarefa "carry" na mão (-1 = nenhuma).
func (p *Play) carrying() int {
	for i, st := range p.ts {
		if st.Load && !st.Done() {
			return i
		}
	}
	return -1
}
