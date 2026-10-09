package actor

import (
	"math"
	"math/rand/v2"

	"duckling/src/world"

	"github.com/hajimehoshi/ebiten/v2"
)

// npc.go — pato passeador: vaga aleatório, trombou freia o jogador.

type NPC struct {
	D        *Duck
	dx, dy   float64
	t        float64
	Still    bool
	fx, fy   float64
	flee     float64
}

func NewNPC(img *ebiten.Image, x, y float64, d Def) *NPC {
	n := &NPC{D: New(img, x, y, d)}
	n.retarget()
	return n
}

func (n *NPC) retarget() {
	a := rand.Float64() * 2 * math.Pi
	n.dx, n.dy = math.Cos(a), math.Sin(a)
	n.t = 1 + rand.Float64()*2
}

func (n *NPC) Update(x0, y0, x1, y1 int, solids []world.Solid) {
	// Parado: congela no frame base, mas ainda trombável.
	if n.Still {
		n.D.Anim(0, 0)
		return
	}
	mx, my, sp := n.dx, n.dy, n.D.def.Speed
	if n.flee > 0 {
		// Assustado: corre pra longe por 1s.
		n.flee -= 1.0 / 60.0
		mx, my, sp = n.fx, n.fy, sp*2.5
	} else {
		n.t -= 1.0 / 60.0
		if n.t <= 0 {
			n.retarget()
		}
	}
	n.D.step(mx*sp, my*sp, solids)
	// Encostou no muro, já vira (sem trimilique quica-e-volta).
	if n.D.hx || n.D.hy {
		n.retarget()
	}
	n.D.Px = min(max(n.D.Px, float64(x0)), float64(x1))
	n.D.Py = min(max(n.D.Py, float64(y0)), float64(y1))
}

// Bump freia o jogador ao encostar no chão: tropeça (pulinho),
// meio passo por 1.2s, e o NPC foge assustado (o parado não).
// 1.5s de carência anti-trava; voando alto passa sem encostar.
func (n *NPC) Bump(p *Duck) bool {
	if p.Alt() > 10 || p.bumpT > 0 {
		return false
	}
	dx, dy := p.Px-n.D.Px, p.Py-n.D.Py
	if math.Abs(dx) >= 16 || math.Abs(dy) >= 12 {
		return false
	}
	p.slowT, p.bumpT = 1.2, 1.5
	if p.z == 0 {
		p.vz = 2.2
	}
	if !n.Still {
		if l := math.Hypot(dx, dy); l > 0 {
			n.fx, n.fy, n.flee = -dx/l, -dy/l, 1.0
		}
	}
	return true
}
