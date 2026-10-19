package play

import (
	"image/color"

	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_dead.go — mortes: corpos no chão, puff/som quando o corpo
// aparece e o aviso de quem virou fantasma.

// deadDraw desenha os corpos caídos no chão.
func (p *Play) deadDraw(img *ebiten.Image) {
	if p.sess == nil {
		return
	}
	for _, b := range p.sess.Bodies() {
		x, y := (b.X-p.camX)*p.zoom, (b.Y-p.camY)*p.zoom
		p.d.DrawDead(img, x, y, p.zoom)
	}
}

// deadUpdate avisa uma vez quando o pato local vira fantasma e
// dispara o efeito dos corpos novos.
func (p *Play) deadUpdate() {
	if p.sess == nil {
		return
	}
	g := p.ghost()
	if g && !p.wasGhost {
		p.toast.Show(i18n.T("match.dead"))
	}
	p.wasGhost = g
	p.deadFx()
}

// deadFx faz som e penas na primeira vez que cada corpo aparece.
func (p *Play) deadFx() {
	for _, b := range p.sess.Bodies() {
		if p.bodySeen[b.ID] {
			continue
		}
		p.bodySeen[b.ID] = true
		_ = audio.Play("gamekill")
		x, y := (b.X-p.camX)*p.zoom, (b.Y-p.camY)*p.zoom
		p.puff.Burst(x, y-12, 16, ui.Gold, ui.Cream, color.RGBA{200, 60, 60, 255})
	}
}
