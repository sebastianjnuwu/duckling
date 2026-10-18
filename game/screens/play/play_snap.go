package play

import (
	"duckling/engine/i18n"
	"duckling/engine/network"
	"duckling/game/session"
)

// play_snap.go — monta o snapshot do host e trata a saída da
// partida (eventos viram toast; erro fatal volta pro menu).

func (p *Play) snapshot() []network.Player {
	out := make([]network.Player, 0, len(p.rem)+1)
	out = append(out, network.Player{
		ID: 0, Name: p.sess.Name(), X: p.d.Px, Y: p.d.Py,
		Z: p.d.Alt(), Dir: p.d.Dir(), Fly: p.d.Flying(),
	})
	for id, d := range p.rem {
		out = append(out, network.Player{
			ID: id, Name: p.rname[id], X: d.Px, Y: d.Py,
			Z: d.Alt(), Dir: d.Dir(), Fly: d.Flying(),
		})
	}
	return out
}

// leaveIfBroken mostra os avisos e, se a sessão morreu, volta
// pro menu. Devolve true quando já trocou de tela.
func (p *Play) leaveIfBroken() bool {
	if p.sess == nil {
		return false
	}
	for _, k := range p.sess.TakeEvents() {
		p.toast.Show(i18n.T(k))
	}
	if f := p.sess.Fatal(); f != "" {
		p.toast.Show(i18n.T(f))
		session.Set(nil)
		p.next("menu")
		return true
	}
	return false
}
