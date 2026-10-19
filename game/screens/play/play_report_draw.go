package play

import (
	"duckling/engine/i18n"
	"duckling/engine/input"
	"duckling/engine/rendering"
	"duckling/game/match"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_report_draw.go — o sino no mundo e a dica de ação perto
// do pato (reportar corpo / chamar reunião).

// bellDraw marca o sino no mundo (o mapa não tem arte pra ele).
func (p *Play) bellDraw(img *ebiten.Image) {
	for _, b := range p.mp.Bells {
		x := (float64(b.X)-p.camX)*p.zoom - 9*p.zoom
		y := (float64(b.Y)-p.camY)*p.zoom - 9*p.zoom
		ui.DrawIcon(img, ui.IconNote, x, y, ui.Gold)
	}
}

// reportDraw mostra a dica "tecla: reportar/chamar" perto do pato.
func (p *Play) reportDraw(img *ebiten.Image) {
	if p.sess == nil || p.phase != match.Playing || !p.canAct() {
		return
	}
	label := ""
	if _, ok := p.nearBody(); ok {
		label = i18n.T("match.report")
	} else if !p.bellUsed && p.nearBell() {
		label = i18n.T("match.bell")
	}
	if label == "" {
		return
	}
	x := (p.d.Px - p.camX) * p.zoom
	y := (p.d.Py-p.camY)*p.zoom - 64*p.zoom
	p.centerText(img, input.Name(input.Key(input.Interact))+": "+label, x, y, ui.Gold, 10)
}
