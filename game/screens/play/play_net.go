package play

import (
	"math"

	"duckling/content"
	"duckling/engine/network"
	"duckling/game/player"
	"duckling/game/session"
)

// play_net.go — costura da partida na tela: o host simula os
// remotos com o input que chega; o cliente interpola os estados.

// syncSession percebe troca de sessão e reinicia o elenco.
func (p *Play) syncSession() {
	s := session.Get()
	if s == p.sess {
		return
	}
	p.sess = s
	p.rem = map[int]*actor.Duck{}
	p.rname = map[int]string{}
	p.rghost = map[int]bool{}
	id := 0
	if s != nil {
		id = s.ID()
	}
	p.d.Px, p.d.Py = p.spawnAt(id)
}

// spawnAt espalha os patos num anel em volta do spawn padrão.
func (p *Play) spawnAt(id int) (float64, float64) {
	sx, sy := p.mp.Spawn()
	if id <= 0 {
		return sx, sy
	}
	a := float64(id) * (math.Pi * 2 / 8)
	return sx + math.Cos(a)*48, sy + math.Sin(a)*48
}

func (p *Play) netUpdate(x0, y0, x1, y1 int) {
	s := p.sess
	if s == nil {
		return
	}
	s.Tick()
	if s.IsHost() {
		p.hostKills()
		for id, in := range s.HostInputs() {
			if p.lock[id] > 0 {
				in = actor.In{}
			}
			p.remote(id).Update(x0, y0, x1, y1, p.solids(), in)
		}
		for id := range p.rem {
			p.rghost[id] = s.Ghost(id)
		}
		s.Publish(p.snapshot())
		return
	}
	s.SendInput(actor.LocalIn().Encode())
	for _, pl := range s.States() {
		p.rghost[pl.ID] = pl.Ghost
		if pl.ID == s.ID() {
			p.reconcile(pl)
			continue
		}
		p.remote(pl.ID).Interp(pl.X, pl.Y, pl.Z, pl.Dir, pl.Fly)
		p.rname[pl.ID] = pl.Name
	}
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

// reconcile corrige o pato local só quando a divergência é grande.
func (p *Play) reconcile(pl network.Player) {
	dx, dy := p.d.Px-pl.X, p.d.Py-pl.Y
	if dx*dx+dy*dy > 48*48 {
		p.d.Snap(pl.X, pl.Y, pl.Z, pl.Dir, pl.Fly)
	}
}
