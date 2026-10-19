package match

import (
	"image"
	"math/rand/v2"

	world "duckling/game/map"
)

// plan.go — o sorteio: quem é impostor, onde cada um nasce e
// quais tarefas cada tripulante leva. Números iniciais (README).

const (
	ShortTasks = 4
	LongTasks  = 2
)

// Build sorteia papéis, spawns e tarefas pra ids (host = 0).
// imps <= 0 = automático: 1 com 4–6 jogadores, 2 com 7+.
func Build(ids []int, spawns []image.Point, tasks []world.Task, imps int) Plan {
	pl := newPlan()
	imps = impostors(len(ids), imps)
	var impIDs []int
	for i, k := range rand.Perm(len(ids)) {
		id := ids[k]
		pl.Roles[id] = Crew
		if i < imps {
			pl.Roles[id] = Imp
			impIDs = append(impIDs, id)
		}
	}
	for _, id := range impIDs {
		for _, o := range impIDs {
			if o != id {
				pl.Mates[id] = append(pl.Mates[id], o)
			}
		}
	}
	for i, id := range ids {
		pl.Spawn[id] = PtAt(spawns, i)
	}
	pl.Tasks = assign(ids, tasks)
	return pl
}

// impostors resolve AUTO (1 até 6, 2 a partir de 7) e limita o
// override pra sempre sobrar tripulante.
func impostors(n, want int) int {
	if want <= 0 {
		if n >= 7 {
			return 2
		}
		return 1
	}
	return min(want, max(1, (n-1)/2))
}

// assign dá 4 curtas + 2 longas por jogador, girando a lista pra
// não repetir o mesmo ponto pra todo mundo. O impostor recebe as
// mesmas (tarefa fake): a lista aparece igual, mas o host não
// conta o progresso dele.
func assign(ids []int, ts []world.Task) map[int][]int {
	var short, long []int
	for i, t := range ts {
		if t.Long {
			long = append(long, i)
		} else {
			short = append(short, i)
		}
	}
	shuffle(short)
	shuffle(long)
	out := map[int][]int{}
	si, li := 0, 0
	for _, id := range ids {
		mine := []int{}
		for k := 0; k < ShortTasks && len(short) > 0; k++ {
			mine = append(mine, short[si%len(short)])
			si++
		}
		for k := 0; k < LongTasks && len(long) > 0; k++ {
			mine = append(mine, long[li%len(long)])
			li++
		}
		out[id] = mine
	}
	return out
}

func shuffle(v []int) {
	rand.Shuffle(len(v), func(i, j int) { v[i], v[j] = v[j], v[i] })
}

// PtAt é o ponto de spawn i (exportado pra testes/preview).
func PtAt(sp []image.Point, i int) Pt {
	p := sp[i%len(sp)]
	return Pt{X: float64(p.X), Y: float64(p.Y)}
}
