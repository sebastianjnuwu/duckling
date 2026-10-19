package over

import (
	"duckling/engine/i18n"
	"duckling/engine/network"
	"duckling/engine/rendering"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
)

// over_votes.go — os votos da última reunião, "quem > alvo"
// (quem pulou aparece como PULAR), em duas colunas.

func (o *Over) votes(img *ebiten.Image, s *session.Session) {
	o.textAt(img, i18n.T("over.votes"), colX, 486, ui.Gold, 12)
	vs := s.LastVotes()
	if len(vs) == 0 {
		o.textAt(img, i18n.T("over.novote"), colX, 512, dimmer, 10)
		return
	}
	for i, v := range vs {
		if i >= 8 {
			break
		}
		col, row := i/4, i%4
		x := colX + float64(col)*colSt
		y := 512.0 + float64(row)*voteRow
		o.textAt(img, voteLine(s, v), x, y, cream, 10)
	}
}

// voteLine monta "quem > alvo" (ou PULAR).
func voteLine(s *session.Session, v network.Vote) string {
	target := i18n.T("meet.skip")
	if v.ID != -1 {
		target = oName(s, v.ID)
	}
	return oName(s, v.By) + " > " + target
}

// oName acha o nome no elenco.
func oName(s *session.Session, id int) string {
	for _, e := range s.Roster() {
		if e.ID == id {
			return e.Name
		}
	}
	return "?"
}

// footer: o dono volta pro lobby (revanche) ou fecha a sala.
func (o *Over) footer(img *ebiten.Image, s *session.Session) {
	if !s.IsHost() {
		ui.DrawCenter(img, i18n.T("match.wait"), o.font, 640)
		return
	}
	o.back.Draw(img, i18n.T("over.again"), o.font)
	o.menu.Draw(img, i18n.T("over.menu"), o.font)
}
