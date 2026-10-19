package play

import (
	"fmt"
	"image/color"
	"math"

	"duckling/engine/i18n"
	"duckling/engine/input"
	"duckling/engine/rendering"
	"duckling/game/match"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// play_kill_hud.go — o botão do MATAR: o anel enche conforme o
// cooldown acaba; a tecla fica no meio (ou os segundos restantes).

const (
	killCX, killCY = 104.0, 606.0
	killR          = 30.0
)

func (p *Play) killDraw(img *ebiten.Image) {
	if p.sess == nil || p.font == nil || p.phase != match.Playing || !p.imp() || p.ghost() {
		return
	}
	cd := p.sess.Cooldown()
	frac := 1 - cd/session.KillCooldown.Seconds()
	if frac < 0 {
		frac = 0
	}
	if frac > 1 {
		frac = 1
	}
	ring := ui.Gold
	if p.lock[p.localID()] > 0 {
		ring = ui.Brass
	}
	vector.DrawFilledCircle(img, killCX, killCY, float32(killR), color.RGBA{18, 20, 18, 200}, true)
	p.killArc(img, frac, ring)
	label := input.Name(input.Key(input.Kill))
	if cd > 0 {
		label = fmt.Sprintf("%.0f", math.Ceil(cd))
	}
	p.killText(img, label, killCX, killCY-11, ui.Cream)
	p.killText(img, i18n.T("settings.kill"), killCX, killCY+killR+10, ui.Brass)
}

// killArc desenha o anel dourado na fração pronta (0..1).
func (p *Play) killArc(img *ebiten.Image, frac float64, c color.Color) {
	const steps = 40
	n := int(frac * steps)
	for i := 0; i < n; i++ {
		a0 := -math.Pi/2 + float64(i)*2*math.Pi/steps
		a1 := a0 + 2*math.Pi/steps
		x0, y0 := killCX+math.Cos(a0)*(killR-5), killCY+math.Sin(a0)*(killR-5)
		x1, y1 := killCX+math.Cos(a1)*(killR-5), killCY+math.Sin(a1)*(killR-5)
		vector.StrokeLine(img, float32(x0), float32(y0), float32(x1), float32(y1), 4, c, true)
	}
}

// killText centraliza um texto compacto no botão.
func (p *Play) killText(img *ebiten.Image, s string, cx, y float64, c color.Color) {
	w, _ := text.Measure(s, p.font, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate(cx-w/2, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, p.font, op)
}
