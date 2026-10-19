package help

import (
	"fmt"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/input"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// help_draw.go — fundo, painel e os blocos na ordem: objetivo,
// papéis, tarefas, teclas (remapeadas) e dicas.

func (h *Help) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, 120, 24, 1040, 672)
	ui.Title(img, ui.IconNote, i18n.T("help.title"), h.big, 44, ui.Gold)

	h.sec(img, "help.goal", 120)
	h.txt(img, i18n.T("help.goal1"), 148)
	h.txt(img, i18n.T("help.goal2"), 170)

	h.sec(img, "help.roles", 214)
	h.txt(img, i18n.T("help.roles1"), 242)
	h.txt(img, i18n.T("help.roles2"), 264)

	h.sec(img, "help.tasks", 308)
	h.txt(img, i18n.T("help.tasks1"), 336)
	h.txt(img, i18n.T("help.tasks2"), 358)

	h.sec(img, "help.keys", 402)
	h.txt(img, keyLine(0), 430)
	h.txt(img, keyLine(1), 452)

	h.sec(img, "help.tips", 496)
	h.txt(img, i18n.T("help.tips1"), 524)
	h.txt(img, i18n.T("help.tips2"), 546)

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
