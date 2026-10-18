package play

import (
	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// play_quit.go — ESC no jogo: confirma antes de voltar ao menu.

func (p *Play) confirmUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		p.confirm = false
		return
	}
	if p.no.Update() {
		p.confirm = false
	}
	if p.qx.Push() {
		p.confirm = false
	}
	if p.yes.Update() {
		p.confirm = false
		session.Set(nil)
		p.next("menu")
	}
}

func (p *Play) confirmDraw(img *ebiten.Image) {
	ui.Dim(img)
	ui.Panel(img, 360, 230, 560, 260)
	p.qx.Draw(img, "X", p.font)
	ui.Title(img, ui.IconDoor, i18n.T("play.quit"), p.font, 330, ui.Cream)
	p.yes.Draw(img, i18n.T("menu.yes"), p.font)
	p.no.Draw(img, i18n.T("menu.no"), p.font)
}
