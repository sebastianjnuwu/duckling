package help

import (
	"image/color"

	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// help_text.go — cartões escuros com borda dourada, ícone e texto.

var (
	cream   = color.RGBA{235, 240, 235, 255}
	crewCol = color.RGBA{130, 225, 165, 255}
	impCol  = color.RGBA{240, 130, 130, 255}
	cardBg  = color.RGBA{24, 32, 52, 255}
	cardRim = color.RGBA{56, 74, 112, 255}
)

// card desenha um bloco temático (992x86) com ícone, título e 2 linhas.
func (h *Help) card(
	img *ebiten.Image,
	y float64,
	ic ui.Icon,
	titleKey, l1, l2 string,
	c1, c2 color.Color,
) {
	const x, w, ch = 144.0, 992.0, 86.0
	ebitenutil.DrawRect(img, x, y, w, ch, cardRim)
	ebitenutil.DrawRect(img, x+2, y+2, w-4, ch-4, cardBg)
	ebitenutil.DrawRect(img, x, y, 6, ch, ui.Gold)

	ui.DrawIcon(img, ic, x+20, y+10, ui.Gold)
	h.left(img, i18n.T(titleKey), h.secF, x+52, y+14, ui.Gold)
	h.left(img, l1, h.subF, x+52, y+40, c1)
	h.left(img, l2, h.subF, x+52, y+62, c2)
}

func (h *Help) left(
	img *ebiten.Image,
	s string,
	f *text.GoTextFace,
	x, y float64,
	c color.Color,
) {
	if f == nil {
		return
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(img, s, f, op)
}
