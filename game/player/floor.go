package actor

import (
	"image"

	"duckling/game/map"
)

// floor.go — colisão vertical e piso: superfícies elevadas (cercas, pedras).

func foot(x, y float64) image.Rectangle {
	return image.Rect(int(x-10), int(y-8), int(x+10), int(y+8))
}

// hit diz se colide com obstáculo mais alto que o pato.
func hit(x, y, z float64, solids []world.Solid) bool {
	r := foot(x, y)
	for _, s := range solids {
		if s.H > z+3 && r.Overlaps(s.Rect) {
			return true
		}
	}
	return false
}

// Floor acha o piso mais alto sob os pés (chão 0 ou topo de sólido).
func Floor(x, y, z float64, solids []world.Solid) float64 {
	r := foot(x, y)
	top := 0.0
	for _, s := range solids {
		if r.Overlaps(s.Rect) && z >= s.H {
			if s.H > top {
				top = s.H
			}
		}
	}
	return top
}

// walled diz se o pato tá colado em paredão mais alto que o voo.
func (d *Duck) walled(solids []world.Solid) bool {
	r := image.Rect(int(d.Px-11), int(d.Py-9), int(d.Px+11), int(d.Py+9))
	for _, s := range solids {
		if s.H > d.z+flyH && r.Overlaps(s.Rect) {
			return true
		}
	}
	return false
}
