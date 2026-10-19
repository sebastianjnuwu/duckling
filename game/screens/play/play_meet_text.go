package play

import (
	"image/color"

	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_meet_text.go — texto da reunião: tamanho 18 (fonte da tela)
// ou menor, alinhado à esquerda, centro e direita.

func (p *Play) face(size float64) *text.GoTextFace {
	if size == 18 {
		return p.font
	}
	return ui.Face("PressStart2P", size)
}

func (p *Play) textAt(img *ebiten.Image, s string, x, y float64, c color.Color, size float64) {
	f := p.face(size)
	if f == nil {
		return
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, f, op)
}

func (p *Play) centerText(img *ebiten.Image, s string, cx, y float64, c color.Color, size float64) {
	f := p.face(size)
	if f == nil {
		return
	}
	w, _ := text.Measure(s, f, 0)
	p.textAt(img, s, cx-w/2, y, c, size)
}

func (p *Play) rightText(img *ebiten.Image, s string, right, y float64, c color.Color, size float64) {
	f := p.face(size)
	if f == nil {
		return
	}
	w, _ := text.Measure(s, f, 0)
	p.textAt(img, s, right-w, y, c, size)
}
