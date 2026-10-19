package session

import (
	"time"

	"duckling/engine/network"
	"duckling/game/match"
)

// session_task.go — progresso das tarefas: o host conta e distribui
// concluídas/total (a barra global); o convidado só avisa quando
// termina (3 envios, porque UDP perde).

// DoneTask marca a tarefa idx de id (host). Ignora fantasma,
// impostor e tarefa que não é dele. Devolve true se contou.
func (s *Session) DoneTask(id, idx int) bool {
	s.mu.Lock()
	if s.plan.Roles[id] != match.Crew || s.ghosts[id] || !hasTask(s.plan.Tasks[id], idx) {
		s.mu.Unlock()
		return false
	}
	if s.tasks[id] == nil {
		s.tasks[id] = map[int]bool{}
	}
	if s.tasks[id][idx] {
		s.mu.Unlock()
		return false
	}
	s.tasks[id][idx] = true
	s.recountLocked()
	m := tasksMsg(s.done, s.total)
	s.mu.Unlock()
	s.broadcast(m)
	return true
}

// SendTask avisa o host que a tarefa idx terminou (cliente).
func (s *Session) SendTask(idx int) {
	if s.host || s.ep == nil {
		return
	}
	data := network.Encode(network.Msg{Ver: network.Version, Kind: "task", Task: idx})
	go func() {
		for i := 0; i < 3; i++ {
			_ = s.ep.Send(data)
			time.Sleep(120 * time.Millisecond)
		}
	}()
}

// GlobalTasks devolve concluídas/total da última contagem.
func (s *Session) GlobalTasks() (int, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.done, s.total
}

// refreshTasks recalcula e distribui o total (host).
func (s *Session) refreshTasks() {
	s.mu.Lock()
	s.recountLocked()
	m := tasksMsg(s.done, s.total)
	s.mu.Unlock()
	s.broadcast(m)
}

// recountLocked conta só tripulantes vivos (fantasma sai do total).
func (s *Session) recountLocked() {
	done, total := 0, 0
	for id, r := range s.plan.Roles {
		if r != match.Crew || s.ghosts[id] {
			continue
		}
		for _, idx := range s.plan.Tasks[id] {
			total++
			if s.tasks[id][idx] {
				done++
			}
		}
	}
	s.done, s.total = done, total
}

func tasksMsg(done, total int) network.Msg {
	return network.Msg{Ver: network.Version, Kind: "tasks", Done: done, Total: total}
}

func hasTask(ids []int, idx int) bool {
	for _, i := range ids {
		if i == idx {
			return true
		}
	}
	return false
}
