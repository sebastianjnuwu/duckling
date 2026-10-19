package play

import (
	"fmt"

	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_meet_draw.go — o painel da reunião: escurece o mundo, título,
// relógio e chama a lista/falas.

func (p *Play) meetDraw(img *ebiten.Image) {
	s := p.sess
	if s == nil || p.font == nil {
		return
	}
	m := s.Meet()
	if !m.Active {
		return
	}
	ui.Dim(img)
	ui.Panel(img, mtX, mtY, mtW, mtH)
	p.meetHead(img, m)
	p.meetList(img, m)
	p.meetChat(img, m)
}

func (p *Play) meetHead(img *ebiten.Image, m session.MeetInfo) {
	title := i18n.T("meet.body")
	if m.Reason == "bell" {
		title = i18n.T("meet.bell")
	}
	ui.DrawCenterColored(img, title, p.font, 84, ui.Gold)
	sub := fmt.Sprintf("%s: %s", i18n.T("meet.by"), p.meetName(m.By))
	if m.Reason == "body" {
		sub += "   " + fmt.Sprintf("%s: %s", i18n.T("meet.of"), p.meetName(m.Of))
	}
	ui.DrawCenterColored(img, sub, ui.Face("PressStart2P", 12), 116, mtCream)
	label := i18n.T("meet.discuss")
	if m.Open {
		label = i18n.T("meet.vote")
	}
	tm := fmt.Sprintf("%s %02d:%02d", label, int(m.Left)/60, int(m.Left)%60)
	p.rightText(img, tm, mtX+mtW-30, 84, ui.Gold, 18)
}
