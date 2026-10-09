package play

import (
	"image/color"

	"duckling/assets"
	"duckling/src/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// play_draw.go — bg atrás, mapa centralizado, pato no spawn.

func (p *Play) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	p.mp.Draw(img, p.camX, p.camY, p.zoom)
	p.d.Draw(img, (p.d.Px-p.camX)*p.zoom, (p.d.Py-p.camY)*p.zoom, p.zoom)
	// Raio-X: sólidos em vermelho, pés do pato em verde.
	if p.dbg {
		z := p.zoom
		for _, s := range p.mp.Solids {
			x, y := (float64(s.Min.X)-p.camX)*z, (float64(s.Min.Y)-p.camY)*z
			ebitenutil.DrawRect(img, x, y, float64(s.Dx())*z, float64(s.Dy())*z, color.RGBA{255, 0, 0, 128})
		}
		fx, fy := (p.d.Px-10-p.camX)*z, (p.d.Py-8-p.camY)*z
		ebitenutil.DrawRect(img, fx, fy, 20*z, 16*z, color.RGBA{0, 255, 0, 128})
	}
}
