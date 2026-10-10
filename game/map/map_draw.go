package world

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// map_draw.go — o mapa virou uma imagem só: 1 DrawImage por frame,
// com zoom (tiles de 24px, pixel crocante no Nearest).

// Zoom do mapa na tela (4x estilo Stardew: tiles de 64px).
const Zoom = 4.0

func (m *Map) Draw(img *ebiten.Image, camX, camY, z float64) {
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	op.GeoM.Scale(z, z)
	op.GeoM.Translate(-camX*z, -camY*z)
	img.DrawImage(m.img, op)
}
