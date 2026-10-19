package help

import (
	"duckling/engine/audio"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// help.go — "como jogar": objetivo, papéis, tarefas, teclas
// (lidas do remapeamento) e dicas. ESC ou o botão volta ao menu.

type Help struct {
	next func(string)
	font *text.GoTextFace
	big  *text.GoTextFace
	back ui.Btn
}

func NewHelp(font *text.GoTextFace, next func(string)) *Help {
	return &Help{
		next: next, font: font,
		big:  &text.GoTextFace{Source: font.Source, Size: 48},
		back: ui.Btn{X: 440, Y: 616, W: 400, H: 64, Icon: ui.IconDoor},
	}
}

func (h *Help) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		h.next("menu")
		return nil
	}
	if h.back.Update() {
		_ = audio.Play("click")
		h.next("menu")
	}
	return nil
}
