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

// NewDuck monta o personagem da fileira row da sheet do pato
// (24 amarelos, 10 laranjas, 8 brancos, 13 laranjas, 10 amarelos).
// cell/cols descrevem a grade: 32/24 no normal, 128/24 no 4x.
func NewDuck(img *ebiten.Image, row, cell, cols int) *Strip {
	n := []int{24, 10, 8, 13, 10}[row%5]
	return NewStrip(img, cell, cols, row*cols, n, 8)
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

// Reset volta ao frame base (parado).
func (s *Strip) Reset() {
	s.i, s.t = 0, 0
}

// Step avança 1 frame na hora (pose imediata no 1º toque).
func (s *Strip) Step() {
	if len(s.frames) == 0 {
		return
	}
	s.t, s.i = 0, (s.i+1)%len(s.frames)
}

// Empty diz se a tira não tem frames (imagem ainda não carregou).
func (s *Strip) Empty() bool {
	return len(s.frames) == 0
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
