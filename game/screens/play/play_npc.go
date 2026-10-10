package play

import (
	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_npc.go — vermelhinhos passeando perto do spawn.

func (p *Play) spawnNPCs(sx, sy float64, red *ebiten.Image) {
	for _, off := range [][2]float64{{60, 0}, {-50, 30}, {10, -60}} {
		p.npcs = append(p.npcs, actor.NewNPC(red, sx+off[0], sy+off[1], actor.Cast["red"]))
	}
	p.npcs[2].Still = true
}
