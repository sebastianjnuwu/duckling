package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// play_area.go — teclas 1-5: teleporta pra zona "areas" do mapa
// (mapa de teste); sem áreas, não faz nada.

func (p *Play) areaKeys() {
	keys := []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5}
	for i, k := range keys {
		if inpututil.IsKeyJustPressed(k) && i < len(p.mp.Areas) {
			p.area = i
			r := p.mp.Areas[i].Rect
			p.d.Px, p.d.Py = float64(r.Min.X+r.Max.X)/2, float64(r.Min.Y+r.Max.Y)/2
			p.follow()
		}
	}
}
