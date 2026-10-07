package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// anim.go — tira de frames de spritesheet já carregada (grade fixa).
// cell = lado do frame, cols = colunas, start = índice inicial.

type Strip struct {
	frames []*ebiten.Image
	i, t   int
	rate   int
}

func NewStrip(img *ebiten.Image, cell, cols, start, n, rate int) *Strip {
	s := &Strip{rate: rate}
	if img == nil {
		return s
	}
	for k := start; k < start+n; k++ {
		x, y := (k%cols)*cell, (k/cols)*cell
		s.frames = append(s.frames, img.SubImage(image.Rect(x, y, x+cell, y+cell)).(*ebiten.Image))
	}
	return s
}

func (s *Strip) Tick() {
	if len(s.frames) == 0 {
		return
	}
	s.t++
	if s.t >= s.rate {
		s.t, s.i = 0, (s.i+1)%len(s.frames)
	}
}

// Draw centraliza o frame em (x,y); face -1 espelha.
func (s *Strip) Draw(img *ebiten.Image, x, y, scale, face float64) {
	if len(s.frames) == 0 {
		return
	}
	fr := s.frames[s.i]
	dw, dh := float64(fr.Bounds().Dx()), float64(fr.Bounds().Dy())
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	op.GeoM.Translate(-dw/2, -dh/2)
	op.GeoM.Scale(scale*face, scale)
	op.GeoM.Translate(x, y)
	img.DrawImage(fr, op)
}
