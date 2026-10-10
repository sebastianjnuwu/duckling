package shop

import (
	"image/color"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// shop_draw.go — fundo, título ouro, EM BREVE, botão voltar.

func (s *Shop) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.DrawCenterColored(img, i18n.T("shop.title"), s.big, 120, color.RGBA{255, 215, 0, 255})
	ui.DrawCenter(img, i18n.T("play.soon"), s.font, 330)
	s.back.Draw(img, i18n.T("shop.back"), s.font)
}
