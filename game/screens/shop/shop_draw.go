package shop

import (
	"fmt"
	"image/color"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/engine/save"
	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
)

// shop_draw.go — fundo, título, saldo, os 4 cartões e o voltar.

var (
	cream = color.RGBA{235, 240, 235, 255}
	okCol = color.RGBA{120, 210, 150, 255}
)

func (s *Shop) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Title(img, ui.IconShop, i18n.T("shop.title"), s.big, 70, ui.Gold)
	s.textMid(img, fmt.Sprintf("%s %d", i18n.T("shop.coins"), save.Coins()), 640, 148, ui.Gold, 14)
	for i := 0; i < int(actor.HatCount); i++ {
		s.card(img, actor.Hat(i))
	}
	s.back.Draw(img, i18n.T("shop.back"), s.font)
	s.toast.Draw(img, s.font)
}

// card desenha um chapéu: nome, pato vestindo e preço/estado.
func (s *Shop) card(img *ebiten.Image, h actor.Hat) {
	i := int(h)
	x := cardX(i)
	ui.Panel(img, x, cardY, cardW, cardH)
	switch {
	case save.Hat() == i:
		frame(img, x-2, cardY-2, cardW+4, cardH+4, 3, ui.Gold)
	case cardHit(i):
		frame(img, x-2, cardY-2, cardW+4, cardH+4, 3, ui.Cream)
	}
	s.textMid(img, i18n.T("hat."+h.Name()), x+cardW/2, cardY+18, cream, 12)
	s.preview(img, h, x+cardW/2, 360)
	label, col := status(h)
	s.textMid(img, label, x+cardW/2, cardY+cardH-52, col, 14)
}

// preview desenha o pato da loja vestindo h (arte carrega lazy).
func (s *Shop) preview(img *ebiten.Image, h actor.Hat, cx, cy float64) {
	if s.duck == nil {
		s.duck = actor.New(assets.Images, 0, 0, actor.Cast["duck"])
	}
	if !s.duck.Ok() {
		s.duck.Load(assets.Images)
	}
	s.duck.SetHat(h)
	s.duck.Draw(img, cx, cy, 3)
}

// status devolve o rodapé do cartão (texto e cor).
func status(h actor.Hat) (string, color.Color) {
	id := int(h)
	switch {
	case save.Hat() == id:
		return i18n.T("shop.equipped"), okCol
	case save.HatOwned(id):
		return i18n.T("shop.equip"), cream
	}
	return fmt.Sprintf("%d %s", h.Price(), i18n.T("shop.coins")), ui.Gold
}
