package play

import (
	"image"
	"math"
)

// play_task_geo.go — alcance, distância e o estado de fantasma
// do pato local.

const taskReach = 46.0 // px do pé do pato até o centro do ponto

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
