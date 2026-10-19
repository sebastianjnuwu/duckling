package play

import (
	"image/color"

	"duckling/game/map"
	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// play_gate.go — porteiras e escuro da sabotagem: colisão extra
// enquanto fechadas e foco de luz por pato vivo.

// solids devolve os sólidos do mapa + as porteiras fechadas.
func (p *Play) solids() []world.Solid {
	if p.sess == nil || p.sess.GateLeft() <= 0 || len(p.mp.Gates) == 0 {
		return p.mp.Solids
	}
	out := make([]world.Solid, 0, len(p.mp.Solids)+len(p.mp.Gates))
	out = append(out, p.mp.Solids...)
	for _, r := range p.mp.Gates {
		out = append(out, world.Solid{Rect: r, H: 1000})
	}
	return out
}

// gateDraw pinta as porteiras fechadas por cima do mapa.
func (p *Play) gateDraw(img *ebiten.Image) {
	if p.sess == nil || p.sess.GateLeft() <= 0 {
		return
	}
	for _, r := range p.mp.Gates {
		x := (float64(r.Min.X) - p.camX) * p.zoom
		y := (float64(r.Min.Y) - p.camY) * p.zoom
		w, h := float64(r.Dx())*p.zoom, float64(r.Dy())*p.zoom
		ebitenutil.DrawRect(img, x, y, w, h, color.RGBA{104, 68, 40, 240})
		ebitenutil.DrawRect(img, x, y, w, 3, color.RGBA{64, 40, 22, 255})
		ebitenutil.DrawRect(img, x, y+h-3, w, 3, color.RGBA{64, 40, 22, 255})
	}
}

// darkDraw aplica o escuro da lamparina com um foco por pato vivo.
func (p *Play) darkDraw(img *ebiten.Image) {
	if p.sess == nil || p.sess.LampLeft() <= 0 || p.ghost() {
		return
	}
	focus := make([][2]float64, 0, len(p.rem)+1)
	focus = append(focus, p.light(p.d))
	for id, d := range p.rem {
		if !p.rghost[id] {
			focus = append(focus, p.light(d))
		}
	}
	p.dark.Draw(img, focus, 150)
}

// light converte o pato em foco de luz (px de tela).
func (p *Play) light(d *actor.Duck) [2]float64 {
	return [2]float64{(d.Px - p.camX) * p.zoom, (d.Py - 24 - p.camY) * p.zoom}
}
