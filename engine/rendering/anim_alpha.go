package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// anim_alpha.go — a mesma tira, translúcida (pato fantasma).

// DrawAlpha desenha o frame atual com alpha (0..1).
func (s *Strip) DrawAlpha(img *ebiten.Image, x, y, scale, face, alpha float64) {
	if len(s.frames) == 0 {
		return
	}
	fr := s.frames[s.i]
	dw, dh := float64(fr.Bounds().Dx()), float64(fr.Bounds().Dy())
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	s.tint(op)
	op.ColorScale.ScaleAlpha(float32(alpha))
	op.GeoM.Translate(-dw/2, -dh/2)
	op.GeoM.Scale(scale*face, scale)
	op.GeoM.Translate(x, y)
	img.DrawImage(fr, op)
}
