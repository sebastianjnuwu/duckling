package play

import (
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// play_over.go — fim de partida (mínimo): o dono volta pro lobby.

func (p *Play) overUpdate() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		p.confirm = true
	}
	if p.sess != nil && p.sess.IsHost() && p.backB.Update() {
		p.sess.Reset()
	}
	return nil
}

func (p *Play) overDraw(img *ebiten.Image) {
	ui.Dim(img)
	ui.Panel(img, 360, 190, 560, 320)
	ui.Title(img, ui.IconPlay, i18n.T("match.over"), p.font, 250, ui.Gold)
	if p.sess != nil && p.sess.IsHost() {
		p.backB.Draw(img, i18n.T("match.back"), p.font)
	} else {
		ui.DrawCenter(img, i18n.T("match.wait"), p.font, 452)
	}
}
