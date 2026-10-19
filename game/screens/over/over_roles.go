package over

import (
	"image/color"

	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/match"
	"duckling/game/player"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// over_roles.go — a lista de papéis (duas colunas) e os votos da
// última reunião, com nome de quem votou e em quem.

const (
	rowH    = 30.0
	colX    = 190.0
	colW    = 440.0
	colSt   = 470.0
	voteRow = 24.0
)

func (o *Over) roles(img *ebiten.Image, s *session.Session) {
	o.textAt(img, i18n.T("over.roles"), colX, 200, ui.Gold, 12)
	roster := s.Roster()
	rev := map[int]match.Role{}
	for _, r := range s.Reveal() {
		if r.Role == "imp" {
			rev[r.ID] = match.Imp
		} else {
			rev[r.ID] = match.Crew
		}
	}
	for i, e := range roster {
		if i >= 16 {
			break
		}
		col, row := i/8, i%8
		x := colX + float64(col)*colSt
		y := 224.0 + float64(row)*rowH
		ebitenutil.DrawRect(img, x, y+5, 18, 18, actor.Color(e.Color).RGB())
		name, c := e.Name, cream
		if e.Ghost {
			c = dimmer
		}
		o.textAt(img, name, x+28, y+3, c, 18)
		if e.Ghost {
			n, _ := text.Measure(name, o.font, 0)
			ebitenutil.DrawRect(img, x+28, y+15, n, 2, c)
		}
		o.textAt(img, roleName(rev[e.ID]), x+colW-150, y+7, roleColor(rev[e.ID]), 12)
	}
}

// roleName/roleColor traduzem o papel do jogador.
func roleName(r match.Role) string {
	if r == match.Imp {
		return i18n.T("over.role.imp")
	}
	return i18n.T("over.role.crew")
}

func roleColor(r match.Role) color.Color {
	if r == match.Imp {
		return winImp
	}
	return cream
}
