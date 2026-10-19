package shop

import (
	"duckling/engine/audio"
	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/engine/save"
	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// shop.go — loja: cartões de chapéu (comprar/equipar), saldo,
// avisos e volta. O desenho fica em shop_draw.go.

type Shop struct {
	next  func(string)
	font  *text.GoTextFace
	big   *text.GoTextFace
	back  ui.Btn
	duck  *actor.Duck
	toast ui.Toast
}

func NewShop(font *text.GoTextFace, next func(string)) *Shop {
	return &Shop{
		next: next, font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 64},
		back: ui.Btn{X: 440, Y: 592, W: 400, H: 64, Icon: ui.IconDoor},
	}
}

func (s *Shop) Update() error {
	s.toast.Update()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.next("menu")
		return nil
	}
	if s.back.Update() {
		_ = audio.Play("click")
		s.next("menu")
		return nil
	}
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	for i := 0; i < int(actor.HatCount); i++ {
		if cardHit(i) {
			s.pick(actor.Hat(i))
			return nil
		}
	}
	return nil
}

// pick compra o que falta e equipa (sem moeda, só avisa).
func (s *Shop) pick(h actor.Hat) {
	id := int(h)
	bought := false
	if !save.HatOwned(id) {
		if !save.BuyHat(id, h.Price()) {
			s.toast.Show(i18n.T("shop.nomoney"))
			return
		}
		bought = true
		s.toast.Show(i18n.T("shop.bought"))
	}
	if save.Hat() != id {
		save.EquipHat(id)
		bought = true
	}
	if bought {
		_ = audio.Play("bumped")
	}
	save.Store()
}
