package play

import (
	"fmt"
	"image/color"
	"strings"

	"duckling/engine/i18n"
	"duckling/game/match"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_hud.go — faixa do topo: fase, vivos e, na partida, o meu
// papel (com os comparsas, se impostor). Texto sempre por i18n.

func (p *Play) hudDraw(img *ebiten.Image) {
	if p.sess == nil || p.font == nil {
		return
	}
	p.hudLeft(img, 20)
}

func (p *Play) hudLeft(img *ebiten.Image, x float64) {
	y := 16.0
	p.hudText(img, i18n.T(phaseKey(p.phase)), x, y)
	if p.phase == match.Lobby {
		return
	}
	p.hudText(img, fmt.Sprintf("%s %d", i18n.T("match.alive"), p.alive()), x, y+30)
	if r, ok := p.sess.Role(p.sess.ID()); ok {
		key := "match.crew"
		if r == match.Imp {
			key = "match.imp"
		}
		p.hudText(img, i18n.T(key)+mates(p.sess.Mates(p.sess.ID())), x, y+60)
	}
}

func (p *Play) hudText(img *ebiten.Image, s string, x, y float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x+2, y+2)
	op.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 180})
	text.Draw(img, s, p.font, op)
	op = &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.RGBA{235, 240, 235, 235})
	text.Draw(img, s, p.font, op)
}

// alive conta quem não é fantasma (host) ou quem veio no snapshot.
func (p *Play) alive() int {
	if p.sess == nil {
		return 1
	}
	if p.sess.IsHost() {
		return p.sess.AliveCount()
	}
	n := 0
	for _, pl := range p.sess.States() {
		if !pl.Ghost {
			n++
		}
	}
	if n > 0 {
		return n
	}
	return p.sess.Count()
}

// mates monta " (COMPARSA 2, 3)" pro impostor.
func mates(ids []int) string {
	if len(ids) == 0 {
		return ""
	}
	parts := make([]string, len(ids))
	for i, id := range ids {
		parts[i] = fmt.Sprintf("%d", id)
	}
	return "  (" + i18n.T("match.mates") + " " + strings.Join(parts, ", ") + ")"
}

func phaseKey(ph match.Phase) string {
	switch ph {
	case match.Playing:
		return "match.playing"
	case match.Meeting:
		return "match.meeting"
	case match.End:
		return "match.over"
	}
	return "match.lobby"
}
