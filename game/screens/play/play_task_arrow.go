package play

import (
	"math"

	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// play_task_arrow.go — seta na borda apontando pra tarefa mais
// perto quando ela está fora da tela.

func (p *Play) arrowDraw(img *ebiten.Image) {
	if p.ghost() {
		return
	}
	best, bd := -1, math.MaxFloat64
	for i, st := range p.ts {
		if st.Done() {
			continue
		}
		if d := p.dist(p.mp.Tasks[i].Center()); d < bd {
			best, bd = i, d
		}
	}
	if best < 0 {
		return
	}
	at := p.mp.Tasks[best].Center()
	x := (float64(at.X) - p.camX) * p.zoom
	y := (float64(at.Y) - p.camY) * p.zoom
	if x >= 0 && x <= 1280 && y >= 0 && y <= 720 {
		return // já está na tela
	}
	p.edgeArrow(img, x-640, y-360)
}

// edgeArrow põe o ícone na borda dominante, apontando pro alvo.
func (p *Play) edgeArrow(img *ebiten.Image, dx, dy float64) {
	x, y := 640+dx*0.45, 360+dy*0.45
	if math.Abs(dx) > math.Abs(dy) {
		if dx > 0 {
			x = 1236
			ui.DrawIcon(img, ui.IconRight, x, clampF(y, 60, 636), ui.Gold)
			return
		}
		ui.DrawIcon(img, ui.IconLeft, 20, clampF(y, 60, 636), ui.Gold)
		return
	}
	if dy > 0 {
		ui.DrawIcon(img, ui.IconDown, clampF(x, 60, 1220), 676, ui.Gold)
		return
	}
	ui.DrawIcon(img, ui.IconUp, clampF(x, 60, 1220), 20, ui.Gold)
}

func clampF(v, lo, hi float64) float64 {
	return min(max(v, lo), hi)
}
