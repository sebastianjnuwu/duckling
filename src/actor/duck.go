package actor

import (
	"duckling/src/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

// duck.go — personagem: sprite, direção e desenho.
// Física e entrada ficam em move.go.

type Duck struct {
	side, down, up, fly *ui.Strip
	def                 Def
	dir                 int
	face                float64
	moving              bool
	kx, ky              float64
	z, vz               float64
	flyT, holdT          float64
	flying, wantFly     bool
	hx, hy              bool
	flap                uint
	slowT               float64
	bumpT               float64
	Px, Py              float64
}

func New(img *ebiten.Image, x, y float64, d Def) *Duck {
	du := &Duck{def: d, face: 1, flyT: flyMax, Px: x, Py: y}
	du.Load(img)
	return du
}

// Load (re)monta as tiras; imagens só existem após o loading.
func (d *Duck) Load(img *ebiten.Image) {
	if img == nil {
		return
	}
	s, dn, up := d.def.Side, d.def.Down, d.def.Up
	c, cs := d.def.Cell, d.def.Cols
	d.side = ui.NewDuck(img, s, c, cs)
	d.down = ui.NewDuck(img, dn, c, cs)
	d.up = ui.NewDuck(img, up, c, cs)
	d.fly = ui.NewFlap(img, d.def.Fly, d.def.FlapUp, c, cs, 6)
}

// Ok diz se o sprite já carregou.
func (d *Duck) Ok() bool {
	return d.side != nil && !d.side.Empty()
}

// Alt devolve a altura atual (px de mundo).
func (d *Duck) Alt() float64 { return d.z }

// Move anda sem colisão (portas, empurrões, testes).
func (d *Duck) Move(dx, dy float64) {
	d.Px += dx
	d.Py += dy
	d.Anim(dx, dy)
}

func (d *Duck) Draw(img *ebiten.Image, x, y, s float64) {
	drawShadow(img, x, y, s, d.z)
	st := d.cur()
	f := d.face
	if d.dir != 0 {
		f = 1
	}
	if st != nil {
		st.Draw(img, x, y-d.z*s, s*d.def.Scale, f)
	}
}

// cur devolve a tira da direção atual (voando de lado usa a de voo).
func (d *Duck) cur() *ui.Strip {
	if d.flying && d.dir == 0 {
		return d.fly
	}
	switch d.dir {
	case 2:
		return d.up
	case 1:
		return d.down
	default:
		return d.side
	}
}
