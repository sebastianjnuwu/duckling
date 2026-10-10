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
}

// Burst solta n faíscas de (x,y) nas cores dadas.
func (s *Sparks) Burst(x, y float64, n int, cs ...color.Color) {
	for i := 0; i < n; i++ {
		a := rand.Float64() * 6.28
		sp := 60 + rand.Float64()*160
		s.ps = append(s.ps, spark{
			x: x, y: y,
			vx: math.Cos(a) * sp, vy: math.Sin(a)*sp - 140,
			life: 0.4 + rand.Float64()*0.4, max: 0.8,
			size: 3 + rand.Float64()*3, c: cs[i%len(cs)],
		})
	}
}

func (s *Sparks) Update() {
	dt := 1.0 / 60.0
	keep := s.ps[:0]
	for _, p := range s.ps {
		p.vy += 600 * dt
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
