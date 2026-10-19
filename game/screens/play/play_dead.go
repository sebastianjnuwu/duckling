package play

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// play_dead.go — os corpos caídos: posição vem do host (ou do
// snapshot) e o desenho é o pato deitado.

// deadDraw desenha os corpos caídos no chão.
func (p *Play) deadDraw(img *ebiten.Image) {
	if p.sess == nil {
		return
	}
	for _, b := range p.sess.Bodies() {
		x, y := (b.X-p.camX)*p.zoom, (b.Y-p.camY)*p.zoom
		p.d.DrawDead(img, x, y, p.zoom)
	}
}
