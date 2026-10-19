package play

import (
	"image"
	"math"
)

// play_task_geo.go — alcance, distância e a escolha do ponto.

const taskReach = 46.0 // px do pé do pato até o centro do ponto

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

// near diz se o pé do pato está perto do ponto.
func (p *Play) near(pt image.Point) bool {
	return p.dist(pt) <= taskReach
}

func (p *Play) dist(pt image.Point) float64 {
	dx := p.d.Px - float64(pt.X)
	dy := p.d.Py - float64(pt.Y)
	return math.Hypot(dx, dy)
}

// ghost diz se o pato local virou fantasma (o host só guarda os
// remotos; o cliente sabe pelo snapshot).
func (p *Play) ghost() bool {
	if p.sess == nil {
		return false
	}
	if pl, ok := p.sess.Me(); ok {
		return pl.Ghost
	}
	return p.sess.Ghost(p.sess.ID())
}
