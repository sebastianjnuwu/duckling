package actor

import (
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// duck.go — personagem: sprite, direção e desenho.
// Física e entrada ficam em move.go.

type Duck struct {
	down, up, left, right                 *ui.Strip
	flyDown, flyUp, flyLeft, flyRight     *ui.Strip
	jumpDown, jumpUp, jumpLeft, jumpRight *ui.Strip
	fx                                    ui.Sparks
	def                                   Def
	dir                                   int
	col                                   Color
	moving                                bool
	vx, vy                                float64
	coyT, bufT                            float64
	jumpHeld                              bool
	z, vz, ground                         float64
	flyT                                  float64
	flying, wantFly                       bool
	flap                                  uint
	Px, Py                                float64
}

func New(imgs map[string]*ebiten.Image, x, y float64, d Def) *Duck {
	du := &Duck{def: d, flyT: flyMax, Px: x, Py: y}
	du.fx.Grav = 120
	du.fx.Size = 0.5
	du.Load(imgs)
	return du
}

// Load monta as 12 tiras a partir das imagens carregadas.
func (d *Duck) Load(imgs map[string]*ebiten.Image) {
	if imgs == nil {
		return
	}
	p := d.def.Prefix
	if p == "" {
		p = "sprites/duck/yellow/128px"
	}
	d.down = ui.NewHStrip(imgs[p+"/down"], 8)
	d.up = ui.NewHStrip(imgs[p+"/up"], 8)
	d.left = ui.NewHStrip(imgs[p+"/left"], 8)
	d.right = ui.NewHStrip(imgs[p+"/right"], 8)

	d.flyDown = ui.NewHStrip(imgs[p+"/fly_down"], 8)
	d.flyUp = ui.NewHStrip(imgs[p+"/fly_up"], 8)
	d.flyLeft = ui.NewHStrip(imgs[p+"/fly_left"], 8)
	d.flyRight = ui.NewHStrip(imgs[p+"/fly_right"], 8)

	d.jumpDown = ui.NewHStrip(imgs[p+"/jump_down"], 8)
	d.jumpUp = ui.NewHStrip(imgs[p+"/jump_up"], 8)
	d.jumpLeft = ui.NewHStrip(imgs[p+"/jump_left"], 8)
	d.jumpRight = ui.NewHStrip(imgs[p+"/jump_right"], 8)
	for _, s := range d.strips() {
		s.Hue(d.col.Hue())
	}
}

// Ok diz se o sprite já carregou.
func (d *Duck) Ok() bool {
	return d.down != nil && !d.down.Empty()
}

// Alt devolve a altura atual (px de mundo).
func (d *Duck) Alt() float64 { return d.z }

// Move anda sem colisão (portas, testes).
func (d *Duck) Move(dx, dy float64) {
	d.Px += dx
	d.Py += dy
	d.Anim(dx, dy)
}

func (d *Duck) Draw(img *ebiten.Image, x, y, s float64) {
	drawShadow(img, x, y, s, d.z-d.ground, d.ground)
	if st := d.cur(); st != nil {
		st.Draw(img, x, y-d.z*s, s*d.def.Scale, 1)
	}
}
