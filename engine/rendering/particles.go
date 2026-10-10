package ui

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// particles.go — faíscas pixel: nascem, sobem e morrem.
// Sem estado de jogo aqui, só física boba.

type spark struct {
	x, y, vx, vy, life, max, size float64
	c                             color.Color
}

type Sparks struct {
	ps []spark
	// Grav é a gravidade (0 = 600 padrão; pena usa ~120).
	Grav float64
	// Size multiplica o tamanho (0 = 1; pena usa 0.5).
	Size float64
}

// Burst solta n faíscas de (x,y) nas cores dadas.
func (s *Sparks) Burst(x, y float64, n int, cs ...color.Color) {
	sz := s.Size
	if sz == 0 {
		sz = 1
	}
	for i := 0; i < n; i++ {
		a := rand.Float64() * 6.28
		sp := 60 + rand.Float64()*160
		s.ps = append(s.ps, spark{
			x: x, y: y,
			vx: math.Cos(a) * sp, vy: math.Sin(a)*sp - 140,
			life: 0.4 + rand.Float64()*0.4, max: 0.8,
			size: (3 + rand.Float64()*3) * sz, c: cs[i%len(cs)],
		})
	}
}

func (s *Sparks) Update() {
	dt := 1.0 / 60.0
	g := s.Grav
	if g == 0 {
		g = 600
	}
	keep := s.ps[:0]
	for _, p := range s.ps {
		p.vy += g * dt
		p.x += p.vx * dt
		p.y += p.vy * dt
		p.life -= dt
		if p.life > 0 {
			keep = append(keep, p)
		}
	}
	s.ps = keep
}

func (s *Sparks) Draw(img *ebiten.Image) {
	for _, p := range s.ps {
		r, g, b, _ := p.c.RGBA()
		c := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(255 * p.life / p.max)}
		ebitenutil.DrawRect(img, p.x, p.y, p.size, p.size, c)
	}
}

// DrawAt atualiza e desenha com câmera (partículas em coords de mundo).
func (s *Sparks) DrawAt(img *ebiten.Image, camX, camY, zoom float64) {
	s.Update()
	for _, p := range s.ps {
		r, g, b, _ := p.c.RGBA()
		c := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(255 * p.life / p.max)}
		ebitenutil.DrawRect(img, (p.x-camX)*zoom, (p.y-camY)*zoom, p.size*zoom, p.size*zoom, c)
	}
}
