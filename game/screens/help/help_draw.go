package help

import (
	"fmt"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/input"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// help_draw.go — fundo, painel principal e os 5 cartões internos:
// objetivo, papéis, tarefas, teclas (remapeadas) e dicas.

func (h *Help) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, 100, 24, 1080, 672)
	h.x.Draw(img, "X", h.font)
	ui.Title(img, ui.IconNote, i18n.T("help.title"), h.big, 48, ui.Gold)

	h.card(img, 108, ui.IconPlay, "help.goal",
		i18n.T("help.goal1"), i18n.T("help.goal2"), cream, cream)

	h.card(img, 206, ui.IconUser, "help.roles",
		i18n.T("help.roles1"), i18n.T("help.roles2"), crewCol, impCol)

	h.card(img, 304, ui.IconHand, "help.tasks",
		i18n.T("help.tasks1"), i18n.T("help.tasks2"), cream, cream)

	h.card(img, 402, ui.IconGear, "help.keys",
		keyLine(0), keyLine(1), ui.Gold, cream)

	h.card(img, 500, ui.IconSpk, "help.tips",
		i18n.T("help.tips1"), i18n.T("help.tips2"), cream, cream)

	h.back.Draw(img, i18n.T("help.back"), h.font)
}

// keyLine monta as duas linhas de teclas com o remapeamento atual.
func keyLine(n int) string {
	if n == 0 {
		return fmt.Sprintf(i18n.T("help.keyMove"),
			input.Name(input.Key(input.Up)), input.Name(input.Key(input.Down)),
			input.Name(input.Key(input.Left)), input.Name(input.Key(input.Right)))
	}
	return fmt.Sprintf(i18n.T("help.keyAct"),
		input.Name(input.Key(input.Jump)), input.Name(input.Key(input.Fly)),
		input.Name(input.Key(input.Interact)), input.Name(input.Key(input.Kill)))
}
