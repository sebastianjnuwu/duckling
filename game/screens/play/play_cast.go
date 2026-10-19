package play

import (
	"math"

	"duckling/content"
	"duckling/game/player"
)

// play_cast.go — o elenco na tela: posição de nascimento em anel e
// o pato de outro jogador criado sob demanda.

// spawnAt espalha os patos num anel em volta do spawn padrão.
func (p *Play) spawnAt(id int) (float64, float64) {
	sx, sy := p.mp.Spawn()
	if id <= 0 {
		return sx, sy
	}
	a := float64(id) * (math.Pi * 2 / 8)
	return sx + math.Cos(a)*48, sy + math.Sin(a)*48
}

// remote devolve (criando na hora) o pato de um jogador.
func (p *Play) remote(id int) *actor.Duck {
	d := p.rem[id]
	if d == nil {
		sx, sy := p.spawnAt(id)
		d = actor.New(assets.Images, sx, sy, actor.Cast["duck"])
		p.rem[id] = d
		if p.sess != nil && p.sess.IsHost() {
			p.rname[id] = p.sess.PeerName(id)
		}
	}
	if !d.Ok() {
		d.Load(assets.Images)
	}
	return d
}
