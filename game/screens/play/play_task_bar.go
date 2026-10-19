package play

import (
	"image"
	"image/color"

	"duckling/engine/rendering"
	"duckling/game/map"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// play_task_bar.go — barrinha da tarefa em andamento (no ponto) e
// marcador do destino do carry.

func (p *Play) barTaskDraw(img *ebiten.Image) {
	if p.ghost() {
		return
	}
	if i := p.carrying(); i >= 0 {
		p.targetMark(img, p.mp.Tasks[i].Target())
	}
	for i, st := range p.ts {
		if st.Done() || st.Kind == world.KindCarry || st.Prog <= 0 {
			continue
		}
		at := p.mp.Tasks[i].Center()
		x := (float64(at.X)-p.camX)*p.zoom - 32
		y := (float64(at.Y)-p.camY)*p.zoom - 34
		ebitenutil.DrawRect(img, x-3, y-3, 70, 14, color.RGBA{0, 0, 0, 190})
		ebitenutil.DrawRect(img, x, y, 64, 8, ui.Slate)
		ebitenutil.DrawRect(img, x, y, 64*st.Prog, 8, ui.Gold)
		if st.Kind == world.KindOrder {
			// 3 pips: soltar zera, as divisórias mostram a sequência.
			for k := 1; k < 3; k++ {
				ebitenutil.DrawRect(img, x+64*float64(k)/3-1, y, 2, 8, ui.Ink)
			}
		}
	}
}

// targetMark marca o destino do carry no chão.
func (p *Play) targetMark(img *ebiten.Image, at image.Point) {
	x := (float64(at.X)-p.camX)*p.zoom - 14
	y := (float64(at.Y)-p.camY)*p.zoom - 14
	ebitenutil.DrawRect(img, x, y, 28, 28, color.RGBA{232, 190, 72, 60})
	ebitenutil.DrawRect(img, x, y, 28, 2, ui.Gold)
	ebitenutil.DrawRect(img, x, y+26, 28, 2, ui.Gold)
	ebitenutil.DrawRect(img, x, y, 2, 28, ui.Gold)
	ebitenutil.DrawRect(img, x+26, y, 2, 28, ui.Gold)
}
