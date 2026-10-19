package over

import (
	"image/color"

	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// over_text.go — texto em tamanho pequeno (a fonte da tela é 18).

func (o *Over) face(size float64) *text.GoTextFace {
	if size == 18 {
		return o.font
	}
	return ui.Face("PressStart2P", size)
}

func (o *Over) textAt(img *ebiten.Image, s string, x, y float64, c color.Color, size float64) {
	f := o.face(size)
	if f == nil {
		return
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, f, op)
}
