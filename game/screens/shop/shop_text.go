package shop

import (
	"image/color"

	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// shop_text.go — texto em tamanhos pequenos (a fonte da tela é 18).

func (s *Shop) face(size float64) *text.GoTextFace {
	if size == 18 {
		return s.font
	}
	return ui.Face("PressStart2P", size)
}

// textMid centraliza o texto em cx.
func (s *Shop) textMid(img *ebiten.Image, str string, cx, y float64, c color.Color, size float64) {
	f := s.face(size)
	if f == nil {
		return
	}
	w, _ := text.Measure(str, f, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate(cx-w/2, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, str, f, op)
}
