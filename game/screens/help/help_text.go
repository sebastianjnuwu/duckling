package help

import (
	"image/color"

	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// help_text.go — linhas curtas em tamanho pequeno, centradas.

var cream = color.RGBA{235, 240, 235, 255}

func (h *Help) face(size float64) *text.GoTextFace {
	if size == 18 {
		return h.font
	}
	return ui.Face("PressStart2P", size)
}

// line desenha o texto centrado em y.
func (h *Help) line(img *ebiten.Image, s string, y float64, c color.Color, size float64) {
	f := h.face(size)
	if f == nil {
		return
	}
	w, _ := text.Measure(s, f, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(img.Bounds().Dx())/2-w/2, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, f, op)
}

// sec desenha um título de seção dourado.
func (h *Help) sec(img *ebiten.Image, key string, y float64) {
	h.line(img, i18n.T(key), y, ui.Gold, 16)
}

// txt desenha uma linha de corpo.
func (h *Help) txt(img *ebiten.Image, s string, y float64) {
	h.line(img, s, y, cream, 12)
}
