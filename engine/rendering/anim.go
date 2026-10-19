package ui

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// anim.go — tira de frames de spritesheet já carregada (grade fixa).
// cell = lado do frame, cols = colunas, start = índice inicial.

type Strip struct {
	frames []*ebiten.Image
	i, t   int
	rate   int
	hue    float64
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

// NewHStrip monta uma tira horizontal de frames quadrados (largura = N * altura).
// Suporta qualquer resolução de sprite (16px, 32px, 64px, 128px).
func NewHStrip(img *ebiten.Image, rate int) *Strip {
	s := &Strip{rate: rate}
	if img == nil {
		return s
	}
	cell := img.Bounds().Dy()
	if cell <= 0 {
		return s
	}
	n := img.Bounds().Dx() / cell
	for k := 0; k < n; k++ {
		s.frames = append(s.frames, img.SubImage(image.Rect(k*cell, 0, (k+1)*cell, cell)).(*ebiten.Image))
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

// Hue gira o matiz do desenho (0 = cor da arte; cor do pato).
func (s *Strip) Hue(h float64) { s.hue = h }

// tint aplica o matiz no desenho, se houver.
func (s *Strip) tint(op *ebiten.DrawImageOptions) {
	if s.hue != 0 {
		op.ColorM.RotateHue(s.hue)
	}
}

// DrawFallen desenha o frame base deitado (corpo caído) e escuro.
func (s *Strip) DrawFallen(img *ebiten.Image, x, y, scale float64) {
	if len(s.frames) == 0 {
		return
	}
	fr := s.frames[0]
	dw, dh := float64(fr.Bounds().Dx()), float64(fr.Bounds().Dy())
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	s.tint(op)
	op.GeoM.Translate(-dw/2, -dh/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Rotate(math.Pi / 2)
	op.GeoM.Translate(0, -dh*scale/4)
	op.GeoM.Translate(x, y)
	op.ColorScale.Scale(0.5, 0.42, 0.42, 1)
	img.DrawImage(fr, op)
}

// Draw centraliza o frame em (x,y); face -1 espelha.
func (s *Strip) Draw(img *ebiten.Image, x, y, scale, face float64) {
	if len(s.frames) == 0 {
		return
	}
	fr := s.frames[s.i]
	dw, dh := float64(fr.Bounds().Dx()), float64(fr.Bounds().Dy())
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterNearest}
	s.tint(op)
	op.GeoM.Translate(-dw/2, -dh/2)
	op.GeoM.Scale(scale*face, scale)
	op.GeoM.Translate(x, y)
	img.DrawImage(fr, op)
}
