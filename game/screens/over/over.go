package over

import (
	"duckling/engine/audio"
	"duckling/engine/rendering"
	"duckling/engine/save"
	"duckling/game/match"
	"duckling/game/session"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// over.go — tela de fim: quem venceu, por quê, os papéis e os
// números. O dono volta pro lobby (revanche) ou fecha a sala.

const winCoins = 10

type Over struct {
	next  func(string)
	font  *text.GoTextFace
	big   *text.GoTextFace
	back  ui.Btn
	menu  ui.Btn
	sess  *session.Session
	paid  bool
	bonus int
}

// NewOver monta a tela (fonte do loading, troca de tela do router).
func NewOver(font *text.GoTextFace, next func(string)) *Over {
	return &Over{
		next: next, font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 26},
		back: ui.Btn{X: 380, Y: 620, W: 240, H: 56, Icon: ui.IconPlay},
		menu: ui.Btn{X: 660, Y: 620, W: 240, H: 56, Icon: ui.IconDoor},
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
		o.paid, o.bonus = false, 0
		o.next("play")
		return nil
	}
	if !o.paid {
		o.pay(s)
	}
	if s.Fatal() != "" {
		session.Set(nil)
		o.next("menu")
		return nil
	}
	if !s.IsHost() {
		return nil
	}
	if o.back.Update() {
		_ = audio.Play("click")
		s.Reset()
	}
	if o.menu.Update() {
		_ = audio.Play("click")
		s.Reset()
		s.Close()
		session.Set(nil)
		o.next("menu")
	}
	return nil
}

// pay credita o bônus de vitória uma vez e toca o som do desfecho.
func (o *Over) pay(s *session.Session) {
	o.paid = true
	r := s.Result()
	role, ok := s.Role(s.ID())
	won := ok && ((r.Winner == match.WinCrew && role == match.Crew) ||
		(r.Winner == match.WinImp && role == match.Imp))
	if won {
		_ = audio.Play("gamewinning")
	} else {
		_ = audio.Play("gameover")
	}
	if !won {
		return
	}
	o.bonus = winCoins
	save.AddCoins(winCoins)
	save.Store()
}
