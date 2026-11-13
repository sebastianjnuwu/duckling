package help

import (
	"duckling/engine/audio"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// help.go — "como jogar": objetivo, papéis, tarefas, teclas
// (lidas do remapeamento) e dicas. ESC, X ou VOLTAR saem.

type Help struct {
	next func(string)
	font *text.GoTextFace
	big  *text.GoTextFace
	secF *text.GoTextFace
	subF *text.GoTextFace
	back ui.Btn
	x    ui.Btn
}

func NewHelp(font *text.GoTextFace, next func(string)) *Help {
	return &Help{
		next: next, font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 34},
		secF: &text.GoTextFace{Source: font.Source, Size: 15},
		subF: &text.GoTextFace{Source: font.Source, Size: 12},
		back: ui.Btn{X: 460, Y: 616, W: 360, H: 58, Icon: ui.IconDoor},
		x:    ui.Btn{X: 1110, Y: 44, W: 44, H: 44},
	}
}

func (h *Help) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		h.next("menu")
		return nil
	}
	if h.x.Push() {
		_ = audio.Play("click")
		h.next("menu")
		return nil
	}
	if h.back.Update() {
		_ = audio.Play("click")
		h.next("menu")
	}
	return nil
}
