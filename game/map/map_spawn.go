package world

import (
	"image"

	"github.com/lafriks/go-tiled"
)

// map_spawn.go — contrato de nascimento: objetos "spawn". Vários
// = um ponto por pato (o host distribui). Sem nenhum, o centro
// visível do mapa.

// Spawns devolve os pontos de nascimento (nunca vazio).
func (m *Map) Spawns() []image.Point {
	if len(m.spawns) == 0 {
		return []image.Point{{X: int(m.sx), Y: int(m.sy)}}
	}
	return m.spawns
}

func (m *Map) takeSpawn(o *tiled.Object) {
	p := image.Pt(int(o.X), int(o.Y))
	m.spawns = append(m.spawns, p)
	if len(m.spawns) == 1 {
		m.sx, m.sy = float64(p.X), float64(p.Y)
	}
}
