package over

import (
	"duckling/engine/audio"
	"duckling/engine/rendering"
	"duckling/game/match"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// over.go — tela de fim: quem venceu, por quê, os papéis e os
// números. O dono volta pro lobby e os convidados acompanham.

type Over struct {
	next func(string)
	font *text.GoTextFace
	big  *text.GoTextFace
	back ui.Btn
	sess *session.Session
}

// NewOver monta a tela (fonte do loading, troca de tela do router).
func NewOver(font *text.GoTextFace, next func(string)) *Over {
	return &Over{
		next: next, font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 26},
		back: ui.Btn{X: 500, Y: 620, W: 280, H: 56, Icon: ui.IconPlay},
	}
}

func (o *Over) Update() error {
	s := session.Get()
	o.sess = s
	if s == nil {
		o.next("menu")
		return nil
	}
	// A partida voltou pro lobby (o dono mandou): a tela sai.
	if s.Phase() != match.End {
		o.next("play")
		return nil
	}
	if s.IsHost() && o.back.Update() {
		_ = audio.Play("click")
		s.Reset()
	}
	return nil
}
