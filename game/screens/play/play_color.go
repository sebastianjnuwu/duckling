package play

import (
	"image/color"

	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/match"
	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_color.go — escolha de cor no lobby: 8 quadradinhos embaixo.
// O host trava repetido; o convidado só pede e espera o snapshot.

const (
	colY = 640.0
	colS = 40.0
	colG = 10.0
)

// colX devolve o x do quadradinho i (linha centrada).
func colX(i int) float64 {
	w := float64(actor.ColorCount)*colS + float64(actor.ColorCount-1)*colG
	return (1280-w)/2 + float64(i)*(colS+colG)
}

// colorUpdate atende o clique num quadradinho livre.
func (p *Play) colorUpdate() {
	s := p.sess
	if s == nil || p.phase != match.Lobby || p.confirm {
		return
	}
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}
	x, y := ebiten.CursorPosition()
	for i := 0; i < int(actor.ColorCount); i++ {
		cx := colX(i)
		if float64(x) < cx || float64(x) > cx+colS || float64(y) < colY || float64(y) > colY+colS {
			continue
		}
		if s.Taken(i) && s.Color(s.ID()) != i {
			return
		}
		s.PickColor(i)
		return
	}
}

// colorDraw desenha a paleta; tomada fica escura, a minha com aro.
func (p *Play) colorDraw(img *ebiten.Image) {
	s := p.sess
	if s == nil || p.phase != match.Lobby || p.confirm {
		return
	}
	mine := s.Color(s.ID())
	for i := 0; i < int(actor.ColorCount); i++ {
		x := colX(i)
		if i == mine {
			ebitenutil.DrawRect(img, x-4, colY-4, colS+8, colS+8, ui.Gold)
		}
		c := actor.Color(i).RGB()
		if s.Taken(i) && i != mine {
			c = color.RGBA{R: c.R/3 + 8, G: c.G/3 + 8, B: c.B/3 + 8, A: 255}
		}
		ebitenutil.DrawRect(img, x, colY, colS, colS, c)
	}
	if p.font == nil {
		return
	}
	label := i18n.T("lobby.color") + " · " + i18n.T("color."+actor.Color(mine).Name())
	w, _ := text.Measure(label, p.font, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate((1280-w)/2, colY-32)
	op.ColorScale.ScaleWithColor(color.RGBA{235, 240, 235, 255})
	text.Draw(img, label, p.font, op)
}
