package play

import (
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

func (p *Play) netUpdate(x0, y0, x1, y1 int) {
	s := p.sess
	if s == nil {
		return
	}
	s.Tick()
	p.d.SetColor(actor.Color(s.Color(s.ID())))
	if s.IsHost() {
		p.hostKills()
		for id, in := range s.HostInputs() {
			if p.lock[id] > 0 {
				in = actor.In{}
			}
			d := p.remote(id)
			d.SetColor(actor.Color(s.Color(id)))
			d.Update(x0, y0, x1, y1, p.solids(), in)
		}
		for id := range p.rem {
			p.rghost[id] = s.Ghost(id)
		}
		s.Publish(p.snapshot())
		return
	}
	s.SendInput(actor.LocalIn())
	for _, pl := range s.States() {
		p.rghost[pl.ID] = pl.Ghost
		if pl.ID == s.ID() {
			p.reconcile(pl)
			continue
		}
		d := p.remote(pl.ID)
		d.Interp(pl.X, pl.Y, pl.Z, pl.Dir, pl.Fly)
		d.SetColor(actor.Color(pl.Color))
		p.rname[pl.ID] = pl.Name
	}
}

// reconcile corrige o pato local só quando a divergência é grande.
func (p *Play) reconcile(pl network.Player) {
	dx, dy := p.d.Px-pl.X, p.d.Py-pl.Y
	if dx*dx+dy*dy > 48*48 {
		p.d.Snap(pl.X, pl.Y, pl.Z, pl.Dir, pl.Fly)
	}
}
