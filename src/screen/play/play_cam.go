package play

import (
	"duckling/src/world"
)

// play_cam.go — câmera presa na área, sempre atrás do pato.

// clampCam prende v em [lo, hi-view]; se couber, centra.
func clampCam(v float64, lo, hi int, view float64) float64 {
	if float64(hi-lo) <= view {
		return (float64(lo+hi) - view) / 2
	}
	return min(max(v, float64(lo)), float64(hi)-view)
}

// bounds devolve a área atual (ou o visível do mapa, se não houver).
func (p *Play) bounds() (int, int, int, int) {
	if p.area < len(p.mp.Areas) {
		r := p.mp.Areas[p.area].Rect
		return r.Min.X, r.Min.Y, r.Max.X, r.Max.Y
	}
	r := p.mp.View()
	return r.Min.X, r.Min.Y, r.Max.X, r.Max.Y
}

// follow gruda a câmera no pato, respeitando as bordas.
// Na intro (2.5s) o zoom sai do mapa todo até o gameplay,
// com o foco deslizando do centro do mapa até o pato.
func (p *Play) follow() {
	p.zT = min(p.zT+1.0/(60.0*2.5), 1)
	e := p.zT * p.zT * (3 - 2*p.zT)
	z := p.zFrom + (world.Zoom-p.zFrom)*e
	fx := float64(p.mp.W)/2 + (p.d.Px-float64(p.mp.W)/2)*e
	fy := float64(p.mp.H)/2 + (p.d.Py-float64(p.mp.H)/2)*e
	vw, vh := 1280/z, 720/z
	x0, y0, x1, y1 := p.bounds()
	p.zoom = z
	p.camX = clampCam(fx-vw/2, x0, x1, vw)
	p.camY = clampCam(fy-vh/2, y0, y1, vh)
}
