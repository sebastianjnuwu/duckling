package shop

import (
	"duckling/src/audio"
	"duckling/src/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// shop.go — loja (placeholder): título, EM BREVE, volta.
// ESC ou botão volta pro menu.

type Shop struct {
	next func(string)
	font *text.GoTextFace
	big  *text.GoTextFace
	back ui.Btn
}

func NewShop(font *text.GoTextFace, next func(string)) *Shop {
	return &Shop{
		next: next,
		font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 64},
		back: ui.Btn{X: 440, Y: 550, W: 400, H: 72},
	}
}

func (s *Shop) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.next("menu")
		return nil
	}
	if s.back.Update() {
		_ = audio.Play("click")
		s.next("menu")
	}
	return nil
}
