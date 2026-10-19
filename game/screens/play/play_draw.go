package play

import (
	"image/color"
	"sort"

	"duckling/content"
	"duckling/engine/rendering"
	"duckling/game/match"
	"duckling/game/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// play_draw.go — bg atrás, mapa centralizado, remotos e o pato
// local por cima (com nome em cima de quem é de fora).

func (p *Play) Draw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	p.mp.Draw(img, p.camX, p.camY, p.zoom)
	p.gateDraw(img)
	p.deadDraw(img)
	for _, id := range p.remoteIDs() {
		d := p.rem[id]
		x, y := (d.Px-p.camX)*p.zoom, (d.Py-p.camY)*p.zoom
		if p.rghost[id] {
			d.DrawGhost(img, x, y, p.zoom)
		} else {
			d.Draw(img, x, y, p.zoom)
		}
		p.nameTag(img, d, p.rname[id])
	}
	if p.ghost() {
		p.d.DrawGhost(img, (p.d.Px-p.camX)*p.zoom, (p.d.Py-p.camY)*p.zoom, p.zoom)
	} else {
		p.d.Draw(img, (p.d.Px-p.camX)*p.zoom, (p.d.Py-p.camY)*p.zoom, p.zoom)
	}
	p.d.DrawFx(img, p.camX, p.camY, p.zoom)
	p.puff.Draw(img)
	p.darkDraw(img)
	// Raio-X: sólidos em vermelho, pés do pato em verde.
	if p.dbg {
		p.drawDbg(img)
	}
	p.hudDraw(img)
	p.taskDraw(img)
	p.killDraw(img)
	p.sabDraw(img)
	p.hostDraw(img)
	if p.phase == match.End {
		p.overDraw(img)
	}
	if p.confirm {
		p.confirmDraw(img)
	}
	p.toast.Draw(img, p.font)
}

// remoteIDs ordena por id pra desenho estável frame a frame.
func (p *Play) remoteIDs() []int {
	ids := make([]int, 0, len(p.rem))
	for id := range p.rem {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	return ids
}

// nameTag escreve o nome acima de um pato remoto.
func (p *Play) nameTag(img *ebiten.Image, d *actor.Duck, name string) {
	if name == "" || p.font == nil {
		return
	}
	w, h := text.Measure(name, p.font, 0)
	x := (d.Px-p.camX)*p.zoom - w/2
	y := (d.Py-p.camY)*p.zoom - 50*p.zoom - h
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.RGBA{235, 240, 235, 220})
	text.Draw(img, name, p.font, op)
}

func (p *Play) drawDbg(img *ebiten.Image) {
	z := p.zoom
	for _, s := range p.mp.Solids {
		x, y := (float64(s.Rect.Min.X)-p.camX)*z, (float64(s.Rect.Min.Y)-p.camY)*z
		ebitenutil.DrawRect(img, x, y, float64(s.Rect.Dx())*z, float64(s.Rect.Dy())*z, color.RGBA{255, 0, 0, 128})
	}
	fx, fy := (p.d.Px-10-p.camX)*z, (p.d.Py-8-p.camY)*z
	ebitenutil.DrawRect(img, fx, fy, 20*z, 16*z, color.RGBA{0, 255, 0, 128})
}
