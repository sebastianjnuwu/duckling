package play

import (
	"fmt"
	"image/color"

	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/player"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_meet_list.go — a lista de patos com cor, morto riscado,
// votos ao vivo e o botão de pular.

var mtCream = color.RGBA{235, 240, 235, 255}
var mtDim = color.RGBA{132, 142, 158, 210}

func (p *Play) meetList(img *ebiten.Image, m session.MeetInfo) {
	roster := p.meetRoster()
	cnt := len(roster)
	size := 18.0
	if cnt > mtRows {
		size = 12
	}
	my := p.sess.ID()
	mine := p.meetVote()
	mx, myy := ebiten.CursorPosition()
	p.textAt(img, i18n.T("meet.alive"), mtRowX, mtRowY-24, mtCream, 10)
	for i, e := range roster {
		x, y, w, h := meetRowRect(i, cnt)
		bg := color.RGBA{26, 30, 44, 230}
		if m.Open && p.canAct() && !e.Ghost && e.ID != my &&
			inRect(float64(mx), float64(myy), x, y, w, h) {
			bg = color.RGBA{40, 48, 72, 235}
		}
		ebitenutil.DrawRect(img, x, y, w, h, bg)
		if e.ID == mine {
			ebitenutil.DrawRect(img, x, y, 4, h, ui.Gold)
		}
		ebitenutil.DrawRect(img, x+10, y+10, 22, 22, actor.Color(e.Color).RGB())
		col := mtCream
		if e.Ghost {
			col = mtDim
		}
		p.textAt(img, e.Name, x+44, y+11, col, size)
		if e.Ghost {
			n, _ := text.Measure(e.Name, p.face(size), 0)
			ebitenutil.DrawRect(img, x+44, y+22, n, 2, col)
		}
		if n := meetVotes(m.Votes, e.ID); n > 0 {
			p.textAt(img, fmt.Sprintf("x%d", n), x+w-30, y+13, ui.Gold, 12)
		}
	}
	p.meetSkip(img, m, mine)
}

func (p *Play) meetSkip(img *ebiten.Image, m session.MeetInfo, mine int) {
	face := color.RGBA{40, 44, 60, 220}
	if m.Open && p.canAct() && mine == session.SkipVote {
		face = ui.Gold
	}
	ebitenutil.DrawRect(img, mtSkipX, mtSkipY, mtSkipW, mtSkipH, face)
	col := mtCream
	if face == ui.Gold {
		col = color.RGBA{16, 16, 30, 255}
	}
	p.centerText(img, i18n.T("meet.skip"), mtSkipX+mtSkipW/2, mtSkipY+11, col, 14)
}
