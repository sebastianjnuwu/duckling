package play

import (
	"fmt"
	"image/color"
	"sort"
	"strings"

	"duckling/engine/i18n"
	"duckling/engine/rendering"
	"duckling/game/match"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_task_hud.go — barra do time (topo) e a minha lista (direita).

func (p *Play) taskDraw(img *ebiten.Image) {
	if p.sess == nil || p.font == nil || p.phase != match.Playing {
		return
	}
	p.barDraw(img)
	if !p.ghost() {
		p.listDraw(img)
	}
	p.arrowDraw(img)
}

// barDraw é a barra do time: preenchida = tarefas concluídas.
func (p *Play) barDraw(img *ebiten.Image) {
	done, total := p.sess.GlobalTasks()
	if total <= 0 {
		return
	}
	label := fmt.Sprintf("%s %d/%d", i18n.T("task.title"), done, total)
	ui.DrawCenterColored(img, label, p.font, 18, ui.Cream)
	x, y, w, h := 430.0, 48.0, 420.0, 16.0
	ebitenutil.DrawRect(img, x-2, y-2, w+4, h+4, color.RGBA{0, 0, 0, 170})
	ebitenutil.DrawRect(img, x, y, w, h, ui.Slate)
	ebitenutil.DrawRect(img, x, y, w*float64(done)/float64(total), h, ui.Gold)
}

// listDraw mostra as minhas tarefas, com ✓ nas prontas.
func (p *Play) listDraw(img *ebiten.Image) {
	ids := p.taskIDs()
	if len(ids) == 0 {
		return
	}
	h := 44 + float64(len(ids))*24
	x, y := 1000.0, 86.0
	ebitenutil.DrawRect(img, x, y, 260, h, color.RGBA{16, 16, 30, 190})
	ebitenutil.DrawRect(img, x, y, 260, 2, ui.Brass)
	p.hudText(img, i18n.T("task.title"), x+14, y+10)
	for k, i := range ids {
		yy := y + 38 + float64(k)*24
		mark, c := "•", ui.Cream
		if p.done[i] {
			mark, c = "✓", ui.Gold
		}
		op := &text.DrawOptions{}
		op.GeoM.Translate(x+14, yy)
		op.ColorScale.ScaleWithColor(c)
		text.Draw(img, mark, p.font, op)
		op = &text.DrawOptions{}
		op.GeoM.Translate(x+34, yy)
		op.ColorScale.ScaleWithColor(c)
		text.Draw(img, taskName(p.mp.Tasks[i].Name), p.font, op)
	}
}

// taskIDs ordena os índices (desenho estável frame a frame).
func (p *Play) taskIDs() []int {
	ids := make([]int, 0, len(p.ts))
	for i := range p.ts {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	return ids
}

// taskName traduz o nome do objeto do Tiled; sem chave, usa o nome.
func taskName(name string) string {
	k := "task." + strings.ToLower(name)
	if v := i18n.T(k); v != k {
		return v
	}
	return name
}
