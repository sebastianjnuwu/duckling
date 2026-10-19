package play

import (
	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/network"
	"duckling/game/session"
)

// play_meet_fx.go — aviso do resultado da votação.

// ejectFx avisa quem saiu (ou que ninguém saiu).
func (p *Play) ejectFx(e session.Eject) {
	switch {
	case e.Skip:
		p.toast.Show(i18n.T("meet.skipped"))
	case e.Imp:
		p.toast.Show(i18n.T("meet.eject.imp"))
	default:
		p.toast.Show(i18n.T("meet.eject.crew"))
	}
	_ = audio.Play("gamekill")
}

// canAct diz se o pato local ainda joga (fantasma não vota/fala).
func (p *Play) canAct() bool { return !p.ghost() }

// meetRoster limita a lista ao que cabe no painel (2 colunas de 8).
func (p *Play) meetRoster() []session.RosterEntry {
	r := p.sess.Roster()
	if len(r) > mtRows*2 {
		r = r[:mtRows*2]
	}
	return r
}

// meetName acha o nome no elenco (id<0 = ninguém).
func (p *Play) meetName(id int) string {
	if id < 0 {
		return "-"
	}
	for _, e := range p.sess.Roster() {
		if e.ID == id {
			return e.Name
		}
	}
	return "?"
}

// meetVote devolve em quem votei (NoVote = ainda não).
func (p *Play) meetVote() int { return p.sess.MyVote() }

// meetVotes conta os votos em id.
func meetVotes(vs []network.Vote, id int) int {
	n := 0
	for _, v := range vs {
		if v.ID == id {
			n++
		}
	}
	return n
}
