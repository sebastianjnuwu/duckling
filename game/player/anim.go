package actor

import (
	"duckling/engine/rendering"
)

// anim.go — direção e quadros: o 1º toque troca a pose na hora,
// parado volta ao frame base.

func (d *Duck) Anim(dx, dy float64) {
	if d.down == nil {
		return
	}
	if dx == 0 && dy == 0 && !d.flying && d.z <= d.ground+0.1 {
		d.down.Reset()
		d.up.Reset()
		d.left.Reset()
		d.right.Reset()
		d.moving = false
		return
	}
	if dx != 0 || dy != 0 {
		switch {
		case dy < 0:
			d.dir = 1
		case dy > 0:
			d.dir = 0
		case dx < 0:
			d.dir = 2
		case dx > 0:
			d.dir = 3
		}
	}
	if d.flying {
		d.flap++
		st := d.cur()
		st.Tick()
		if d.z < FlyPeak-1 || d.flap%2 == 0 {
			st.Tick()
		}
		return
	}
	if !d.moving {
		d.moving = true
		d.cur().Step()
	}
	d.Tick()
}

func (d *Duck) Tick() {
	if st := d.cur(); st != nil {
		st.Tick()
	}
}

// cur devolve a tira da ação e direção atual.
func (d *Duck) cur() *ui.Strip {
	if d.flying {
		switch d.dir {
		case 1:
			return d.flyUp
		case 2:
			return d.flyLeft
		case 3:
			return d.flyRight
		default:
			return d.flyDown
		}
	}
	if d.z > d.ground+0.1 {
		switch d.dir {
		case 1:
			return d.jumpUp
		case 2:
			return d.jumpLeft
		case 3:
			return d.jumpRight
		default:
			return d.jumpDown
		}
	}
	switch d.dir {
	case 1:
		return d.up
	case 2:
		return d.left
	case 3:
		return d.right
	default:
		return d.down
	}
}
