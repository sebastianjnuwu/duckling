package match

// match.go — regras da partida: fases, papéis e o plano que o
// host sorteia. Puro: não conhece rede nem tela.

// Phase é a fase da partida (o host é a autoridade).
type Phase int

const (
	Lobby Phase = iota
	Playing
	Meeting
	End
)

// Role é o papel de um pato.
type Role int

const (
	Crew Role = iota
	Imp
)

// Pt é uma posição de mundo.
type Pt struct{ X, Y float64 }

// Plan é o sorteio do host: papel, spawn e tarefas de cada um.
type Plan struct {
	Roles map[int]Role
	Mates map[int][]int
	Tasks map[int][]int
	Spawn map[int]Pt
}

func newPlan() Plan {
	return Plan{
		Roles: map[int]Role{}, Mates: map[int][]int{},
		Tasks: map[int][]int{}, Spawn: map[int]Pt{},
	}
}

// Put grava a parte de um jogador (o convidado usa no "role").
func (p *Plan) Put(id int, role Role, mates, tasks []int, at Pt) {
	if p.Roles == nil {
		*p = newPlan()
	}
	p.Roles[id], p.Mates[id] = role, mates
	p.Tasks[id], p.Spawn[id] = tasks, at
}

// Clear zera o plano (volta pro lobby).
func (p *Plan) Clear() { *p = newPlan() }
