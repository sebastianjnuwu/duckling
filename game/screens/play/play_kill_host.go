package play

import (
	"duckling/game/player"
	"duckling/game/session"
)

// play_kill_host.go — o lado do host do MATAR: consome os pedidos
// dos convidados, confere o alcance e trava quem matou.

// hostKills consome os pedidos dos convidados e valida cada um.
func (p *Play) hostKills() {
	for _, r := range p.sess.TakeKills() {
		p.tryKill(r.By, r.Target)
	}
}

// tryKill valida o alcance e chama a sessão (host).
func (p *Play) tryKill(by, target int) bool {
	if p.sess == nil || !p.sess.CanKill(by, target) {
		return false
	}
	a, v := p.duckOf(by), p.duckOf(target)
	if a == nil || v == nil {
		return false
	}
	if dx, dy := a.Px-v.Px, a.Py-v.Py; dx*dx+dy*dy > session.KillReach*session.KillReach {
		return false
	}
	if !p.sess.Kill(by, target, v.Px, v.Py) {
		return false
	}
	p.lock[by] = killLock
	return true
}

// duckOf acha o pato local ou remoto pelo id (host).
func (p *Play) duckOf(id int) *actor.Duck {
	if p.sess == nil {
		return nil
	}
	if id == p.sess.ID() {
		return p.d
	}
	return p.rem[id]
}

// lockUpdate desconta o tempo travado depois de matar.
func (p *Play) lockUpdate(dt float64) {
	for id, t := range p.lock {
		if t -= dt; t > 0 {
			p.lock[id] = t
		} else {
			delete(p.lock, id)
		}
	}
}
