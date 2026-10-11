package shop

import (
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
	ui.Title(img, ui.IconShop, i18n.T("shop.title"), s.big, 120, ui.Gold)
	ui.DrawCenter(img, i18n.T("play.soon"), s.font, 330)
	s.back.Draw(img, i18n.T("shop.back"), s.font)
}
